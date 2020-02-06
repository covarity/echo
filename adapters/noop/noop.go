package noop

import (
	"github.com/covarity/echo/adapters/metadata"
	"fmt"
	"context"
	"github.com/covarity/echo/pkg/adapter"
)

type handler struct{}

// GetInfo returns the Info associated with this adapter implementation.
func GetInfo() adapter.Info {
	info := metadata.GetInfo("noop")
	return info
}

func (*handler) Close() error { return nil }

func (*handler) HandleRequestNothing(context.Context) {
	fmt.Print("noop:HandleRequestNothing")
}
