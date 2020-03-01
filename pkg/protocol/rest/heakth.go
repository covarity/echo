package rest

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	v1 "github.com/covarity/echo/pkg/api/v1"
)

// RunServer runs HTTP/REST gateway
func RunHealthServer(ctx context.Context, grpcPort, mgmtPort string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	mux := runtime.NewServeMux()
	if err := v1.RegisterHealthHandlerFromEndpoint(ctx, mux, "localhost:"+grpcPort, opts); err != nil {
		log.Fatalf("failed to start Health endpoint %v", err)
	}


	srv := &http.Server{
		Addr:    ":" + mgmtPort,
		Handler: mux,
	}

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
		}

		_, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		_ = srv.Shutdown(ctx)
	}()

	log.Println("starting Health Endpoint")
	return srv.ListenAndServe()
}
