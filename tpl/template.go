package tpl

import _ "embed"

var (
	//go:embed etc.tpl
	EtcTemplate string

	//go:embed config.tpl
	ConfigTemplate string

	//go:embed main.tpl
	MainTemplate string

	//go:embed logic.tpl
	LogicTemplate string

	//go:embed server.tpl
	ServerTemplate string

	//go:embed server_setup.tpl
	ServerSetupTemplate string
)
