package generator

import (
	"fmt"
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

	return genFile(fileGenConfig{
		dir:             GrpcOutDir,
		subDir:          logicDir,
		filename:        filename + ".go",
		templateName:    "logicTemplate",
		builtinTemplate: tpl.LogicTemplate,
		data: map[string]interface{}{
			"imports":  fmt.Sprintf("\"%s\"", pbDir),
			"function": util.Title(strings.TrimSuffix(logic, "Logic")),
			"request":  fmt.Sprintf("%s.%s", pbPkg, rpc.RequestType),
			"response": fmt.Sprintf("%s.%s", pbPkg, rpc.ReturnsType),
		},
	})
}
