package handler

import (
	"fmt"
	"time"

	"github.com/covarity/echo/pkg/adapter"
	"github.com/covarity/echo/pkg/pool"
	"github.com/covarity/echo/pkg/config"
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

func NewTable(adapters map[string]*adapter.Info, gp *pool.GoroutinePool) *Table {
	t := &Table{
		entries:                   make(map[string]Entry, len(adapters)),
		strayWorkersCheckRetries:  defaultRetryChecks,
		strayWorkersRetryDuration: defaultRetryDuration,
	}
	for _, info := range adapters {
		createEntry(t, info, gp)
	}
	return t
}

func createEntry(t *Table, info *adapter.Info, gp *pool.GoroutinePool) {
	e := NewEnv(info.Name, gp).(env)
	statichandler := &config.HandlerStatic{
		Name:    info.Name,
		Adapter: info,
		Params:  nil,
	}
	handler, err := config.BuildHandler(statichandler, e)
	if err != nil {
		fmt.Errorf("unable to initilize adapter: handler='%s'", statichandler.GetName())
	}
	t.entries[info.Name] = Entry{
		Name:        statichandler.GetName(),
		Handler:     handler,
		AdapterName: statichandler.AdapterName(),
		env:         e,
	}
}

func (t *Table) Len() int {
	return len(t.entries)
}

// Get returns the entry for a Handler with the given name, if it exists.
func (t *Table) Get(handlerName string) (Entry, bool) {
	e, found := t.entries[handlerName]
	if !found {
		return Entry{}, false
	}

	return e, true
}

func (t *Table) Entries() map[string]Entry {
	return t.entries
}
