package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Clone a public repository",
	Run: func(cmd *cobra.Command, args []string) {
		// Prompt the user for the repository URL
		repoURL := getRepoURL()

		// Extract the repository name from the URL
		repoName := getRepoName(repoURL)

		// Check if the repository already exists
		if _, err := os.Stat("./repos/" + repoName); err == nil {
			fmt.Println("Repository already installed.")
			return
		}

		// Create the repos directory if it doesn't exist
		err := os.MkdirAll("./repos", 0755)
		if err != nil {
			fmt.Println("Error creating repos directory:", err)
			return
		}

		// Clone the repository using git
		cmdGit := exec.Command("git", "clone", repoURL, "./repos/"+repoName)
		cmdGit.Stdout = os.Stdout
		cmdGit.Stderr = os.Stderr
		if err := cmdGit.Run(); err != nil {
			fmt.Println("Failed to clone repo:", err)
		} else {
			fmt.Println("Repository installed:", repoName)
		}
	},
}

// getRepoURL prompts the user to enter the repository URL
func getRepoURL() string {
	prompt := promptui.Prompt{
		Label: "Enter the repository URL",
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Println("Failed to get input:", err)
		os.Exit(1)
	}

	return result
}

// getRepoName extracts the repository name from the URL
func getRepoName(repoURL string) string {
	// Extract the base name from the URL
	base := path.Base(repoURL)
	// Remove the ".git" extension
	return strings.TrimSuffix(base, ".git")
}
