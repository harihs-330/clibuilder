package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
)

func InstallAction() {
	repoURL := promptRepo("Install")
	fmt.Printf("Installing repository: %s\n", repoURL)
	repoName := getRepoName(repoURL)

	if _, err := os.Stat("./repos/" + repoName); err == nil {
		fmt.Println("Repository already installed.")
		return
	}

	if err := os.MkdirAll("./repos", 0755); err != nil {
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
}

func getRepoName(repoURL string) string {
	base := path.Base(repoURL)
	return strings.TrimSuffix(base, ".git")
}
