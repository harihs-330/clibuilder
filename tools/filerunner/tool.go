package filerunner

import (
	"clibuilder/tools/filerunner/cmd"
	"fmt"
)

type Filerunner struct{}

func (b Filerunner) Name() string {
	return "filerunner"
}

func (b Filerunner) Run() {
	fmt.Println("🧪 Running Filerunner tool!")
	cmd.Execute()

}

func New() Filerunner {
	return Filerunner{}
}
