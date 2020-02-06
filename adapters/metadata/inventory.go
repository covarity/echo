package metadata

import (
	"fmt"

	"github.com/covarity/echo/pkg/adapter"
)

var (
	Infos = []adapter.Info{
		{
			Name:        "noop",
			Impl:        "github.com/covarity/echo/adapter/noop",
			Description: "Does nothing (useful for testing)",
		},
		{
			Name:        "tcp",
			Impl:        "github.com/covarity/echo/adapter/tcp",
			Description: "TCP based interactions",
		},
	}
)

// GetInfo looks up an adapter info from the declaration list by name
func GetInfo(name string) adapter.Info {
	for _, info := range Infos {
		if info.Name == name {
			return info
		}
	}
	panic(fmt.Errorf("requesting a missing descriptor %q", name))
}
