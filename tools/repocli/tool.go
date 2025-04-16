package repocli

import (
	"clibuilder/tools/repocli/cmd"
	"fmt"
)

type Repocli struct{}

func (b Repocli) Name() string {
	return "repocli"
}

func (b Repocli) Run() {
	fmt.Println("🧪 Running Repocli tool!")
	cmd.Execute()

}

func New() Repocli {
	return Repocli{}
}
