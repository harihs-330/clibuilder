package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var upgradeCmd = &cobra.Command{
	Use:   "upgrade [repo_name]",
	Short: "Pull latest changes for a repo",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		repoName := args[0]
		repoPath := "./repos/" + repoName

		if _, err := os.Stat(repoPath); os.IsNotExist(err) {
			fmt.Println("Repository not found:", repoName)
			return
		}

		cmdGit := exec.Command("git", "-C", repoPath, "pull")
		cmdGit.Stdout = os.Stdout
		cmdGit.Stderr = os.Stderr
		if err := cmdGit.Run(); err != nil {
			fmt.Println("Failed to upgrade repo:", err)
		} else {
			fmt.Println("Repository upgraded:", repoName)
		}
	},
}
