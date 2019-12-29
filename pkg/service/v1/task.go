package v1

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	v1 "github.com/covarity/echo/pkg/api/v1"
	"github.com/covarity/echo/pkg/queue"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

// taskServiceServer is implementation of v1.taskServiceServer proto interface
type taskServiceServer struct {
	queue queue.Queue
}

// NewToDoServiceServer creates task service
func NewTaskServiceServer(q *queue.Queue) v1.TaskServiceServer {
	return &taskServiceServer{}
}

// checkAPI checks if the API version requested by client is supported by server
func (s *taskServiceServer) checkAPI(api string) error {
	// API version is "" means use current version of the service
	if len(api) > 0 {
		if apiVersion != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
		}
	}
	return nil
}

// Create new task
func (s *taskServiceServer) Create(ctx context.Context, req *v1.CreateRequest) (*v1.CreateResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	fmt.Printf(req.Task.GetProtocol().String())

	// reminder, err := ptypes.Timestamp(req.Task.Reminder)
	// if err != nil {
	// 	return nil, status.Error(codes.InvalidArgument, "reminder field has invalid format-> "+err.Error())
	// }

	// fmt.Printf("Task:Create:reminder:%s", reminder)

	s.queue.PushBack(queue.Item{Value: req.Task.GetProtocol(), Priority: 0})

	var id int64 = 1

	return &v1.CreateResponse{
		Api: apiVersion,
		Id:  id,
	}, nil
}

// Delete task
func (s *taskServiceServer) Delete(ctx context.Context, req *v1.DeleteRequest) (*v1.DeleteResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	return &v1.DeleteResponse{
		Api:     apiVersion,
		Deleted: 1,
	}, nil
}

// Update task
func (s *taskServiceServer) Update(ctx context.Context, req *v1.UpdateRequest) (*v1.UpdateResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	return &v1.UpdateResponse{
		Api:     apiVersion,
		Updated: 1,
	}, nil
}

// Read task
func (s *taskServiceServer) Read(ctx context.Context, req *v1.ReadRequest) (*v1.ReadResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	return &v1.ReadResponse{
		Api:  apiVersion,
		Task: &v1.Task{Title: "example"},
	}, nil

}

// Read all tasks
func (s *taskServiceServer) ReadAll(ctx context.Context, req *v1.ReadAllRequest) (*v1.ReadAllResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	title := s.queue.String()

	return &v1.ReadAllResponse{
		Api:  apiVersion,
		Task: []*v1.Task{&v1.Task{Title: title}},
	}, nil
}
