package cmd

import (
	"github.com/MasterJoyHunan/genrpc/generator"
	"github.com/MasterJoyHunan/genrpc/prepare"
	"github.com/spf13/cobra"
)

var clientCmd = &cobra.Command{
	Use:     "client",
	Short:   "生成 GRPC client 的文件",
	Example: "  genrpc client xxx.proto",
	Args:    cobra.ExactValidArgs(1),
	RunE:    GenRpcClient,
}

func GenRpcClient(cmd *cobra.Command, args []string) error {
	prepare.ProtoFile = args[0]
	prepare.Setup()
	return generator.GenClient()
}
