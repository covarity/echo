package handler

import (
	"time"

	"github.com/covarity/echo/pkg/adapter"
)

const (
	defaultRetryDuration = 1 * time.Second
	defaultRetryChecks   = 10
)

// Table contains a set of instantiated and configured adapter handlers.
type Table struct {
	entries map[string]Entry

	// monitoringCtx context.Context

	strayWorkersRetryDuration time.Duration
	strayWorkersCheckRetries  int
}

// Entry in the handler table.
type Entry struct {
	// Name of the Handler
	Name string

	// Handler is the initialized Handler object.
	Handler adapter.Handler

	// AdapterName that was used to create this Entry.
	AdapterName string

	// env refers to the adapter.Env passed to the handler.
	env env
}

var emptyTable = &Table{}

// Empty returns an empty table instance.
func Empty() *Table {
	return emptyTable
}
