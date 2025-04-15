package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install [repo_url]",
	Short: "Clone a public repository",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		repoURL := args[0]
		repoName := getRepoName(repoURL)

		if _, err := os.Stat("./repos/" + repoName); err == nil {
			fmt.Println("Repository already installed.")
			return
		}

		err := os.MkdirAll("./repos", 0755)
		if err != nil {
			fmt.Println("Error creating repos directory:", err)
			return
		}

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

func getRepoName(repoURL string) string {
	base := path.Base(repoURL)
	return strings.TrimSuffix(base, ".git")
}
