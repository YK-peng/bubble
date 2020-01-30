package gate

import (
	"context"
	"io"
	"net"
	"strings"
	"time"

	"gate/conf"
	"gate/grpc"
	"gate/xframe/bufio"
	"gate/xframe/bytes"
	xtime "gate/xframe/time"
	"gate/xframe/xlog"
)

const (
	maxInt = 1<<31 - 1
)

func InitTCP(server *Server, addrs []string, accept int) (err error) {
	var (
		bind     string
		listener *net.TCPListener
		addr     *net.TCPAddr
	)
	for _, bind = range addrs {
		if addr, err = net.ResolveTCPAddr("tcp", bind); err != nil {
			xlog.Errorf("net.ResolveTCPAddr(tcp, %s) error(%v)", bind, err)
			return
		}
		if listener, err = net.ListenTCP("tcp", addr); err != nil {
			xlog.Errorf("net.ListenTCP(tcp, %s) error(%v)", bind, err)
			return
		}
		xlog.Infof("start tcp listen: %s", bind)

		for i := 0; i < accept; i++ {
			go acceptTCP(server, listener)
		}
	}
	return
}

func acceptTCP(server *Server, lis *net.TCPListener) {
	var (
		conn *net.TCPConn
		err  error
		r    int
	)
	for {
		if conn, err = lis.AcceptTCP(); err != nil {
			xlog.Errorf("listener.Accept(\"%s\") error(%v)", lis.Addr().String(), err)
			return
		}
		if err = conn.SetKeepAlive(server.c.TCP.KeepAlive); err != nil {
			xlog.Errorf("conn.SetKeepAlive() error(%v)", err)
			return
		}
		if err = conn.SetReadBuffer(server.c.TCP.Rcvbuf); err != nil {
			xlog.Errorf("conn.SetReadBuffer() error(%v)", err)
			return
		}
		if err = conn.SetWriteBuffer(server.c.TCP.Sndbuf); err != nil {
			xlog.Errorf("conn.SetWriteBuffer() error(%v)", err)
			return
		}
		go serveTCP(server, conn, r)
		if r++; r == maxInt {
			r = 0
		}
	}
}

func serveTCP(s *Server, conn *net.TCPConn, r int) {
	var (
		tr = s.round.Timer(r)
		rp = s.round.Reader(r)
		wp = s.round.Writer(r)

		lAddr = conn.LocalAddr().String()
		rAddr = conn.RemoteAddr().String()
	)
	if conf.Conf.Debug {
		xlog.Infof("start tcp serve \"%s\" with \"%s\"", lAddr, rAddr)
	}
	s.ServeTCP(conn, rp, wp, tr)
}

// ServeTCP serve a tcp connection.
func (s *Server) ServeTCP(conn *net.TCPConn, rp, wp *bytes.Pool, tr *xtime.Timer) {
	var (
		err    error
		hb     time.Duration
		p      *grpc.Proto
		b      *Bucket
		trd    *xtime.TimerData
		lastHb = time.Now()
		rb     = rp.Get()
		wb     = wp.Get()
		ch     = NewChannel(s.c.Protocol.CliProto, s.c.Protocol.SvrProto)
		rr     = &ch.Reader
		wr     = &ch.Writer
	)

	//读写缓冲区
	ch.Reader.ResetBuffer(conn, rb.Bytes())
	ch.Writer.ResetBuffer(conn, wb.Bytes())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// handshake
	step := 0
	trd = tr.Add(time.Duration(s.c.Protocol.HandshakeTimeout), func() {
		conn.Close()
		xlog.Errorf("key: %s remoteIP: %s step: %d tcp handshake timeout", ch.Key, conn.RemoteAddr().String(), step)
	})
	ch.IP, _, _ = net.SplitHostPort(conn.RemoteAddr().String())
	// must not setadv, only used in auth
	step = 1
	if p, err = ch.CliProto.Set(); err == nil {
		if ch.Mid, ch.Key, _, _, hb, err = s.authTCP(ctx, rr, wr, p); err == nil {
			b = s.Bucket(ch.Key)
			err = b.Put(ch)
			if conf.Conf.Debug {
				xlog.Infof("tcp connnected key:%s mid:%d proto:%+v", ch.Key, ch.Mid, p)
			}
		}
	}

	// 握手失败，回收资源
	if err != nil {
		conn.Close()
		rp.Put(rb)
		wp.Put(wb)
		tr.Del(trd)
		xlog.Errorf("key: %s handshake failed error(%v)", ch.Key, err)
		return
	}
	trd.Key = ch.Key
	tr.Set(trd, hb)

	step = 3
	// hanshake ok start dispatch goroutine

	go s.dispatchTCP(conn, wr, wp, wb, ch)
	serverHeartbeat := s.RandServerHearbeat()

	//循环读
	for {
		if p, err = ch.CliProto.Set(); err != nil {
			break
		}

		//从client TCP中读取数据包
		if err = p.ReadTCP(rr); err != nil {
			break
		}

		//处理数据包
		if p.Op == grpc.OpHeartbeat {
			tr.Set(trd, hb)
			p.Op = grpc.OpHeartbeatReply
			p.Body = nil
			// NOTE: send server heartbeat for a long time
			if now := time.Now(); now.Sub(lastHb) > serverHeartbeat {
				if err1 := s.Heartbeat(ctx, ch.Mid, ch.Key); err1 == nil {
					lastHb = now
				}
			}
			if conf.Conf.Debug {
				xlog.Infof("tcp heartbeat receive key:%s, mid:%d", ch.Key, ch.Mid)
			}
			step++
		} else {
			if err = s.Operate(ctx, p, ch, b); err != nil {
				break
			}
		}

		//从ring buffer中"放出".变为可读
		ch.CliProto.SetAdv()
		ch.Signal()
	}

	/*退出*/
	if err != nil && err != io.EOF && !strings.Contains(err.Error(), "closed") {
		xlog.Errorf("key: %s server tcp failed error(%v)", ch.Key, err)
	}

	//倒序释放channel关联的资源
	b.Del(ch)
	tr.Del(trd)
	rp.Put(rb)
	conn.Close()
	ch.Close()
	if err = s.Disconnect(ctx, ch.Mid, ch.Key); err != nil {
		xlog.Errorf("key: %s mid: %d operator do disconnect error(%v)", ch.Key, ch.Mid, err)
	}

	if conf.Conf.Debug {
		xlog.Infof("tcp disconnected key: %s mid: %d", ch.Key, ch.Mid)
	}
}

