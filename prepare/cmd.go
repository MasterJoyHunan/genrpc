package prepare

import (
	"github.com/spf13/cobra"
	"os"
)

var (
	GrpcOutDir string
	ProtoFile  string

	rootCmd = &cobra.Command{
		Use:   "genrpc",
		Short: "生成 GRPC 的项目结构",
		Args:  cobra.ArbitraryArgs,
		Run: func(cmd *cobra.Command, args []string) {
			ProtoFile = args[0]
		},
	}
)

func init() {
	rootCmd.Flags().StringVar(&GrpcOutDir, "dir", ".", "-go_out 和 -go-grpc_out 的参数")

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
