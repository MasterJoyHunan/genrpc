package prepare

import (
	"github.com/spf13/cobra"
)

var (
	GrpcOutDir string
	ProtoFile  string
	Without    string
	Only       string

	rootCmd = &cobra.Command{
		Use:     "genrpc",
		Short:   "生成 GRPC 的项目结构",
		Example: "  genrpc --dir=. --only=server,logic xxx.proto",
		Args:    cobra.ExactValidArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ProtoFile = args[0]
			return nil
		},
	}
)

func init() {
	rootCmd.Flags().StringVar(&GrpcOutDir, "dir", ".", "-go_out 和 -go-grpc_out 的参数")
	rootCmd.Flags().StringVar(&Without, "without", "", "不生成的目录,多个以英文逗号分割,可选参数:pb,etc,config,main,server,logic")
	rootCmd.Flags().StringVar(&Only, "only", "", "当次只生成的目录,多个以英文逗号分割,优先级比 without 参数高,可选参数:pb,etc,config,main,server,logic")
}
