package v1

import (
	"context"

	v1 "github.com/covarity/echo/pkg/api/v1"
)

const (
	// apiVersion is version of API is provided by server
	healthApiVersion = "v1"
)

// taskServiceServer is implementation of v1.taskServiceServer proto interface
type healthServer struct {
}

// NewToDoServiceServer creates task service
func NewHealthServer() v1.HealthServer {
	return &healthServer{}
}

func (s *healthServer) Check(ctx context.Context, req *v1.HealthCheckRequest) (*v1.HealthCheckResponse, error) {
	return &v1.HealthCheckResponse{Status: v1.HealthCheckResponse_SERVING}, nil
}
