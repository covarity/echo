package routing

import (
	adptTmpl "github.com/covarity/echo/api/adapter/model/v1"
	"github.com/covarity/echo/pkg/runtime/handler"
	"github.com/covarity/echo/pkg/template"
	"github.com/covarity/echo/pkg/adapter"
)

type builder struct {
	table     *Table
	handlers  *handler.Table
	adapters  map[string]*adapter.Info
	templates map[string]*template.Info
}

func BuildTable(
	handlers *handler.Table,
	adapters map[string]*adapter.Info,
	templates map[string]*template.Info,
) *Table {
	b := &builder{
		table: &Table{
			entries: make(map[adptTmpl.TemplateVariety]*varietyTable, 3),
		},
		handlers:  handlers,
		adapters:  adapters,
		templates: templates,
	}
	b.build()
	return b.table
}

func (b *builder) build() {
	for _, entry := range b.handlers.Entries() {
		for _, templateInfo := range b.templates {
			b.add(buildTemplateInfo(templateInfo), entry)
		}
	}
}

// templateInfo build method needed dispatch this template
// func (b *builder) templateInfo(tmpl *config.Template) *TemplateInfo {
// 	ti := &TemplateInfo{
// 		Name:    tmpl.Name,
// 		Variety: tmpl.Variety,
// 	}
// }

func (b *builder) add(
		t *TemplateInfo,
		entry handler.Entry,
	) {
		variety := t.Variety
		byVariety, found := b.table.entries[variety]
		if !found {
			byVariety = &varietyTable{
				entries: make(map[string]*Destination),
			}
			b.table.entries[variety] = byVariety
		}

		var byHandler *Destination
		for _, d := range byVariety.Entries() {
			if d.HandlerName == entry.Name && d.Template.Name == t.Name {
				byHandler = d
				break
			}
		}
		if byHandler == nil {
			byHandler = &Destination{
				Handler:     entry.Handler,
				HandlerName: entry.Name,
				AdapterName: entry.AdapterName,
				Template:    t,	
			}
			byVariety.entries[entry.Name] = byHandler
		}
	}
