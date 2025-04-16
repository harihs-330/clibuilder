package main

import (
	"clibuilder/tools/repo-cli/cmd"
	"fmt"
	"log"

	"github.com/manifoldco/promptui"
)

func main() {
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
		cmd.InstallAction()
	case "upgrade":
		cmd.UpgradeAction()
	case "view":
		cmd.ViewAction()
	default:
		fmt.Println("Unknown action")
	}
}
