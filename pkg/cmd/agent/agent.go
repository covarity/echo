package agent

import (
	"context"
	"fmt"

	"github.com/covarity/echo/pkg/protocol/grpc"
	"github.com/covarity/echo/pkg/protocol/rest"
	v1 "github.com/covarity/echo/pkg/service/v1"
	"github.com/covarity/echo/pkg/queue"
)

// Config is configuration for Server
type Config struct {
	// gRPC server start parameters section
	// gRPC is TCP port to listen by gRPC server
	GRPCPort string

	// HTTP/REST gateway start parameters section
	// HTTPPort is TCP port to listen by HTTP/REST gateway
	HTTPPort string
}

func RunServer() error {
	ctx := context.Background()

	// get configuration
	var cfg Config
	cfg.GRPCPort = "3000"
	cfg.HTTPPort = "3001"

	if len(cfg.GRPCPort) == 0 {
		return fmt.Errorf("invalid TCP port for gRPC server: '%s'", cfg.GRPCPort)
	}

	if len(cfg.HTTPPort) == 0 {
		return fmt.Errorf("invalid TCP port for HTTP gateway: '%s'", cfg.HTTPPort)
	}
	q := queue.New()

	v1API := v1.NewTaskServiceServer(q)

	// run HTTP gateway
	go func() {
		_ = rest.RunServer(ctx, cfg.GRPCPort, cfg.HTTPPort)
	}()

	return grpc.RunServer(ctx, v1API, cfg.GRPCPort)
}
