package generator

import (
	"github.com/MasterJoyHunan/genrpc/prepare"
	"github.com/MasterJoyHunan/genrpc/tpl"

	"github.com/zeromicro/go-zero/tools/goctl/util/format"
)

func GenMain() error {
	serverName, err := format.FileNamingFormat("go_zero", prepare.GrpcProto.Package.Name)
	if err != nil {
		return err
	}

	return GenFile(
		serverName+".go",
		tpl.MainTemplate,
		WithData(map[string]any{
			"rootPkg":    prepare.RootPkg,
			"serverName": serverName,
		}),
	)
}
