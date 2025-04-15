package cmd

import (
	"fmt"
	"log"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "repo-cli",
	Short: "CLI tool to manage public repos",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func promptRepo(action string) string {
	prompt := promptui.Prompt{
		Label: fmt.Sprintf("%s Repository Name", action),
		Validate: func(input string) error {
			if len(input) == 0 {
				return fmt.Errorf("repository name cannot be empty")
			}
			return nil
		},
	}

	repo, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed: %v\n", err)
	}
	return repo
}
