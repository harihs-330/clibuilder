package cmd

import (
	"fmt"
	"os"
	"os/exec"
)

func UpgradeAction() {
	repoURL := promptRepo("Install")
	fmt.Printf("Installing repository: %s\n", repoURL)
	repoName := getRepoName(repoURL)

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

}
