package test

import (
	"testing"

	"github.com/MasterJoyHunan/genrpc/generator"
	"github.com/MasterJoyHunan/genrpc/prepare"
)

func TestMain(m *testing.M) {
	prepare.ProtoFile = "../tpl/myrpc.proto"
	prepare.GrpcOutDir = "example"

	prepare.Setup()

	m.Run()
}

func TestGenPB(t *testing.T) {
	if err := generator.GenPB(); err != nil {
		t.Failed()
	}
}

func TestGenEtc(t *testing.T) {
	if err := generator.GenEtc(); err != nil {
		t.Failed()
	}
}

func TestGenConfig(t *testing.T) {
	if err := generator.GenConfig(); err != nil {
		t.Failed()
	}
}

func TestGenMain(t *testing.T) {
	if err := generator.GenMain(); err != nil {
		t.Failed()
	}
}

func TestGenServer(t *testing.T) {
	if err := generator.GenServer(); err != nil {
		t.Failed()
	}
}

func TestGenLogic(t *testing.T) {
	if err := generator.GenLogic(); err != nil {
		t.Failed()
	}
}

func TestGenClient(t *testing.T) {
	if err := generator.GenClient(); err != nil {
		t.Failed()
	}
}

func TestGenSvc(t *testing.T) {
	if err := generator.GenSvc(); err != nil {
		t.Failed()
	}
}
