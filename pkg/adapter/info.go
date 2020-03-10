package adapter

import (
	"github.com/gogo/protobuf/proto"
)

type Info struct {
	Name        string
	Impl        string
	Description string
	// NewBuilder is a function that creates a Builder which implements Builders associated
	// with the SupportedTemplates.
	NewBuilder NewBuilderFn

	// SupportedTemplates expresses all the templates the Adapter wants to serve.
	SupportedTemplates []string

	// DefaultConfig is a default configuration struct for this
	// adapter. This will be used by the configuration system to establish
	// the shape of the block of configuration state passed to the HandlerBuilder.Build method.
	DefaultConfig proto.Message
}

// NewBuilderFn is a function that creates a Builder.
type NewBuilderFn func() HandlerBuilder

type InfoFn func() Info
