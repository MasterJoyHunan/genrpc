package generator

import (
	"github.com/MasterJoyHunan/genrpc/tpl"
)

func GenClient() error {
	return GenFile(
		"grpc_client.go",
		tpl.RpcClientTemplate,
		WithSubDir("pkg/grpc_client"),
	)
}
