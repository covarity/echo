package synthetic

import (
	"context"

	"github.com/covarity/echo/pkg/adapter"
)

// Fully qualified name of the template
const TemplateName = "synthetic"

// HandlerBuilder must be implemented by adapters if they want to
// process data associated with the 'synthetic' template.
//
// Echo uses this interface to call into the adapter at configuration time to configure
// it with adapter-specific configuration as well as all template-specific type information.
type HandlerBuilder interface {
	adapter.HandlerBuilder
}

// Handler must be implemented by adapter code if it wants to
// process data associated with the 'synthetic' template.
//
type Handler interface {
	adapter.Handler

	// HandleSynthetic is called by Echo at request time to deliver to
	// to an adapter.
	HandleSynthetic(context.Context) error
}
