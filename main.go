package main

import (
	"grpc-http-example/http"
	"grpc-http-example/rpc"
)

func main() {
	go rpc.Serve()
	http.Serve()
}
