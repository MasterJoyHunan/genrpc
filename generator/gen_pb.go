package generator

import (
	"fmt"
	"github.com/zeromicro/go-zero/tools/goctl/rpc/execx"
	"os"
	"path/filepath"

	"github.com/MasterJoyHunan/genrpc/prepare"
)

func GenPB() error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}
	cmdStr := fmt.Sprintf("protoc -I=%s %s --go_out=%s --go-grpc_out=%s", filepath.Dir(prepare.ProtoFile), filepath.Base(prepare.ProtoFile), prepare.GrpcOutDir, prepare.GrpcOutDir)
	fmt.Println(cmdStr)

	//cmd := exec.Command("protoc", "-I="+filepath.Dir(prepare.ProtoFile), filepath.Base(prepare.ProtoFile), "--go_out="+prepare.GrpcOutDir, "--go-grpc_out="+prepare.GrpcOutDir)
	//cmd.Stdout = os.Stdout
	//cmd.Stderr = os.Stderr
	//cmd.Dir = pwd

	run, err := execx.Run(cmdStr, pwd)
	fmt.Println(run)
	return err
}
