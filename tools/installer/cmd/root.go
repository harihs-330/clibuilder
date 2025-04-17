package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "installer",
	Short: "Tool Installer CLI",
	Long:  "Install tools from a YAML config using interactive prompt.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(installCmd)
}
