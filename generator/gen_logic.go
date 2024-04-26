package generator

import (
	"path"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/MasterJoyHunan/genrpc/prepare"
	"github.com/MasterJoyHunan/genrpc/tpl"

	"github.com/emicklei/proto"
	"github.com/zeromicro/go-zero/tools/goctl/util/format"
)

func GenLogic() error {
	for _, e := range prepare.GrpcProto.Service.Elements {

		rpc := e.(*proto.RPC)

		filename, err := format.FileNamingFormat("go_zero", rpc.Name+"Logic")
		if err != nil {
			return err
		}

		serverName, err := format.FileNamingFormat(dirFmt, prepare.GrpcProto.Service.Name)
		if err != nil {
			return err
		}

		pbPkg := path.Join(prepare.RootPkg, prepare.GrpcProto.GoPackage)

		err = GenFile(
			filename+".go",
			tpl.LogicTemplate,
			WithSubDir(path.Join("logic", serverName)),
			WithData(map[string]string{
				"pkgName":   path.Base(serverName),
				"rootPkg":   prepare.RootPkg,
				"pbPkg":     pbPkg,
				"pbLastPkg": path.Base(pbPkg),
				"funcName":  cases.Title(language.English, cases.NoLower).String(rpc.Name),
				"request":   cases.Title(language.English, cases.NoLower).String(rpc.RequestType),
				"response":  cases.Title(language.English, cases.NoLower).String(rpc.ReturnsType),
			}),
		)

		if err != nil {
			return err
		}
	}
	return nil
}
