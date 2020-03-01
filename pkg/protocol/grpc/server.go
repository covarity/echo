package grpc

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"

	v1 "github.com/covarity/echo/pkg/api/v1"
)

// RunServer runs gRPC service to publish ToDo service
func RunServer(ctx context.Context, v1API v1.TaskServiceServer, v1Health v1.HealthServer, port string) error {
	listen, err := net.Listen("tcp", ":" + port)
	if err != nil {
		return err
	}

	// register service
	server := grpc.NewServer()
	v1.RegisterTaskServiceServer(server, v1API)
	v1.RegisterHealthServer(server,v1Health)
	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			log.Println("shutting down gRPC server...")

			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	// start gRPC server
	log.Println("starting gRPC server...")
	return server.Serve(listen)
}
