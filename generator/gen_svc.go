package generator

import (
	"github.com/MasterJoyHunan/genrpc/tpl"
)

func GenSvc() error {
	return GenFile(
		"grpc_context.go",
		tpl.GrpcContext,
		WithSubDir("svc"),
	)
}
