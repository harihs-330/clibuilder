package toolregistry

import (
	"clibuilder/core"
	"clibuilder/tools/binrunner"
	"clibuilder/tools/filerunner"
	"clibuilder/tools/installer"
	"clibuilder/tools/plugincli"
	"clibuilder/tools/repocli"
)

func GetTools() []core.Tool {
	return []core.Tool{
		binrunner.New(),
		filerunner.New(),
		plugincli.New(),
		repocli.New(),
		installer.New(),
	}
}
