package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/manifoldco/promptui"
)

func ViewAction() {

	repoDir := "./repos"
	files, err := ioutil.ReadDir(repoDir)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("No repositories found.")
		} else {
			fmt.Println("Error reading repo directory:", err)
		}
		return
	}

	// Prompt user to select a repo from the list of installed repositories
	repos := []string{}
	for _, f := range files {
		if f.IsDir() {
			repos = append(repos, f.Name())
		}
	}

	if len(repos) == 0 {
		fmt.Println("No repositories available.")
		return
	}

	// Prompt user to select a repository
	prompt := promptui.Select{
		Label: "Select a repository to view",
		Items: repos,
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Println("Failed to select repository:", err)
		return
	}

	fmt.Println("Selected repository:", result)

}
