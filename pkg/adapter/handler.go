package adapter

import (
	"context"
	"io"
)

type (
	// Handler represents default functionality every Adapter must implement.
	Handler interface {
		io.Closer
	}

	HandlerBuilder interface {
		// SetAdapterConfig gives the builder the adapter-level configuration state.
		SetAdapterConfig(Config)
		Build(context.Context, Env) (Handler, error)
	}
)
