package server

import (
	"context"

	"github.com/MasterJoyHunan/genrpc/test/example/logic"
	"github.com/MasterJoyHunan/genrpc/test/example/proto/myrpc"
)

type MyrpcServer struct {
	myrpc.UnimplementedMyrpcServer
}

func (s *MyrpcServer) Ping(ctx context.Context, req *myrpc.Request) (*myrpc.Response, error) {
	return logic.Ping(req)
}
