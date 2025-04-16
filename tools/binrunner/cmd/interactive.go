package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var interactiveCmd = &cobra.Command{
	Use:   "interactive",
	Short: "Run binrunner in interactive mode using config-defined options",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := LoadConfig("../config.json")
		if err != nil {
			log.Fatalf("❌ Failed to load config: %v", err)
		}

		// Command selection
		cmdSelect := promptui.Select{
			Label: "Select command to run",
			Items: config.Commands,
		}
		_, selectedCommand, err := cmdSelect.Run()
		if err != nil {
			log.Fatalf("❌ Command prompt failed: %v", err)
		}
		if selectedCommand == "Custom" {
			cmdPrompt := promptui.Prompt{Label: "Enter custom command"}
			selectedCommand, err = cmdPrompt.Run()
			if err != nil {
				log.Fatalf("❌ Command input failed: %v", err)
			}
		}

		// Path selection
		pathSelect := promptui.Select{
			Label: "Select path to binary",
			Items: config.Paths,
		}
		_, selectedPath, err := pathSelect.Run()
		if err != nil {
			log.Fatalf("❌ Path prompt failed: %v", err)
		}
		if selectedPath == "Custom" {
			pathPrompt := promptui.Prompt{Label: "Enter custom path"}
			selectedPath, err = pathPrompt.Run()
			if err != nil {
				log.Fatalf("❌ Path input failed: %v", err)
			}
		}

		fullPath := filepath.Join(strings.TrimSpace(selectedPath), strings.TrimSpace(selectedCommand))

		execCmd := exec.Command(fullPath)
		output, err := execCmd.CombinedOutput()
		if err != nil {
			log.Fatalf("❌ Error running %q: %v\nOutput:\n%s", fullPath, err, output)
		}

		fmt.Printf("✅ Output of %s:\n%s\n", fullPath, output)
	},
}
