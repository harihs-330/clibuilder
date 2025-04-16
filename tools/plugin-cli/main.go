package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	pb "clibuilder/tools/plugin-cli/proto"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// PluginInterface defines the method our CLI expects.
type PluginInterface interface {
	Run(args []string) (string, error)
}

// PluginServer implements the gRPC Plugin service.
type PluginServer struct {
	pb.UnimplementedPluginServer
}

// PluginGRPCClient wraps the generated gRPC client.
type PluginGRPCClient struct {
	client pb.PluginClient
}

// PluginGRPC is the go-plugin wrapper for our client.
type PluginGRPC struct {
	plugin.NetRPCUnsupportedPlugin
}

// Run calls the plugin's Run method over gRPC.
func (p *PluginGRPCClient) Run(args []string) (string, error) {
	resp, err := p.client.Run(context.Background(), &pb.RunRequest{Args: args})
	if err != nil {
		return "", err
	}
	return resp.Message, nil
}

// GRPCServer registers our gRPC server.
func (p *PluginGRPC) GRPCServer(broker *plugin.GRPCBroker, server *grpc.Server) error {
	pb.RegisterPluginServer(server, &PluginServer{})
	return nil
}

// GRPCClient creates a new gRPC client.
func (p *PluginGRPC) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, client *grpc.ClientConn) (interface{}, error) {
	return &PluginGRPCClient{
		client: pb.NewPluginClient(client),
	}, nil
}

// loadPlugins scans the given directory for plugin binaries.
func loadPlugins(directory string) (map[string]*plugin.Client, error) {
	plugins := make(map[string]*plugin.Client)
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		pluginName := info.Name()
		pluginPath := filepath.Join(directory, pluginName)
		client := plugin.NewClient(&plugin.ClientConfig{
			HandshakeConfig: plugin.HandshakeConfig{
				ProtocolVersion:  1,
				MagicCookieKey:   "PLUGIN_MAGIC_COOKIE",
				MagicCookieValue: "cli_builder",
			},
			Plugins: map[string]plugin.Plugin{
				"cli_plugin": &PluginGRPC{},
			},
			Cmd:              exec.Command("sh", "-c", pluginPath),
			SyncStdout:       os.Stdout,
			SyncStderr:       os.Stderr,
			Logger:           hclog.New(&hclog.LoggerOptions{Name: "plugin-client", Level: hclog.Error}),
			AllowedProtocols: []plugin.Protocol{plugin.ProtocolGRPC},
		})
		plugins[pluginName] = client
		return nil
	})
	return plugins, err
}

// executePlugin runs the plugin with the given arguments.
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
	pluginInstance, ok := raw.(PluginInterface)
	if !ok {
		log.Fatalf("Error: plugin does not implement PluginInterface")
	}
	result, err := pluginInstance.Run(args)
	if err != nil {
		log.Fatalf("Error running plugin: %v", err)
	}
	fmt.Println(result)
}

func main() {
	// Load plugins from the plugins directory.
	plugins, err := loadPlugins("./plugins")
	if err != nil {
		log.Fatalf("Error loading plugins: %v", err)
	}

	// Get plugin names for the interactive prompt.
	var pluginNames []string
	for name := range plugins {
		pluginNames = append(pluginNames, name)
	}

	// Interactive prompt to select a plugin
	prompt := promptui.Select{
		Label: "Select a plugin to run",
		Items: pluginNames,
	}

	_, selectedPlugin, err := prompt.Run()
	if err != nil {
		log.Fatalf("Error selecting plugin: %v", err)
	}

	// Run the selected plugin interactively.
	pluginClient := plugins[selectedPlugin]
	cmd := &cobra.Command{
		Use:   selectedPlugin,
		Short: fmt.Sprintf("Run plugin %s", selectedPlugin),
		Run: func(cmd *cobra.Command, args []string) {
			executePlugin(pluginClient, args)
		},
	}

	// Optionally, allow user to pass arguments
	var args []string
	for {
		promptArgs := promptui.Prompt{
			Label: "Enter arguments for the plugin (comma separated), or press Enter to run",
		}
		result, err := promptArgs.Run()
		if err != nil {
			log.Fatalf("Error reading arguments: %v", err)
		}
		if result == "" {
			break
		}
		args = append(args, result)
	}

	// Execute the plugin
	cmd.Run(cmd, args)
}
