package generator

import (
	"fmt"
	"os"
	"path"
	"sort"
	"strings"

	. "github.com/MasterJoyHunan/genrpc/prepare"
	"github.com/MasterJoyHunan/genrpc/tpl"
	"github.com/emicklei/proto"
	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"github.com/zeromicro/go-zero/tools/goctl/util/format"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

func GenServer() error {
	pbDir := path.Join(RootPkg, GrpcProto.GoPackage)
	pbPkg := path.Base(pbDir)
	filename := "server.go"

	os.Remove(path.Join(GrpcOutDir, serverDir, filename))

	serverFmtName, err := format.FileNamingFormat(dirFmt, GrpcProto.Service.Name)
	logicPath := pathx.JoinPackages(RootPkg, logicPacket, serverFmtName)
	logicPkgAlias := logicPath[strings.LastIndex(logicPath, "/")+1:] + "Logic"

	if err != nil {
		return err
	}

	err = genFile(fileGenConfig{
		dir:             GrpcOutDir,
		subDir:          serverDir,
		filename:        filename,
		templateName:    "serverTemplate",
		builtinTemplate: tpl.ServerTemplate,
		data: map[string]interface{}{
			"importPackages": genServerImport(pbDir, logicPath, logicPkgAlias),
			"pbPkg":          pbPkg,
			"serverName":     util.Title(GrpcProto.Service.Name),
			"func":           genFunc(pbPkg, logicPkgAlias),
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
	return %s.%s(svc.NewGrpcContext(ctx), req)
}`,
			util.Title(GrpcProto.Service.Name),
			util.Title(rpc.Name),
			pbPkg+"."+util.Title(rpc.RequestType),
			pbPkg+"."+util.Title(rpc.ReturnsType),
			logicPkg,
			util.Title(rpc.Name),
		))
	}
	return sb.String()
}

func genServerImport(pbDir, logicDir, logicPkgAlias string) string {
	importSet := collection.NewSet()

	importSet.AddStr(fmt.Sprintf("\"%s/svc\"", RootPkg))
	// pb pkg
	importSet.AddStr(fmt.Sprintf("\"%s\"", pbDir))
	// logic pkg
	importSet.AddStr(fmt.Sprintf("%s \"%s\"", logicPkgAlias, logicDir))

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
			"serverName":     util.Title(GrpcProto.Service.Name),
		},
	})
}
