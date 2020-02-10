package noop

import (
	"context"
	"fmt"
	"github.com/covarity/echo/adapters/metadata"
	"github.com/covarity/echo/pkg/adapter"
)

type handler struct{}

// GetInfo returns the Info associated with this adapter implementation.
func GetInfo() adapter.Info {
	info := metadata.GetInfo("noop")
	info.NewBuilder = func() adapter.HandlerBuilder { return &builder{} }
	return info
}

func (*handler) Close() error { return nil }

func (*handler) HandleRequestNothing(context.Context) {
	fmt.Print("noop:HandleRequestNothing")
}

type builder struct{}

func (*builder) SetAdapterConfig(adapter.Config) {}

func (b *builder) Build(context context.Context, env adapter.Env) (adapter.Handler, error) {
	return &handler{}, nil
}
