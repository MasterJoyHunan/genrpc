package myrpc

import (
	"github.com/MasterJoyHunan/genrpc/test/example/proto/foo/bar"
	"github.com/MasterJoyHunan/genrpc/test/example/svc"
)

func Pong(ctx *svc.GrpcContext, req *bar.Request) (reps *bar.Response, err error) {
	// todo: add your logic here and delete this line

	return reps, nil
}
