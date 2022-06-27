package main

import (
	"flag"

	"github.com/MasterJoyHunan/genrpc/test/example/server"
)

var release string

func init() {
	flag.StringVar(&release, "release", "local", "release model, optional local/dev/prod")
}

func main() {
	flag.Parse()

	// configFile := fmt.Sprintf("etc/myrpc-%s.yaml", release)

	server.Setup()
}
