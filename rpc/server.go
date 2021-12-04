package rpc

import (
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"grpc-http-example/service"
	"grpc-http-example/util/conf"
	"grpc-http-example/util/logger"
	"net"
)

// Serve 启动grpc服务
func Serve() {
	// 监听端口
	grpcPort := conf.GetString("app.grpcPort")
	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		logger.Fatal("listen port "+grpcPort+" failed", zap.Error(err))
	}

	// 创建服务
	s := grpc.NewServer()
	registerFuncs := service.GetServiceRegisterFuncs()
	for _, f := range registerFuncs {
		f(s)
	}

	logger.Info("start grpc server...", zap.Any("services", s.GetServiceInfo()))
	err = s.Serve(lis)
	if err != nil {
		logger.Fatal("start grpc server failed: ", zap.Error(err))
	}
}
