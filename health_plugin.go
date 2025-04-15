package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"

	pb "clibuilder/clibuilder/usecase4/proto"
)

// PluginServer implements health check logic
type PluginServer struct {
	pb.UnimplementedPluginServer
}

func (p *PluginServer) Run(ctx context.Context, req *pb.RunRequest) (*pb.RunResponse, error) {
	if len(req.Args) == 0 {
		return &pb.RunResponse{Message: "No URL provided"}, nil
	}

	url := req.Args[0]
	resp, err := http.Get(url)
	if err != nil {
		return &pb.RunResponse{Message: fmt.Sprintf("❌  Failed to reach %s: %v", url, err)}, nil
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return &pb.RunResponse{Message: fmt.Sprintf("✅ %s is healthy", url)}, nil
	}

	return &pb.RunResponse{Message: fmt.Sprintf("❌ %s returned status %d", url, resp.StatusCode)}, nil
}

// Plugin struct for go-plugin
type PluginGRPC struct {
	plugin.NetRPCUnsupportedPlugin
}

func (p *PluginGRPC) GRPCServer(broker *plugin.GRPCBroker, server *grpc.Server) error {
	pb.RegisterPluginServer(server, &PluginServer{})
	return nil
}

func (p *PluginGRPC) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, cc *grpc.ClientConn) (interface{}, error) {
	return nil, fmt.Errorf("this is a plugin binary only")
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: plugin.HandshakeConfig{
			ProtocolVersion:  1,
			MagicCookieKey:   "PLUGIN_MAGIC_COOKIE",
			MagicCookieValue: "cli_builder",
		},
		Plugins: map[string]plugin.Plugin{
			"cli_plugin": &PluginGRPC{},
		},
		GRPCServer: plugin.DefaultGRPCServer,
		Logger:     hclog.New(&hclog.LoggerOptions{Name: "url-health-plugin", Level: hclog.Debug}),
	})
}
