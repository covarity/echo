package adapter

type Info struct {
	Name        string
	Impl        string
	Description string
	// NewBuilder is a function that creates a Builder which implements Builders associated
	// with the SupportedTemplates.
	NewBuilder NewBuilderFn

	// SupportedTemplates expresses all the templates the Adapter wants to serve.
	SupportedTemplates []string
}

// NewBuilderFn is a function that creates a Builder.
type NewBuilderFn func() HandlerBuilder

type InfoFn func() Info
