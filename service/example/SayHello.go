package example

import (
	"context"
	"grpc-http-example/pb/example"
)

func (s *ExampleService) SayHello(ctx context.Context, req *example.SayHelloRequest) (*example.SayHelloResponse, error) {
	resp := &example.SayHelloResponse{
		Msg: "hello, " + req.User,
	}
	return resp, nil
}
