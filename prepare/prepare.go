package prepare

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/emicklei/proto"
	"github.com/zeromicro/go-zero/tools/goctl/util/ctx"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

var (
	RootPkg   string
	GrpcProto Proto
)

type Proto struct {
	Package   *proto.Package
	GoPackage string
	Import    *proto.Import
	Message   []*proto.Message
	Service   *proto.Service
}

func Setup() {
	abs, err := filepath.Abs(ProtoFile)
	r, err := os.Open(abs)
	if err != nil {
		fmt.Println("无法找到文件")
		os.Exit(1)
	}
	parser := proto.NewParser(r)
	GrpcParse, err := parser.Parse()
	if err != nil {
		fmt.Println("无法解析proto文件")
		os.Exit(1)
	}

	proto.Walk(
		GrpcParse,
		proto.WithImport(func(i *proto.Import) {
			fmt.Println("proto.Import")
			GrpcProto.Import = i
		}),
		proto.WithMessage(func(message *proto.Message) {
			fmt.Println("proto.Message")
			GrpcProto.Message = append(GrpcProto.Message, message)
		}),
		proto.WithPackage(func(p *proto.Package) {
			fmt.Println("proto.Package")
			GrpcProto.Package = p
		}),
		proto.WithService(func(service *proto.Service) {
			fmt.Println("proto.Service")
			GrpcProto.Service = service
		}),
		proto.WithOption(func(option *proto.Option) {
			fmt.Println("proto.Option")
			if option.Name == "go_package" {
				GrpcProto.GoPackage = option.Constant.Source
			}
		}),
	)

	RootPkg, err = GetParentPackage(GrpcOutDir)
	if err != nil {
		panic(err)
	}
}

func GetParentPackage(dir string) (string, error) {
	abs, err := filepath.Abs(dir)
	if err != nil {
		return "", err
	}

	projectCtx, err := ctx.Prepare(abs)
	if err != nil {
		return "", err
	}

	// fix https://github.com/zeromicro/go-zero/issues/1058
	wd := projectCtx.WorkDir
	d := projectCtx.Dir
	same, err := pathx.SameFile(wd, d)
	if err != nil {
		return "", err
	}

	trim := strings.TrimPrefix(projectCtx.WorkDir, projectCtx.Dir)
	if same {
		trim = strings.TrimPrefix(strings.ToLower(projectCtx.WorkDir), strings.ToLower(projectCtx.Dir))
	}

	return filepath.ToSlash(filepath.Join(projectCtx.Path, trim)), nil
}
