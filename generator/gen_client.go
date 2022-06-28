package generator

import (
	. "github.com/MasterJoyHunan/genrpc/prepare"
	"github.com/MasterJoyHunan/genrpc/tpl"
)

func GenClient() error {

	return genFile(fileGenConfig{
		dir:             GrpcOutDir,
		subDir:          "rpcclient",
		filename:        "rpcclient.go",
		templateName:    "rpcclientTemplate",
		builtinTemplate: tpl.RpcClientTemplate,
	})
}
