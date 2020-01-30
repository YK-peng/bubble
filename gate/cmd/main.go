package main

import (
	"flag"
	"gate"
	"gate/conf"
	"gate/grpc"
	"gate/xframe/xlog"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

func main() {
	//配置
	flag.Parse()
	if err := conf.Init(); err != nil {
		panic(err)
	}
	runtime.GOMAXPROCS(runtime.NumCPU())
	println(conf.Conf.Debug)

	xlog.Info("bubble-gate [version: %s env: %+v] start")
	// register discovery
	//dis := naming.New(conf.Conf.Discovery)
	//resolver.Register(dis)

	//开始构建bubble-gate server
	srv := gate.NewServer(conf.Conf)
	if err := gate.InitTCP(srv, conf.Conf.TCP.Bind, runtime.NumCPU()); err != nil {
		panic(err)
	}
	if err := gate.InitWebsocket(srv, conf.Conf.Websocket.Bind, runtime.NumCPU()); err != nil {
		panic(err)
	}
	if conf.Conf.Websocket.TLSOpen {
		if err := gate.InitWebsocketWithTLS(srv, conf.Conf.Websocket.TLSBind, conf.Conf.Websocket.CertFile, conf.Conf.Websocket.PrivateFile, runtime.NumCPU()); err != nil {
			panic(err)
		}
	}

	//构建grpc server
	rpcSrv := grpc.New(conf.Conf.RPCServer, srv)
	//cancel := register(dis, srv)

	//监听信号
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		xlog.Infof("bubble-gate get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			//if cancel != nil {
			//	cancel()
			//}
			rpcSrv.GracefulStop()
			srv.Close()
			xlog.Infof("bubble-gate exit")
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
