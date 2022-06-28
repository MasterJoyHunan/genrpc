package generator

import (
	"fmt"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
	"path"
	"strings"

	. "github.com/MasterJoyHunan/genrpc/prepare"
	"github.com/MasterJoyHunan/genrpc/tpl"

	"github.com/emicklei/proto"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"github.com/zeromicro/go-zero/tools/goctl/util/format"
)

func GenLogic() error {
	for _, e := range GrpcProto.Service.Elements {
		rpc := e.(*proto.RPC)
		err := genLogicByRpc(rpc)
		if err != nil {
			return err
		}
	}
	return nil
}

func genLogicByRpc(rpc *proto.RPC) error {
	logic := rpc.Name + "Logic"
	filename, err := format.FileNamingFormat("go_zero", logic)
	if err != nil {
		return err
	}

	pbDir := path.Join(RootPkg, GrpcProto.GoPackage)
	pbPkg := path.Base(pbDir)

	fmtName, err := format.FileNamingFormat(dirFmt, GrpcProto.Service.Name)
	if err != nil {
		return err
	}
	dirPath := pathx.JoinPackages(logicDir, fmtName)

	return genFile(fileGenConfig{
		dir:             GrpcOutDir,
		subDir:          dirPath,
		filename:        filename + ".go",
		templateName:    "logicTemplate",
		builtinTemplate: tpl.LogicTemplate,
		data: map[string]interface{}{
			"pkgName":  fmtName[strings.LastIndex(fmtName, "/")+1:],
			"imports":  fmt.Sprintf("\"%s\"", pbDir),
			"function": util.Title(strings.TrimSuffix(logic, "Logic")),
			"request":  fmt.Sprintf("%s.%s", pbPkg, util.Title(rpc.RequestType)),
			"response": fmt.Sprintf("%s.%s", pbPkg, util.Title(rpc.ReturnsType)),
		},
	})
}
