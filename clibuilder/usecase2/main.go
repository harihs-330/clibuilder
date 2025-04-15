package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

func main() {
	// Define flags
	configFile := flag.String("config", "commands.yaml", "Path to YAML or JSON file")
	flag.Parse()

	// Get command key from args
	args := flag.Args()
	if len(args) < 1 {
		log.Fatal("❌ Please provide a command key. Example: go run main.go -config=commands.yaml list")
	}
	key := args[0]

	// Load config
	commands, err := loadCommands(*configFile)
	if err != nil {
		log.Fatalf("❌ Error loading config: %v", err)
	}

	// Get shell command
	cmdStr, exists := commands[key]
	if !exists {
		log.Fatalf("❌ Key '%s' not found in config", key)
	}

	// Execute command using shell
	out, err := exec.Command("bash", "-c", cmdStr).CombinedOutput()
	if err != nil {
		log.Fatalf("❌ Execution error: %v\nOutput:\n%s", err, out)
	}
	fmt.Printf("✅ Output:\n%s\n", out)
}

func loadCommands(path string) (map[string]string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	ext := strings.ToLower(filepath.Ext(path))
	commands := make(map[string]string)

	switch ext {
	case ".yaml", ".yml":
		err = yaml.Unmarshal(data, &commands)
	case ".json":
		err = json.Unmarshal(data, &commands)
	default:
		err = fmt.Errorf("unsupported file type: %s", ext)
	}
	return commands, err
}
