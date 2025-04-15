package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

var path string

var runCmd = &cobra.Command{
	Use:   "run <command>",
	Short: "Run a binary command from the specified or default path",
	Args:  cobra.MinimumNArgs(1),
	Example: `
  binrunner run ls --path /bin
  binrunner run date
`,
	Run: func(cmd *cobra.Command, args []string) {
		command := args[0]
		fullPath := command

		if path != "" {
			fullPath = filepath.Join(path, command)
		}

		execCmd := exec.Command(fullPath)
		output, err := execCmd.CombinedOutput()
		if err != nil {
			log.Fatalf("❌ Error running %q: %v\nOutput:\n%s", fullPath, err, output)
		}

		fmt.Printf("✅ Output:\n%s\n", output)
	},
}
