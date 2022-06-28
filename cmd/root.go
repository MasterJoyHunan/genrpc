package cmd

import (
	"fmt"
	"github.com/MasterJoyHunan/genrpc/prepare"
	"github.com/spf13/cobra"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Use:   "genrpc",
		Short: "生成 GRPC 的项目结构",
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVar(&prepare.GrpcOutDir, "dir", ".", "-go_out 和 -go-grpc_out 的参数")
	rootCmd.AddCommand(serverCmd)
	rootCmd.AddCommand(clientCmd)
}
