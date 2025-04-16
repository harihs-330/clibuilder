package cmd

import (
	"fmt"
	"os"

	"os/exec"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var configFile = "tools/filerunner/commands.yaml"
var commands map[string]string

var rootCmd = &cobra.Command{
	Use:   "filerunner",
	Short: "Run commands from config",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := LoadConfig(configFile)
		if err != nil {
			return fmt.Errorf("failed to load config: %v", err)
		}
		commands = cfg.Commands
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		interactiveMode(commands)
	},
}

var runCmd = &cobra.Command{
	Use:   "run [key]",
	Short: "Run command by key from config",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		executeCommand(key, commands)
	},
}

func Execute() error {
	rootCmd.AddCommand(runCmd)
	return rootCmd.Execute()
}

func executeCommand(key string, commands map[string]string) {
	cmdStr, exists := commands[key]
	if !exists {
		fmt.Printf("❌ Command '%s' not found\n", key)
		os.Exit(1)
	}

	out, err := exec.Command("bash", "-c", cmdStr).CombinedOutput()
	if err != nil {
		fmt.Printf("❌ Error executing command: %v\nOutput:\n%s", err, out)
		os.Exit(1)
	}

	fmt.Printf("✅ Output:\n%s\n", out)
}

func interactiveMode(commands map[string]string) {
	var keys []string
	for k := range commands {
		keys = append(keys, k)
	}

	prompt := promptui.Select{
		Label: "Select Command",
		Items: keys,
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Println("Prompt failed:", err)
		os.Exit(1)
	}

	executeCommand(result, commands)
}
