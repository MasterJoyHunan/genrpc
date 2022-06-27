package main

import (
	"flag"

	{{.importPkg}}
)

var release string

func init() {
	flag.StringVar(&release, "release", "local", "release model, optional local/dev/prod")
}

func main() {
	flag.Parse()

	// configFile := fmt.Sprintf("{{.etcDir}}/{{.configName}}-%s.yaml", release)

	{{.setup}}
}
