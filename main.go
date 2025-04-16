package main

import (
	"fmt"
	"log"
	"time"

	"github.com/manifoldco/promptui"

	"clibuilder/toolregistry"
)

func main() {
	fmt.Println("👋 Welcome to the CLI Launcher!")
	fmt.Println("📂 Loading available tools...")
	time.Sleep(1 * time.Second)

	tools := toolregistry.GetTools()

	if len(tools) == 0 {
		log.Fatal("No tools registered")
	}

	toolNames := []string{}
	for _, t := range tools {
		toolNames = append(toolNames, t.Name())
	}

	prompt := promptui.Select{
		Label: "Select a Tool",
		Items: toolNames,
	}

	i, _, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed: %v", err)
	}

	selectedTool := tools[i]
	fmt.Printf("🚀 Launching '%s'...\n\n", selectedTool.Name())
	selectedTool.Run()
}
