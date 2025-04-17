package toolregistry

import (
	"clibuilder/tools/binrunner"
	"clibuilder/tools/filerunner"
	"clibuilder/tools/installer"
	"clibuilder/tools/plugincli"
	"clibuilder/tools/repocli"
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Tool struct {
	Name        string
	Description string
	Run         func()
}

type toolConfig struct {
	Description string `yaml:"description"`
}

var runFuncs = map[string]func(){
	"binrunner": func() {
		fmt.Println("🔧 Running binrunner...")
		binrunner.New().Run()
	},
	"filerunner": func() {
		fmt.Println("📁 Running filerunner...")
		filerunner.New().Run()
	},
	"plugincli": func() {
		fmt.Println("🔌 Running plugincli...")
		plugincli.New().Run()
	},
	"repocli": func() {
		fmt.Println("🔌 Running repocli...")
		repocli.New().Run()
	},
	"installer": func() {
		fmt.Println("📦 Running installer...")
		installer.New().Run()
	},
}

// GetTools loads the YAML and returns a list of Tool structs
func GetTools() []Tool {
	data, err := ioutil.ReadFile("toolregistry/tools.yaml")
	if err != nil {
		log.Fatalf("Failed to read tools.yaml: %v", err)
	}

	var raw map[string]toolConfig
	if err := yaml.Unmarshal(data, &raw); err != nil {
		log.Fatalf("Failed to parse tools.yaml: %v", err)
	}

	var tools []Tool
	for name, conf := range raw {
		runFunc, ok := runFuncs[name]
		if !ok {
			runFunc = func() { fmt.Printf("🚫 No run function defined for '%s'\n", name) }
		}

		tools = append(tools, Tool{
			Name:        name,
			Description: conf.Description,
			Run:         runFunc,
		})
	}

	return tools
}
