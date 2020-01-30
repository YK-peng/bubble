package main

import (
	"flag"
	"os"
	"os/signal"
	"router"
	"router/conf"
	"router/grpc"
	"router/xframe/xlog"
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
	xlog.Info("bubble-router start")

	//开始构建bubble-router server
	srv := router.NewServer(conf.Conf)

	rpcSrv := grpc.New(conf.Conf.RPCServer, srv)

	//监听信号
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		xlog.Infof("bubble-router get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			//if cancel != nil {
			//	cancel()
			//}
			rpcSrv.GracefulStop()
			srv.Close()
			xlog.Infof("bubble-router exit")
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
