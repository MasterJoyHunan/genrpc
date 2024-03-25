package main

import (
	"flag"

	"{{.rootPkg}}/server"
)

var release string

func init() {
	flag.StringVar(&release, "release", "local", "release model, optional local/dev/prod")
}

func main() {
	flag.Parse()

	// configFile := fmt.Sprintf("etc/{{.serverName}}-%s.yaml", release)

	server.Setup()
}
