package routing

import (
	"fmt"
	adptTmpl "github.com/covarity/echo/api/adapter/model/v1"
	"github.com/covarity/echo/pkg/adapter"
	"github.com/covarity/echo/pkg/template"
)

// Destination for a given dispatch
type Destination struct {
	id          uint32
	Handler     adapter.Handler
	HandlerName string
	AdapterName string
	Template    *TemplateInfo
}

// Table of Destinations for various adapter & handlers
type Table struct {
	entries map[adptTmpl.TemplateVariety]*varietyTable
}

// varietyTable contains destination sets for a given template variety. It contains a mapping from namespaces
// to a flattened list of destinations. It also contains the defaultSet, which gets returned if no namespace-specific
// destination entry is found.
type varietyTable struct {
	// destinations
	entries map[string]*Destination
}

var emptyTable = &Table{}

// Empty returns an empty routing table.
func Empty() *Table {
	return emptyTable
}

func (t *Table) GetDestination(variety adptTmpl.TemplateVariety, name string) *Destination {

	destinations, ok := t.entries[variety]

	if !ok {
		fmt.Printf("No destinations found for variety: variety='%d'", variety)

		return &Destination{}
	}
	destinationSet := destinations.entries[name]
	if destinationSet == nil {
		fmt.Printf("no destination found for %s", name)
		destinationSet = &Destination{}
	}

	return destinationSet
}

// Entries in the table.
func (d *varietyTable) Entries() map[string]*Destination {
	return d.entries
}

// TemplateInfo is the common data that is needed from a template
type TemplateInfo struct {
	Name            string
	Variety         adptTmpl.TemplateVariety
	DispatchRequest template.DispatchRequestFn
}

func buildTemplateInfo(info *template.Info) *TemplateInfo {
	return &TemplateInfo{
		Name:            info.Name,
		Variety:         info.Variety,
		DispatchRequest: info.DispatchRequest,
	}
}
