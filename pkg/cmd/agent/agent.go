package agent

import (
	"context"
	"fmt"

	"github.com/covarity/echo/pkg/pool"
	"github.com/covarity/echo/pkg/protocol/grpc"
	"github.com/covarity/echo/pkg/protocol/rest"
	"github.com/covarity/echo/pkg/runtime"
	v1 "github.com/covarity/echo/pkg/service/v1"
)

// Config is configuration for Agent
type Server struct {
	// gRPC server start parameters section
	// gRPC is TCP port to listen by gRPC server
	GRPCPort  string
	gp        *pool.GoroutinePool
	adapterGP *pool.GoroutinePool
	// HTTP/REST gateway start parameters section
	// HTTPPort is TCP port to listen by HTTP/REST gateway
	HTTPPort string
}

func RunServer() error {
	ctx := context.Background()

	// get configuration
	var server Server
	server.GRPCPort = "3000"
	server.HTTPPort = "3001"
	server.gp = pool.New(10, false)
	server.gp.AddWorkers(10)
	server.adapterGP = pool.New(10, false)
	server.adapterGP.AddWorkers(10)

	if len(server.GRPCPort) == 0 {
		return fmt.Errorf("invalid TCP port for gRPC server: '%s'", server.GRPCPort)
	}

	if len(server.HTTPPort) == 0 {
		return fmt.Errorf("invalid TCP port for HTTP gateway: '%s'", server.HTTPPort)
	}
	m := runtime.New(server.gp, server.adapterGP)

	v1API := v1.NewTaskServiceServer(m.Dispatcher())

	// run HTTP gateway
	go func() {
		_ = rest.RunServer(ctx, server.GRPCPort, server.HTTPPort)
	}()

	return grpc.RunServer(ctx, v1API, server.GRPCPort)
}
