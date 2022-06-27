package server

import (
    "context"

    {{.importPackages}}
)

type {{.serverName}}Server struct {
    {{.pbPkg}}.Unimplemented{{.serverName}}Server
}

{{.func}}