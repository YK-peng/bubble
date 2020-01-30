package grpc

import (
	"context"
	"net"
	"router"
	"router/conf"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

// New router grpc server.
func New(c *conf.RPCServer, s *router.Server) *grpc.Server {
	keepParams := grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle:     time.Duration(c.IdleTimeout),
		MaxConnectionAgeGrace: time.Duration(c.ForceCloseWait),
		Time:                  time.Duration(c.KeepAliveInterval),
		Timeout:               time.Duration(c.KeepAliveTimeout),
		MaxConnectionAge:      time.Duration(c.MaxLifeTime),
	})
	srv := grpc.NewServer(keepParams)
	RegisterRouterServer(srv, &server{s})
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
	srv *router.Server
}

var _ RouterServer = &server{}

func (s server) Deliver(context.Context, *CliReq) (*CliRep, error) {
	panic("implement me")
}
