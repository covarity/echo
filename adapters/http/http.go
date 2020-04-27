package tcp

import (
	"context"
	"fmt"
	"time"

	"github.com/covarity/echo/adapters/metadata"
	"github.com/covarity/echo/pkg/adapter"
	"github.com/covarity/echo/templates/synthetic"
)

// Handler implements tracespan.Handler using an OpenCensus TraceExporter.
type Handler struct {
	// CloseFunc will be called when this handler is closed. Optional.
	CloseFunc func() error
}

type handler struct{}

var _ synthetic.Handler = &handler{}

var checkResult = adapter.CheckResult{
	Status:        "OK",
	ValidDuration: 1000000000 * time.Second,
	ValidUseCount: 1000000000,
}

// GetInfo returns the Info associated with this adapter implementation.
func GetInfo() adapter.Info {
	info := metadata.GetInfo("http")
	info.NewBuilder = func() adapter.HandlerBuilder { return &builder{} }
	return info
}

func (*handler) Close() error { return nil }

func (*handler) HandleRequest(context.Context) error {
	return nil
}

func (*handler) HandleSynthetic(context.Context) error {
	// TODO: add http synethetic check logic here
	fmt.Printf("Adapter:HTTP:HandleSynthetic")
	return nil
}

type builder struct{}

var _ synthetic.HandlerBuilder = &builder{}

func (*builder) SetAdapterConfig(adapter.Config) {}

func (b *builder) Build(context context.Context, env adapter.Env) (adapter.Handler, error) {
	return &handler{}, nil
}
