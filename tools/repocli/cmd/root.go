package cmd

import (
	"fmt"
	"log"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "repo-cli",
	Short: "CLI to manage repositories",
	Run: func(cmd *cobra.Command, args []string) {
		actions := []string{"install", "upgrade", "view"}

		prompt := promptui.Select{
			Label: "Select Action",
			Items: actions,
		}

		_, action, err := prompt.Run()
		if err != nil {
			log.Fatalf("Prompt failed %v\n", err)
		}

		switch action {
		case "install":
			InstallAction()
		case "upgrade":
			UpgradeAction()
		case "view":
			ViewAction()
		default:
			fmt.Println("Unknown action")
		}
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func promptRepo(action string) string {
	prompt := promptui.Prompt{
		Label: fmt.Sprintf("%s Repository URL", action),
		Validate: func(input string) error {
			if input == "" {
				return fmt.Errorf("repository URL cannot be empty")
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
