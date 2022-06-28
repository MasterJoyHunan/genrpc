package generator

import (
	"fmt"
	"sort"
	"strings"

	. "github.com/MasterJoyHunan/genrpc/prepare"
	"github.com/MasterJoyHunan/genrpc/tpl"

	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/tools/goctl/util/format"
)

func GenMain() error {
	filename, err := format.FileNamingFormat("go_zero", GrpcProto.Package.Name)
	if err != nil {
		return err
	}

	return genFile(fileGenConfig{
		dir:             GrpcOutDir,
		subDir:          "",
		filename:        filename + ".go",
		templateName:    "mainTemplate",
		builtinTemplate: tpl.MainTemplate,
		data: map[string]interface{}{
			"importPkg":  genMainImportPkg(),
			"etcDir":     etcDir,
			"configName": filename,
			"setup":      genMainSetup(),
			"host":       defaultHost,
			"port":       defaultPort,
		},
	})
}

func genMainImportPkg() string {
	set := collection.NewSet()

	// server
	set.AddStr(fmt.Sprintf("\"%s/%s\"", RootPkg, serverPacket))

	importArr := set.KeysStr()
	sort.Strings(importArr)
	return strings.Join(importArr, "\n\t")
}

func genMainSetup() string {
	return fmt.Sprintf(`%s.Setup()`, serverPacket)
}
