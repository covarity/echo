package routing

import (
	"github.com/covarity/echo/pkg/adapter"
	"github.com/covarity/echo/pkg/runtime/handler"
)

type builder struct {
	table    *Table
	handlers *handler.Table
	adapters map[string]*adapter.Info
}

func BuildTable(
	handlers *handler.Table,
	adapters map[string]*adapter.Info,
) *Table {
	b := &builder{
		table: &Table{
			entries: make(map[string]*Destination, handlers.Len()),
		},
		handlers: handlers,
		adapters: adapters,
	}
	b.build()
	return b.table
}

func (b *builder) build() {
	for handler, entry := range b.handlers.Entries() {
		b.table.entries[handler] = &Destination{
			Handler: entry.Handler,
			HandlerName: handler,
			AdapterName: entry.AdapterName,
		}
	}
}
