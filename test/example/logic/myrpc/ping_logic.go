package myrpc

import (
	"github.com/MasterJoyHunan/genrpc/test/example/proto/myrpc"
	"github.com/MasterJoyHunan/genrpc/test/example/svc"
)

func Ping(ctx *svc.GrpcContext, req *myrpc.Request) (reps *myrpc.Response, err error) {
	// todo: add your logic here and delete this line

	return reps, nil
}
