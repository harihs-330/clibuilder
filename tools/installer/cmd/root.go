package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"sort"

	"io/ioutil"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

type ToolMap map[string]string

const configPath = "tools/installer/cmd/tools.yaml"

var rootCmd = &cobra.Command{
	Use:   "installer",
	Short: "Install tools from a YAML config",
	Run: func(cmd *cobra.Command, args []string) {
		tools, err := loadToolsConfig(configPath)
		if err != nil {
			log.Fatalf("❌ Failed to load tools.yaml: %v", err)
		}

		// Prompt for mode
		modePrompt := promptui.Select{
			Label: "Select Mode",
			Items: []string{"Interactive Mode", "Quiet Mode"},
		}

		modeIndex, _, err := modePrompt.Run()
		if err != nil {
			log.Fatalf("Prompt failed: %v", err)
		}

		switch modeIndex {
		case 0:
			runInteractiveMode(tools)
		case 1:
			showQuietModeHelp(tools)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func runInteractiveMode(tools ToolMap) {
	var toolNames []string
	for name := range tools {
		toolNames = append(toolNames, name)
	}
	sort.Strings(toolNames)

	toolPrompt := promptui.Select{
		Label: "Select a tool to install",
		Items: toolNames,
	}

	index, _, err := toolPrompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed: %v", err)
	}

	selectedTool := toolNames[index]
	cmdStr := tools[selectedTool]

	fmt.Printf("📦 Installing %s...\n", selectedTool)

	cmd := exec.Command("bash", "-c", cmdStr)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatalf("Install failed: %v", err)
	}

	fmt.Println("✅ Installation complete.")
}

func showQuietModeHelp(tools ToolMap) {
	fmt.Println("\n📖 Available Tools:")
	fmt.Println("---------------------")
	for name, cmd := range tools {
		fmt.Printf("🔹 %s -> %s\n", name, cmd)
	}

	fmt.Println("\n📌 Example Usage:")
	fmt.Println("  clibuilder installer --tool git --quiet")
	fmt.Println("\n📝 future: support tags, batch install, YAML metadata.")
}

func loadToolsConfig(filename string) (ToolMap, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var tools ToolMap
	err = yaml.Unmarshal(data, &tools)
	return tools, err
}
