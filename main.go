package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/manifoldco/promptui"
)

func main() {
	fmt.Println("👋 Welcome to the CLI Launcher!")
	fmt.Println("📂 Scanning available tools...")
	time.Sleep(2 * time.Second)

	// Set tools folder path
	toolsFolder := "./tools"

	// Get list of tool folder names
	toolDirs, err := listToolDirectories(toolsFolder)
	if err != nil {
		log.Fatalf("Error listing tools: %v", err)
	}

	if len(toolDirs) == 0 {
		log.Fatal("No tools found in the 'tools' directory")
	}

	// Prompt user to select tool
	prompt := promptui.Select{
		Label: "Select a Tool",
		Items: toolDirs,
	}

	_, selectedTool, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed: %v", err)
	}

	// Build path to selected tool's main.go
	toolPath := filepath.Join(toolsFolder, selectedTool, "main.go")

	// Execute: go run tools/<tool>/main.go
	cmd := exec.Command("go", "run", toolPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	fmt.Printf("🚀 Launching '%s' tool...\n", selectedTool)
	if err := cmd.Run(); err != nil {
		log.Fatalf("Error running tool: %v", err)
	}
}

// listToolDirectories lists all subdirectories inside tools/
func listToolDirectories(root string) ([]string, error) {
	dirs := []string{}
	entries, err := os.ReadDir(root)
	if err != nil {
		return dirs, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			dirs = append(dirs, entry.Name())
		}
	}

	return dirs, nil
}
