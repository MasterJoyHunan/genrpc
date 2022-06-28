package main

import (
	"github.com/MasterJoyHunan/genrpc/generator"
	"github.com/MasterJoyHunan/genrpc/prepare"
	"regexp"
	"strings"
)

func main() {
	prepare.Setup()

	if IfNeedGenerate("pb") {
		Must(generator.GenPB())
	}
	if IfNeedGenerate("etc") {
		Must(generator.GenEtc())
	}
	if IfNeedGenerate("config") {
		Must(generator.GenConfig())
	}
	if IfNeedGenerate("main") {
		Must(generator.GenMain())
	}
	if IfNeedGenerate("server") {
		Must(generator.GenServer())
	}
	if IfNeedGenerate("logic") {
		Must(generator.GenLogic())
	}
}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func IfNeedGenerate(target string) bool {
	spaceRe, _ := regexp.Compile(`,|，`)
	if len(prepare.Only) > 0 {
		only := spaceRe.Split(prepare.Only, -1)
		for _, s := range only {
			if strings.ToLower(s) == target {
				return true
			}
		}
		return false
	}

	if len(prepare.Without) > 0 {
		without := spaceRe.Split(prepare.Without, -1)
		for _, s := range without {
			if strings.ToLower(s) == target {
				return false
			}
		}
		return true
	}
	return true
}
