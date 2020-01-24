package adapter

import (
	"io"
)

type (
	// Handler represents default functionality every Adapter must implement.
	Handler interface {
		io.Closer
	}
)
