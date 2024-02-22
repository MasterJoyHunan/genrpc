package generator

import (
	. "github.com/MasterJoyHunan/genrpc/prepare"
	"github.com/MasterJoyHunan/genrpc/tpl"

	"github.com/zeromicro/go-zero/tools/goctl/util/format"
)

func GenSvc() error {
	filename, err := format.FileNamingFormat("go_zero", GrpcProto.Package.Name)
	if err != nil {
		return err
	}

	return genFile(fileGenConfig{
		dir:             GrpcOutDir,
		subDir:          "svc",
		filename:        filename + ".go",
		templateName:    "grpcContextTemplate",
		builtinTemplate: tpl.GrpcContext,
	})
}
