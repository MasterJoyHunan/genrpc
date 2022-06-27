package generator

import (
	"fmt"
	"strings"

	. "github.com/MasterJoyHunan/genrpc/prepare"
	"github.com/MasterJoyHunan/genrpc/tpl"

	"github.com/zeromicro/go-zero/tools/goctl/util/format"
)

const (
	defaultHost = "127.0.0.1"
	defaultPort = 8877
	devModel    = "local|dev|prod"
)

func GenEtc() error {
	filename, err := format.FileNamingFormat("go_zero", GrpcProto.Package.Name)
	if err != nil {
		return err
	}

	mode := strings.Split(devModel, "|")

	for _, m := range mode {
		err = genFile(fileGenConfig{
			dir:             GrpcOutDir,
			subDir:          etcDir,
			filename:        fmt.Sprintf("%s-%s.yaml", filename, m),
			templateName:    "etcTemplate",
			builtinTemplate: tpl.EtcTemplate,
			data: map[string]interface{}{
				"serviceName": filename,
				"host":        defaultHost,
				"port":        defaultPort,
			},
		})
		if err != nil {
			return err
		}
	}

	return nil
}
