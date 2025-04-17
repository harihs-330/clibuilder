package main

import (
	"fmt"
	"log"
	"time"

	"clibuilder/toolregistry"

	"github.com/manifoldco/promptui"
)

func main() {
	fmt.Println("🤔 Hmm... what are you trying to do?")
	fmt.Println("📂 Loading available tools...")
	time.Sleep(1 * time.Second)

	tools := toolregistry.GetTools()

	if len(tools) == 0 {
		log.Fatal("⚠️  No tools registered.")
	}

	var toolOptions []string
	for _, t := range tools {
		toolOptions = append(toolOptions, fmt.Sprintf("%s - %s", t.Name, t.Description))
	}

	prompt := promptui.Select{
		Label: "Select a Tool",
		Items: toolOptions,
	}

	i, _, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed: %v", err)
	}

	selected := tools[i]
	fmt.Printf("🚀 Launching '%s'...\n\n", selected.Name)
	selected.Run()
}
