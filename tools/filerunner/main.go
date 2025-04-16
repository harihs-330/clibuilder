package main

import (
	"clibuilder/tools/filerunner/config"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/manifoldco/promptui"
)

func main() {
	// Define flags
	configFile := "/Users/hariharasudhan/Documents/clibuilder/mycli/tools/filerunner/commands.yaml" // Path to YAML config file

	// Load config
	config, err := config.LoadConfig(configFile)
	fmt.Println("ďdddd", config)
	if err != nil {
		log.Fatalf("❌ Error loading config: %v", err)
	}

	// Check if no arguments are provided (interactive mode)
	if len(os.Args) < 2 {
		// Interactive prompt
		InteractiveMode(config.Commands)
		return
	}

	// Non-interactive: Get command key from args
	key := os.Args[1]

	// Execute command from config
	ExecuteCommand(key, config.Commands)
}

// ShowHelp prints the help message from the loaded config
func ShowHelp(help config.HelpInfo) {
	fmt.Println(help.Description)
	fmt.Println("\nExamples:")
	for _, example := range help.Examples {
		fmt.Println("  " + example)
	}
	fmt.Println("\nUsage:")
	for _, usage := range help.Usage {
		fmt.Println("  " + usage)
	}
	fmt.Println("\nAvailable Commands:")
	for cmd, desc := range help.AvailableCommands {
		fmt.Printf("  %s: %s\n", cmd, desc)
	}
	fmt.Println("\nFlags:")
	for _, flag := range help.Flags {
		fmt.Println("  " + flag)
	}
}

// ExecuteCommand will execute a command based on the user's input
func ExecuteCommand(cmd string, commands map[string]string) {
	cmdStr, exists := commands[cmd]
	if !exists {
		log.Fatalf("❌ Command '%s' not found in config", cmd)
	}

	// Execute command using shell
	out, err := exec.Command("bash", "-c", cmdStr).CombinedOutput()
	if err != nil {
		log.Fatalf("❌ Execution error: %v\nOutput:\n%s", err, out)
	}
	fmt.Printf("✅ Output:\n%s\n", out)
}

// InteractiveMode provides an interactive CLI prompt for the user
func InteractiveMode(commands map[string]string) {
	// Create a list of command names from the config
	var commandList []string
	for cmd := range commands {
		commandList = append(commandList, cmd)
	}

	// Create a prompt with the available commands
	prompt := promptui.Select{
		Label: "Select Command",
		Items: commandList,
	}

	_, result, err := prompt.Run()
	if err != nil {
		log.Fatalf("❌ Prompt error: %v", err)
	}

	// Execute the selected command
	ExecuteCommand(result, commands)
}
