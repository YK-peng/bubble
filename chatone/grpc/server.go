package grpc

import (
	"chatone"
	"chatone/conf"
	"context"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

// New chatone grpc server.
func New(c *conf.RPCServer, s *chatone.Server) *grpc.Server {
	keepParams := grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle:     time.Duration(c.IdleTimeout),
		MaxConnectionAgeGrace: time.Duration(c.ForceCloseWait),
		Time:                  time.Duration(c.KeepAliveInterval),
		Timeout:               time.Duration(c.KeepAliveTimeout),
		MaxConnectionAge:      time.Duration(c.MaxLifeTime),
	})
	srv := grpc.NewServer(keepParams)
	RegisterChatServer(srv, &server{s})
	lis, err := net.Listen(c.Network, c.Addr)
	if err != nil {
		panic(err)
	}
	go func() {
		if err := srv.Serve(lis); err != nil {
			panic(err)
		}
	}()
	return srv
}

type server struct {
	srv *chatone.Server
}

var _ ChatServer = &server{}

func (s server) SendChatMsg(context.Context, *NewChatMsg) (*NewChatMsgAck, error) {
	panic("implement me")
}

func (s server) SendAppMsg(context.Context, *NewAppMsg) (*NewAppMsgAck, error) {
	panic("implement me")
}

func (s server) RevokeMsg(context.Context, *RevokeMsgReq) (*RevokeMsgRep, error) {
	panic("implement me")
}

func (s server) ReadMsg(context.Context, *MsgReadReq) (*MsgReadRep, error) {
	panic("implement me")
}

func (s server) GetMsgUnread(context.Context, *MsgReadReq) (*MsgReadRep, error) {
	panic("implement me")
}
