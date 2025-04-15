package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base/root command
var rootCmd = &cobra.Command{
	Use:   "binrunner",
	Short: "A CLI tool to run system binaries from custom paths",
	Long: `binrunner is a lightweight CLI tool that allows you to execute 
binary commands from your system or a specific directory path.

Examples:
  binrunner run ls --path /bin
  binrunner run date
`,
	// If no subcommand is given, show help
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

// Execute is called by main.main()
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "❌ Error: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	runCmd.Flags().StringVar(&path, "path", "", "Path to the binary directory")
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(interactiveCmd)

}
