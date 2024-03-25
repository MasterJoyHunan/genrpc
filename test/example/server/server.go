package server

import (
	"context"

	myrpcLogic "github.com/MasterJoyHunan/genrpc/test/example/logic/myrpc"
	"github.com/MasterJoyHunan/genrpc/test/example/proto/foo/bar"
	"github.com/MasterJoyHunan/genrpc/test/example/svc"
)

type MyrpcServer struct {
	bar.UnimplementedMyrpcServer
}

func (s *MyrpcServer) Ping(ctx context.Context, req *bar.Request) (*bar.Response, error) {
	return myrpcLogic.Ping(svc.NewGrpcContext(ctx), req)
}

// Pong  pong comment
func (s *MyrpcServer) Pong(ctx context.Context, req *bar.Request) (*bar.Response, error) {
	return myrpcLogic.Pong(svc.NewGrpcContext(ctx), req)
}
