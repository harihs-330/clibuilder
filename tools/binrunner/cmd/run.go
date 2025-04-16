package cmd

import (
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

var path string

var runCmd = &cobra.Command{
	Use:   "run [binary]",
	Short: "Run a binary from the specified path",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		binary := args[0]
		fullPath := filepath.Join(path, binary)

		out, err := exec.Command(fullPath, args[1:]...).CombinedOutput()
		if err != nil {
			fmt.Printf("❌ Failed to run command: %v\n", err)
			fmt.Printf("Output:\n%s\n", out)
			return
		}
		fmt.Printf("✅ Output:\n%s\n", out)
	},
}
