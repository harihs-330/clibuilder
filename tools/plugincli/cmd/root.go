package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"clibuilder/tools/plugincli/grpc"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "plugin-cli",
	Short: "Run available plugins",
	Run: func(cmd *cobra.Command, args []string) {
		runInteractive()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func runInteractive() {
	plugins, err := loadPlugins("./plugins")
	if err != nil {
		log.Fatalf("Error loading plugins: %v", err)
	}

	if len(plugins) == 0 {
		log.Fatal("No plugins found")
	}

	var pluginNames []string
	for name := range plugins {
		pluginNames = append(pluginNames, name)
	}

	prompt := promptui.Select{
		Label: "Select a plugin to run",
		Items: pluginNames,
	}

	_, selectedPlugin, err := prompt.Run()
	if err != nil {
		log.Fatalf("Error selecting plugin: %v", err)
	}

	pluginClient := plugins[selectedPlugin]

	// Get args from user
	promptArgs := promptui.Prompt{
		Label: "Enter args (comma separated)",
	}
	result, _ := promptArgs.Run()
	args := strings.Split(result, ",")

	executePlugin(pluginClient, args)
}

func loadPlugins(directory string) (map[string]*plugin.Client, error) {
	plugins := make(map[string]*plugin.Client)
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return err
		}
		client := plugin.NewClient(&plugin.ClientConfig{
			HandshakeConfig: plugin.HandshakeConfig{
				ProtocolVersion:  1,
				MagicCookieKey:   "PLUGIN_MAGIC_COOKIE",
				MagicCookieValue: "cli_builder",
			},
			Plugins: map[string]plugin.Plugin{
				"cli_plugin": &grpc.PluginGRPC{},
			},
			Cmd:              exec.Command("sh", "-c", path),
			SyncStdout:       os.Stdout,
			SyncStderr:       os.Stderr,
			Logger:           hclog.New(&hclog.LoggerOptions{Name: "plugin-client", Level: hclog.Error}),
			AllowedProtocols: []plugin.Protocol{plugin.ProtocolGRPC},
		})
		plugins[info.Name()] = client
		return nil
	})
	return plugins, err
}

func executePlugin(client *plugin.Client, args []string) {
	rpcClient, err := client.Client()
	if err != nil {
		log.Fatalf("Error starting plugin: %v", err)
	}
	defer client.Kill()

	raw, err := rpcClient.Dispense("cli_plugin")
	if err != nil {
		log.Fatalf("Error dispensing plugin: %v", err)
	}

	pluginInstance, ok := raw.(grpc.PluginInterface)
	if !ok {
		log.Fatalf("Error: plugin does not implement PluginInterface")
	}

	result, err := pluginInstance.Run(args)
	if err != nil {
		log.Fatalf("Plugin execution failed: %v", err)
	}

	fmt.Println("✅ Output:")
	fmt.Println(result)
}
