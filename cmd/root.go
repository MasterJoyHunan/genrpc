package cmd

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/MasterJoyHunan/genrpc/generator"
	"github.com/MasterJoyHunan/genrpc/prepare"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:     "genrpc",
		Short:   "生成 GRPC 的项目结构",
		Example: "genrpc --dir=. --only=server,logic xxx.proto",
		Args:    cobra.ExactValidArgs(1),
		RunE:    GenGrpcCode,
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
	rootCmd.Flags().StringVar(&prepare.Without, "without", "", "不生成的目录,多个以英文逗号分割,可选参数:pb,etc,config,main,server,logic,svc,client")
	rootCmd.Flags().StringVar(&prepare.Only, "only", "", "当次只生成的目录,多个以英文逗号分割,优先级比 without 参数高,可选参数:pb,etc,config,main,server,logic,svc,client")
}

func GenGrpcCode(cmd *cobra.Command, args []string) error {
	prepare.ProtoFile = args[0]
	prepare.Setup()

	if IfNeedGenerate("pb") {
		if err := generator.GenPB(); err != nil {
			return err
		}
	}
	if IfNeedGenerate("etc") {
		if err := generator.GenEtc(); err != nil {
			return err
		}
	}
	if IfNeedGenerate("config") {
		if err := generator.GenConfig(); err != nil {
			return err
		}
	}
	if IfNeedGenerate("main") {
		if err := generator.GenMain(); err != nil {
			return err
		}
	}
	if IfNeedGenerate("server") {
		if err := generator.GenServer(); err != nil {
			return err
		}
	}
	if IfNeedGenerate("logic") {
		if err := generator.GenLogic(); err != nil {
			return err
		}
	}
	if IfNeedGenerate("svc") {
		if err := generator.GenSvc(); err != nil {
			return err
		}
	}
	if IfNeedGenerate("client") {
		if err := generator.GenClient(); err != nil {
			return err
		}
	}
	return nil
}

func IfNeedGenerate(target string) bool {
	spaceRe, _ := regexp.Compile(`,|，`)
	if len(prepare.Only) > 0 {
		only := spaceRe.Split(prepare.Only, -1)
		for _, s := range only {
			if strings.ToLower(s) == target {
				return true
			}
		}
		return false
	}

	if len(prepare.Without) > 0 {
		without := spaceRe.Split(prepare.Without, -1)
		for _, s := range without {
			if strings.ToLower(s) == target {
				return false
			}
		}
		return true
	}
	return true
}
