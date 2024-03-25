package server

import (
    "context"

    "{{.rootPkg}}/svc"
    "{{.pbPkg}}"
    {{.logicPkgAlias}} "{{.logicPkg}}"
)

type {{.serverName}}Server struct {
    {{.pbLastPkg}}.Unimplemented{{.serverName}}Server
}

{{ range .funcArr -}}
{{if .comment}}// {{.funcName}} {{ .comment}}{{end}}
func (s *{{$.serverName}}Server) {{.funcName}} (ctx context.Context, req *{{$.pbLastPkg}}.{{.request}}) (*{{$.pbLastPkg}}.{{.response}}, error) {
    return {{$.logicPkgAlias}}.{{.funcName}}(svc.NewGrpcContext(ctx), req)
}

{{ end -}}
