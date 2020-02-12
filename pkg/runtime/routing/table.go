package routing

import (
	"github.com/covarity/echo/pkg/adapter"
)

// Destination for a given dispatch
type Destination struct {
	id          uint32
	Handler     adapter.Handler
	HandlerName string
	AdapterName string
}

// Table of Destinations for various adapter & handlers
type Table struct {
	entries map[string]*Destination
}

// Destination retrieval
func (t *Table) Destination(name string) *Destination {
	return t.entries[name]
}

var emptyTable = &Table{}

// Empty returns an empty routing table.
func Empty() *Table {
	return emptyTable
}

func (t *Table) GetDestination(name string) *Destination {

	destination := t.entries[name]
	if destination == nil {
		destination = &Destination{}
	}

	return destination
}
