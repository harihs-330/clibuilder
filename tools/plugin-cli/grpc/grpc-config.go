package grpc

import (
	pb "clibuilder/tools/plugin-cli/proto"
	"context"

	"github.com/hashicorp/go-plugin"
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
