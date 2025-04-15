package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/manifoldco/promptui"
)

func main() {
	// Folder to scan for binaries (can change to "./bin" or any path)
	binFolder := "."

	// List executable files
	binaries, err := listBinaries(binFolder)
	if err != nil {
		log.Fatalf("Error listing binaries: %v", err)
	}

	if len(binaries) == 0 {
		log.Fatal("No binaries found in the folder")
	}

	// Prompt user to select binary
	prompt := promptui.Select{
		Label: "Select Binary to Execute",
		Items: binaries,
	}

	_, selected, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed: %v", err)
	}

	// Run the binary with bash -c './binary'
	err = execBinary(selected)
	if err != nil {
		log.Fatalf("Error executing binary: %v", err)
	}
}

func listBinaries(folder string) ([]string, error) {
	files, err := os.ReadDir(folder)
	if err != nil {
		return nil, err
	}

	var binaries []string
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		fullPath := filepath.Join(folder, file.Name())

		info, err := os.Stat(fullPath)
		if err != nil {
			continue
		}

		// Check if it's executable
		mode := info.Mode()
		if mode&0111 != 0 {
			binaries = append(binaries, file.Name())
		}
	}

	return binaries, nil
}

func execBinary(name string) error {
	cmd := exec.Command("bash", "-c", fmt.Sprintf("./%s", strings.TrimSpace(name)))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin // allow interactive input if needed

	return cmd.Run()
}
