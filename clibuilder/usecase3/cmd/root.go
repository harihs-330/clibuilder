package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "repo-cli",
	Short: "CLI tool to manage public repos",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.AddCommand(installCmd)
	rootCmd.AddCommand(upgradeCmd)
	rootCmd.AddCommand(viewCmd)
}
