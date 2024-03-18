package generator

import (
	. "github.com/MasterJoyHunan/genrpc/prepare"
	"github.com/MasterJoyHunan/genrpc/tpl"
)

func GenSvc() error {
	return genFile(fileGenConfig{
		dir:             GrpcOutDir,
		subDir:          "svc",
		filename:        "grpc_context.go",
		templateName:    "grpcContextTemplate",
		builtinTemplate: tpl.GrpcContext,
	})
}
