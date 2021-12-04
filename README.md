# grpc-http-example

## 编译PB文件

使用gRPC和[gRPC-Gateway](https://github.com/grpc-ecosystem/grpc-gateway) 实现一套代码同时提供 RPC服务和 HTTP服务

PB文件的编译请参考：https://github.com/grpc-ecosystem/grpc-gateway#usage

1. 编译pb文件

```
protoc \
-I=. \
-I=$GOROOT/src \
--go_out=. \
--go_opt=paths=source_relative \
--go-grpc_out=. \
--go-grpc_opt=paths=source_relative \
*.proto
```
2. 编译gw文件
```
protoc \
-I . \
-I $GOROOT/src \
--grpc-gateway_out . \
--grpc-gateway_opt logtostderr=true \
--grpc-gateway_opt paths=source_relative \
--grpc-gateway_opt generate_unbound_methods=true \
*.proto
```

## 运行结果

```shell
# curl -X POST -k http://localhost:8080/v1/SayHello -d '{"User": "SongLee24"}'
{"Msg":"hello, SongLee24"}
```

