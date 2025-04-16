package cmd

import (
	"clibuilder/utils"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "binrunner",
	Short: "A CLI tool to run system binaries from custom paths",
	Long: `binrunner is a lightweight CLI tool that allows you to execute 
binary commands from your system or a specific directory path.

Examples:
  binrunner run ls --path /bin
  binrunner run date
`,
	// Prompt UI logic added here
	Run: func(cmd *cobra.Command, args []string) {
		mode, customPath, err := PromptUser()
		if err != nil {
			fmt.Fprintf(os.Stderr, "❌ Prompt failed: %v\n", err)
			os.Exit(1)
		}

		switch mode {
		case "interactive":
			interactiveCmd.Run(cmd, args)
		case "custom":
			// Prompt for binary name
			binaryPrompt := promptui.Prompt{
				Label: "Enter the binary to run (e.g., ls, pwd, date)",
			}
			binaryName, err := binaryPrompt.Run()
			if err != nil {
				fmt.Fprintf(os.Stderr, "❌ Binary input error: %v\n", err)
				os.Exit(1)
			}

			// Optional: Prompt for arguments (space-separated)
			argsPrompt := promptui.Prompt{
				Label:   "Enter arguments (optional)",
				Default: "",
			}
			argString, _ := argsPrompt.Run()
			userArgs := []string{}
			if argString != "" {
				userArgs = append(userArgs, binaryName)
				userArgs = append(userArgs, utils.SplitArgs(argString)...)
			} else {
				userArgs = append(userArgs, binaryName)
			}

			path = customPath
			runCmd.Run(cmd, userArgs)
		default:
			_ = cmd.Help()
		}
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

func PromptUser() (string, string, error) {
	prompt := promptui.Select{
		Label: "Select Mode",
		Items: []string{"interactive", "custom"},
	}

	_, result, err := prompt.Run()
	if err != nil {
		return "", "", err
	}

	if result == "custom" {
		pathPrompt := promptui.Prompt{
			Label: "Enter Custom Path to Binary Directory (e.g., /usr/local/bin)",
		}
		customPath, err := pathPrompt.Run()
		if err != nil {
			return "", "", err
		}
		return "custom", customPath, nil
	}

	return "interactive", "", nil
}
