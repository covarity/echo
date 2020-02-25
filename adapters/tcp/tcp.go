package tcp

import (
	"context"
	"fmt"
	"time"

	"github.com/covarity/echo/adapters/metadata"
	"github.com/covarity/echo/pkg/adapter"
)

type handler struct{}

var checkResult = adapter.CheckResult{
	Status:        "OK",
	ValidDuration: 1000000000 * time.Second,
	ValidUseCount: 1000000000,
}

// GetInfo returns the Info associated with this adapter implementation.
func GetInfo() adapter.Info {
	info := metadata.GetInfo("tcp")
	info.NewBuilder = func() adapter.HandlerBuilder { return &builder{} }
	return info
}

func (*handler) Close() error { return nil }

func (*handler) HandleRequest(context.Context) error {
	fmt.Println("handling tcp request")
	return nil
}

type builder struct{}

func (*builder) SetAdapterConfig(adapter.Config) {}

func (b *builder) Build(context context.Context, env adapter.Env) (adapter.Handler, error) {
	return &handler{}, nil
}
