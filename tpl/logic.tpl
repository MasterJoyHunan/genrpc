package {{.pkgName}}

import (
	"{{.rootPkg}}/svc"
	"{{.pbPkg}}"
)

func {{.funcName}}(ctx *svc.GrpcContext, req *{{.pbLastPkg}}.{{.request}}) (resp *{{.pbLastPkg}}.{{.response}}, err error) {
	// todo: add your logic here and delete this line

	return resp, nil
}
