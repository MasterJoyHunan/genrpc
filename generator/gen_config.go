package generator

import (
	. "github.com/MasterJoyHunan/genrpc/prepare"
	"github.com/MasterJoyHunan/genrpc/tpl"
)

func GenConfig() error {
	return genFile(fileGenConfig{
		dir:             GrpcOutDir,
		subDir:          configDir,
		filename:        configPacket + ".go",
		templateName:    "configTemplate",
		builtinTemplate: tpl.ConfigTemplate,
		data: map[string]interface{}{
			"pkgName": configPacket,
		},
	})
}
