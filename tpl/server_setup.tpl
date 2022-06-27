package server

import (
    "net"

    {{.importPackages}}

    "google.golang.org/grpc"
)

func Setup() *grpc.Server {
    rpcServer := grpc.NewServer()
    {{.pbPkg}}.Register{{.serverName}}Server(rpcServer, &{{.serverName}}Server{})
    listen, err := net.Listen("tcp", "{{.host}}:{{.port}}")
    if err != nil {
        panic(err)
    }
    go func() {
        rpcServer.Serve(listen)
    }()
    return rpcServer
}
