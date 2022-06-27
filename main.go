package main

import (
	"github.com/MasterJoyHunan/genrpc/generator"
	"github.com/MasterJoyHunan/genrpc/prepare"
)

func main() {
	prepare.Setup()
	Must(generator.GenEtc())
	Must(generator.GenConfig())
	Must(generator.GenMain())
	Must(generator.GenServer())
	Must(generator.GenLogic())

}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}
