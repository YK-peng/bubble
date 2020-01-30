package main

import (
	"auth"
	"auth/conf"
	"auth/grpc"
	"auth/xframe/xlog"
	"flag"
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

	xlog.NewGbizLogger(conf.Conf.Xlog.Path, conf.Conf.Xlog.Level)
	xlog.Info("bubble-auth start")

	//开始构建bubble-auth server
	srv := auth.NewServer(conf.Conf)
	rpcSrv := grpc.New(conf.Conf.RPCServer, srv)
	// signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		xlog.Infof("bubble-auth get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			srv.Close()
			rpcSrv.GracefulStop()
			xlog.Infof("bubble-auth exit")
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
