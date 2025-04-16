package plugincli

import (
	"clibuilder/tools/plugincli/cmd"
	"fmt"
)

type Plugincli struct{}

func (b Plugincli) Name() string {
	return "plugincli"
}

func (b Plugincli) Run() {
	fmt.Println("🧪 Running Plugincli tool!")
	cmd.Execute()
}

func New() Plugincli {
	return Plugincli{}
}
