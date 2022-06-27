package generator

import (
	"fmt"
	. "github.com/MasterJoyHunan/genrpc/prepare"
	"github.com/MasterJoyHunan/genrpc/tpl"
	"github.com/emicklei/proto"
	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/tools/goctl/util/format"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
	"path"
	"sort"
	"strings"
)

func GenServer() error {
	pbDir := path.Join(RootPkg, GrpcProto.GoPackage)
	pbPkg := path.Base(pbDir)

	fmtName, err := format.FileNamingFormat(dirFmt, GrpcProto.Service.Name)
	if err != nil {
		return err
	}

	err = genFile(fileGenConfig{
		dir:             GrpcOutDir,
		subDir:          serverDir,
		filename:        "server.go",
		templateName:    "serverTemplate",
		builtinTemplate: tpl.ServerTemplate,
		data: map[string]interface{}{
			"importPackages": genServerImport(pbDir, fmtName),
			"pbPkg":          pbPkg,
			"serverName":     GrpcProto.Service.Name,
			"func":           genFunc(pbPkg, fmtName[strings.LastIndex(fmtName, "/")+1:]),
		},
	})
	if err != nil {
		return err
	}
	return genSetup(pbDir, pbPkg)
}

func genFunc(pbPkg, logicPkg string) string {
	var sb strings.Builder
	for _, e := range GrpcProto.Service.Elements {
		rpc := e.(*proto.RPC)
		sb.WriteString(fmt.Sprintf(`
func (s *%sServer) %s (ctx context.Context, req *%s) (*%s, error) {
	return %s.%s(req)
}`,
			GrpcProto.Service.Name,
			rpc.Name,
			pbPkg+"."+rpc.RequestType,
			pbPkg+"."+rpc.ReturnsType,
			logicPkg,
			rpc.Name,
		))
	}
	return sb.String()
}

func genServerImport(pbDir, logicPath string) string {
	importSet := collection.NewSet()
	// pb pkg
	importSet.AddStr(fmt.Sprintf("\"%s\"", pbDir))
	// logic pkg
	logicDir := pathx.JoinPackages(RootPkg, logicPath)
	importSet.AddStr(fmt.Sprintf("\"%s\"", logicDir))

	imports := importSet.KeysStr()
	sort.Strings(imports)

	return strings.Join(imports, "\n\t")
}

func genSetup(pbDir, pbPkg string) error {
	return genFile(fileGenConfig{
		dir:             GrpcOutDir,
		subDir:          serverDir,
		filename:        "setup.go",
		templateName:    "serverSetupTemplate",
		builtinTemplate: tpl.ServerSetupTemplate,
		data: map[string]interface{}{
			"importPackages": fmt.Sprintf("\"%s\"", pbDir),
			"pbPkg":          pbPkg,
			"host":           defaultHost,
			"port":           defaultPort,
			"serverName":     GrpcProto.Service.Name,
		},
	})
}