// dispatch accepts connections on the listener and serves requests
// for each incoming connection.  dispatch blocks; the caller typically
// invokes it in a go statement.
func (s *Server) dispatchTCP(conn *net.TCPConn, wr *bufio.Writer, wp *bytes.Pool, wb *bytes.Buffer, ch *Channel) {
	var (
		err    error
		finish bool
		online int32
	)
	if conf.Conf.Debug {
		xlog.Infof("key: %s start dispatch tcp goroutine", ch.Key)
	}
	for {

		var p = ch.Ready()

		if conf.Conf.Debug {
			xlog.Infof("key:%s dispatch msg:%v", ch.Key, *p)
		}
		switch p {
		case grpc.ProtoFinish:
			if conf.Conf.Debug {
				xlog.Infof("key: %s wakeup exit dispatch goroutine", ch.Key)
			}
			finish = true
			goto failed
		case grpc.ProtoReady:
			// fetch message from svrbox(client send)
			for {
				if p, err = ch.CliProto.Get(); err != nil {
					break
				}

				if p.Op == grpc.OpHeartbeatReply {
					if err = p.WriteTCPHeart(wr, online); err != nil {
						goto failed
					}
				} else {
					if err = p.WriteTCP(wr); err != nil {
						goto failed
					}
				}
				p.Body = nil // avoid memory leak
				ch.CliProto.GetAdv()
			}
		default:
			//服务端主动推送给client的业务数据
			if err = p.WriteTCP(wr); err != nil {
				goto failed
			}
			if conf.Conf.Debug {
				xlog.Infof("tcp sent a message key:%s mid:%d proto:%+v", ch.Key, ch.Mid, p)
			}
		}
		// only hungry flush response
		if err = wr.Flush(); err != nil {
			break
		}
	}
failed:
	if err != nil {
		xlog.Errorf("key: %s dispatch tcp error(%v)", ch.Key, err)
	}
	conn.Close()
	wp.Put(wb)
	// must ensure all channel message discard, for reader won't blocking Signal
	for !finish {
		finish = ch.Ready() == grpc.ProtoFinish
	}
	if conf.Conf.Debug {
		xlog.Infof("key: %s dispatch goroutine exit", ch.Key)
	}
}

func (s *Server) authTCP(ctx context.Context, rr *bufio.Reader, wr *bufio.Writer, p *grpc.Proto) (mid int64, key, rid string, accepts []int32, hb time.Duration, err error) {
	for {
		if err = p.ReadTCP(rr); err != nil {
			return
		}
		if p.Op == grpc.OpAuth {
			break
		} else {
			xlog.Errorf("tcp request operation(%d) not auth", p.Op)
		}
	}
	if mid, key, rid, accepts, hb, err = s.Connect(ctx, p, ""); err != nil {
		xlog.Errorf("authTCP.Connect(key:%v).err(%v)", key, err)
		return
	}
	p.Op = grpc.OpAuthReply
	p.Body = nil
	if err = p.WriteTCP(wr); err != nil {
		xlog.Errorf("authTCP.WriteTCP(key:%v).err(%v)", key, err)
		return
	}
	err = wr.Flush()
	return
}
