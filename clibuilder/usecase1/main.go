package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
)

func main() {
	// Parse path flag
	path := flag.String("path", "", "Path to the binary directory")
	flag.Parse()

	// Get command key from args
	args := flag.Args()
	fmt.Println("args:", args)
	if len(args) < 1 {
		log.Fatal("❗ Please provide a command. Example: go run main.go ls -path=/bin/")
	}
	commandKey := args[0]

	// Build full binary path
	var fullPath string
	fmt.Println("path:", *path)
	if *path != "" {
		fullPath = filepath.Join(*path, commandKey)
	} else {
		fullPath = commandKey
	}
	fmt.Println("fullPath:", fullPath)

	cmd := exec.Command(fullPath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("❌ Failed to execute %s: %v\nOutput:\n%s", fullPath, err, string(output))
	}

	fmt.Printf("✅ Output:\n%s\n", string(output))
}
