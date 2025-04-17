package main

import (
	"fmt"
	"log"
	"time"

	"clibuilder/toolregistry"

	"github.com/manifoldco/promptui"
)

func main() {
	fmt.Println("🤔 Welcome to hmm — your CLI toolbox")
	fmt.Println("🔍 Scanning for tools...")
	time.Sleep(1 * time.Second)

	tools := toolregistry.GetTools()

	if len(tools) == 0 {
		log.Fatal("⚠️  No tools registered.")
	}

	var toolNames []string
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
