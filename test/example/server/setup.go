package server

import (
	"net"

	"github.com/MasterJoyHunan/genrpc/test/example/proto/myrpc"

	"google.golang.org/grpc"
)

func Setup() *grpc.Server {
	rpcServer := grpc.NewServer()
	myrpc.RegistermyrpcServer(rpcServer, &myrpcServer{})
	listen, err := net.Listen("tcp", "127.0.0.1:8877")
	if err != nil {
		panic(err)
	}
	go func() {
		rpcServer.Serve(listen)
	}()
	return rpcServer
}
