package config

import (
	"github.com/covarity/echo/pkg/adapter"
	"github.com/gogo/protobuf/proto"
)

type (
	// HandlerStatic configuration for compiled in adapters. Fully resolved.
	HandlerStatic struct {

		// Name of the Handler. Fully qualified.
		Name string

		// Associated adapter. Always resolved.
		Adapter *adapter.Info

		// parameters used to construct the Handler.
		Params proto.Message
	}
)

// GetName gets name
func (h HandlerStatic) GetName() string {
	return h.Name
}

// AdapterName gets adapter name
func (h HandlerStatic) AdapterName() string {
	return h.Adapter.Name
}

// AdapterParams gets AdapterParams
func (h HandlerStatic) AdapterParams() interface{} {
	return h.Params
}
