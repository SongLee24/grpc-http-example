package http

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"grpc-http-example/pb/example"
	"grpc-http-example/util/conf"
	"grpc-http-example/util/logger"
	"net/http"
)

func Serve() {
	grpcPort := conf.GetString("app.grpcPort")
	httpPort := conf.GetString("app.httpPort")
	// 开启http服务
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	//  注册gRPC服务endpoint
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := example.RegisterExampleHandlerFromEndpoint(ctx, mux, "127.0.0.1:"+grpcPort, opts)
	if err != nil {
		logger.Fatal("RegisterExampleHandlerFromEndpoint failed: " + err.Error())
	}

	// 启动http
	err = http.ListenAndServe(":"+httpPort, mux)
	if err != nil {
		logger.Fatal("start http server failed: ", zap.Error(err))
	}
}
