package binrunner

import (
	"clibuilder/tools/binrunner/cmd"
	"fmt"
)

type Binrunner struct{}

func (b Binrunner) Name() string {
	return "binrunner"
}

func (b Binrunner) Run() {
	fmt.Println("🧪 Running Binrunner tool!")
	cmd.Execute()
}

func New() Binrunner {
	return Binrunner{}
}
