package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "List all installed repositories",
	Run: func(cmd *cobra.Command, args []string) {
		repoDir := "./repos"
		files, err := ioutil.ReadDir(repoDir)
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Println("No repositories found.")
			} else {
				fmt.Println("Error reading repo directory:", err)
			}
			return
		}

		fmt.Println("Installed repositories:")
		for _, f := range files {
			if f.IsDir() {
				fmt.Println(" -", f.Name())
			}
		}
	},
}
