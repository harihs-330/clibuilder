package installer

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"sort"

	"github.com/manifoldco/promptui"
	"gopkg.in/yaml.v2"
)

type Installer struct{}

func (i Installer) Name() string {
	return "installer"
}

func (i Installer) Run() {
	fmt.Println("🛠️ Tool Installer: loading available tools...")

	tools, err := loadToolsConfig("tools/installer/tools.yaml")
	if err != nil {
		log.Fatalf("❌ Failed to load tools.yaml: %v", err)
	}

	var toolNames []string
	for name := range tools {
		toolNames = append(toolNames, name)
	}
	sort.Strings(toolNames)

	prompt := promptui.Select{
		Label: "Select a tool to install",
		Items: toolNames,
	}

	index, _, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed: %v", err)
	}

	selectedTool := toolNames[index]
	cmdStr := tools[selectedTool]

	fmt.Printf("📦 Installing %s...\n", selectedTool)

	cmd := exec.Command("bash", "-c", cmdStr)
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	if err := cmd.Run(); err != nil {
		log.Fatalf("Install failed: %v", err)
	}

	fmt.Println("✅ Installation complete.")
}

// ✅ This function must be exported
func New() Installer {
	return Installer{}
}

type ToolMap map[string]string

func loadToolsConfig(filename string) (ToolMap, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var tools ToolMap
	err = yaml.Unmarshal(data, &tools)
	return tools, err
}
