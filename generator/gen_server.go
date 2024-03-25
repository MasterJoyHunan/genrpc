package generator

import (
	"os"
	"path"

	"github.com/MasterJoyHunan/genrpc/prepare"
	"github.com/MasterJoyHunan/genrpc/tpl"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/emicklei/proto"
	"github.com/zeromicro/go-zero/tools/goctl/util/format"
)

func GenServer() error {
	// 每次都重新生成
	os.Remove(path.Join(prepare.GrpcOutDir, "server", "server.go"))

	pbPkg := path.Join(prepare.RootPkg, prepare.GrpcProto.GoPackage)

	serverFmtName, err := format.FileNamingFormat(dirFmt, prepare.GrpcProto.Service.Name)
	logicPkg := path.Join(prepare.RootPkg, "logic", serverFmtName)
	logicPkgAlias := path.Base(logicPkg) + "Logic"

	if err != nil {
		return err
	}

	err = GenFile(
		"server.go",
		tpl.ServerTemplate,
		WithSubDir("server"),
		WithData(map[string]any{
			"rootPkg":       prepare.RootPkg,
			"pbPkg":         pbPkg,
			"pbLastPkg":     path.Base(pbPkg),
			"logicPkg":      logicPkg,
			"logicPkgAlias": logicPkgAlias,
			"serverName":    cases.Title(language.English).String(prepare.GrpcProto.Service.Name),
			"funcArr":       genFunc(),
		}),
	)
	if err != nil {
		return err
	}
	return GenFile(
		"setup.go",
		tpl.ServerSetupTemplate,
		WithSubDir("server"),
		WithData(map[string]any{
			"pbPkg":      pbPkg,
			"pbLastPkg":  path.Base(pbPkg),
			"host":       defaultHost,
			"port":       defaultPort,
			"serverName": cases.Title(language.English).String(prepare.GrpcProto.Service.Name),
		}),
	)
}

func genFunc() (arr []map[string]string) {
	for _, e := range prepare.GrpcProto.Service.Elements {

		rpc := e.(*proto.RPC)

		arr = append(arr, map[string]string{
			"comment":  rpc.Comment.Message(),
			"funcName": cases.Title(language.English).String(rpc.Name),
			"request":  cases.Title(language.English).String(rpc.RequestType),
			"response": cases.Title(language.English).String(rpc.ReturnsType),
		})
	}
	return
}
