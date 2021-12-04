package service

import (
	"google.golang.org/grpc"
	"grpc-http-example/service/example"
)

type ServiceRegisterFunc func(*grpc.Server)

var (
	registerFuncs []ServiceRegisterFunc
)

func init() {
	registerFuncs = append(registerFuncs,
		example.GetRegisterFunc(),
		// append as follows
	)
}

func GetServiceRegisterFuncs() []ServiceRegisterFunc {
	return registerFuncs
}
