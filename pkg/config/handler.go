package config

import (
	"context"
	"fmt"

	"github.com/covarity/echo/pkg/adapter"
)

// BuildHandler instantiates a handler object using the passed in handler and instances configuration.
func BuildHandler(handler *HandlerStatic, env adapter.Env) (h adapter.Handler, err error) {
	var builder adapter.HandlerBuilder
	// Adapter should always be present for a valid configuration (reference integrity should already be checked).
	info := handler.Adapter
	builder = info.NewBuilder()
	h, err = buildHandler(builder, env)

	if err != nil {
		h = nil
		err = fmt.Errorf("adapter instantiation error: %v", err)
		return
	}

	return h, nil
}

func buildHandler(builder adapter.HandlerBuilder, env adapter.Env) (handler adapter.Handler, err error) {
	return builder.Build(context.Background(), env)
}
