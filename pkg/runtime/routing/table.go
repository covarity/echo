package routing

import "github.com/covarity/echo/pkg/adapter"

// Destination for a given dispatch
type Destination struct {
	id uint32
	Handler adapter.Handler
	HandlerName string
	AdapterName string
}