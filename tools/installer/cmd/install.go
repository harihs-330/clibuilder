package cmd

// import (
// 	"fmt"
// 	"log"
// 	"os/exec"
// 	"sort"

// 	"github.com/manifoldco/promptui"
// 	"github.com/spf13/cobra"
// )

// var installCmd = &cobra.Command{
// 	Use:   "install",
// 	Short: "Install a tool from tools.yaml",
// 	Run: func(cmd *cobra.Command, args []string) {
// 		tools, err := loadToolsConfig("tools.yaml")
// 		if err != nil {
// 			log.Fatalf("Failed to load tools.yaml: %v", err)
// 		}

// 		var toolNames []string
// 		for name := range tools {
// 			toolNames = append(toolNames, name)
// 		}
// 		sort.Strings(toolNames)

// 		prompt := promptui.Select{
// 			Label: "Select a tool to install",
// 			Items: toolNames,
// 		}

// 		index, _, err := prompt.Run()
// 		if err != nil {
// 			log.Fatalf("Prompt failed: %v", err)
// 		}

// 		tool := toolNames[index]
// 		cmdStr := tools[tool]

// 		fmt.Printf("📦 Installing %s...\n", tool)
// 		execCmd := exec.Command("bash", "-c", cmdStr)
// 		execCmd.Stdout = log.Writer()
// 		execCmd.Stderr = log.Writer()

// 		if err := execCmd.Run(); err != nil {
// 			log.Fatalf("Failed to install %s: %v", tool, err)
// 		}

// 		fmt.Println("✅ Installation complete.")
// 	},
// }
