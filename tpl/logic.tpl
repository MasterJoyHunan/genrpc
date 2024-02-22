package {{.pkgName}}

import (
	{{.imports}}
)

func {{.function}}(ctx *svc.GrpcContext, req *{{.request}}) (reps *{{.response}}, err error) {
	// todo: add your logic here and delete this line

	return reps, nil
}
