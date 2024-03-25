package generator

import (
	"fmt"
	"strings"

	"github.com/MasterJoyHunan/genrpc/prepare"
	"github.com/MasterJoyHunan/genrpc/tpl"

	"github.com/zeromicro/go-zero/tools/goctl/util/format"
)

const (
	defaultHost = "127.0.0.1"
	defaultPort = 8887
	devModel    = "local|dev|prod"
)

func GenEtc() error {
	serverName, err := format.FileNamingFormat("go_zero", prepare.GrpcProto.Package.Name)
	if err != nil {
		return err
	}

	mode := strings.Split(devModel, "|")

	for _, m := range mode {
		err = GenFile(
			fmt.Sprintf("%s-%s.yaml", serverName, m),
			tpl.EtcTemplate,
			WithSubDir("etc"),
			WithData(map[string]any{
				"serverName": serverName,
				"host":       defaultHost,
				"port":       defaultPort,
			}),
		)
		if err != nil {
			return err
		}
	}

	return nil
}
