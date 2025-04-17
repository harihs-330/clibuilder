package installer

import (
	"clibuilder/tools/installer/cmd"
	"fmt"
)

type Installer struct{}

func (i Installer) Name() string {
	return "installer"
}

func (i Installer) Run() {
	fmt.Println("🛠️ Running Installer tool!")
	cmd.Execute()
}

func New() Installer {
	return Installer{}
}
