package example

import (
	"google.golang.org/grpc"
	"grpc-http-example/pb/example"
)

func GetRegisterFunc() func(server *grpc.Server) {
	return func(server *grpc.Server) {
		example.RegisterExampleServer(server, &ExampleService{})
	}
}

type ExampleService struct {
	*example.UnimplementedExampleServer
}
