package tcp

import (
	"context"
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
	return info
}

func (*handler) HandleRequest(context.Context) error {
	return nil
}