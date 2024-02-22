package server

import (
	"context"

	myrpcLogic "github.com/MasterJoyHunan/genrpc/test/example/logic/myrpc"
	"github.com/MasterJoyHunan/genrpc/test/example/proto/myrpc"
	"github.com/MasterJoyHunan/genrpc/test/example/svc"
)

type MyrpcServer struct {
	myrpc.UnimplementedMyrpcServer
}

func (s *MyrpcServer) Ping(ctx context.Context, req *myrpc.Request) (*myrpc.Response, error) {
	return myrpcLogic.Ping(svc.NewGrpcContext(ctx), req)
}
