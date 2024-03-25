package generator

import (
	"github.com/MasterJoyHunan/genrpc/tpl"
)

func GenConfig() error {
	return GenFile(
		"config.go",
		tpl.ConfigTemplate,
		WithSubDir("config"),
	)
}
