package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/manifoldco/promptui"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Help     HelpInfo          `yaml:"help"`
	Commands map[string]string `yaml:"commands"`
}

type HelpInfo struct {
	Description       string            `yaml:"description"`
	Examples          []string          `yaml:"examples"`
	Usage             []string          `yaml:"usage"`
	AvailableCommands map[string]string `yaml:"available_commands"`
	Flags             []string          `yaml:"flags"`
}

func main() {
	// Define flags
	configFile := "commands.yaml" // Path to YAML config file

	// Load config
	config, err := loadConfig(configFile)
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

// loadConfig reads the YAML config file and returns the parsed Config object
func loadConfig(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// ShowHelp prints the help message from the loaded config
func ShowHelp(help HelpInfo) {
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
