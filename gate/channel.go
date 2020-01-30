package gate

import (
	"sync"

	"gate/xframe/bufio"

	"github.com/Terry-Mao/goim/api/comet/grpc"
)

//对连接的封装
//从Reader协程(in ServeTCP func)中获取client发过来的数据包到CliProto。处理完业务后转换成给客户端的回复数据包，通过signal通知Writer协程(dispatchTCP func)送给client
type Channel struct {
	Writer bufio.Writer
	Reader bufio.Reader

	CliProto Ring
	signal   chan *grpc.Proto

	Next *Channel
	Prev *Channel

	Mid int64
	Key string
	IP  string

	mutex sync.RWMutex
}

// NewChannel new a channel.
func NewChannel(cli, svr int) *Channel {
	c := new(Channel)
	c.CliProto.Init(cli)
	c.signal = make(chan *grpc.Proto, svr)
	return c
}

// Ready check the channel ready or close?
func (c *Channel) Ready() *grpc.Proto {
	return <-c.signal
}

func (c *Channel) Push(p *grpc.Proto) (err error) {
	select {
	case c.signal <- p:
	default:
	}
	return
}

// Signal send signal to the channel, protocol ready.
func (c *Channel) Signal() {
	c.signal <- grpc.ProtoReady
}

// Close close the channel.
func (c *Channel) Close() {
	c.signal <- grpc.ProtoFinish
}
