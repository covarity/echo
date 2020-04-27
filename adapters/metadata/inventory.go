package metadata

import (
	"fmt"
	"time"

	tcp "github.com/covarity/echo/adapters/tcp/config"
	http "github.com/covarity/echo/adapters/http/config"
	"github.com/covarity/echo/pkg/adapter"
	"github.com/covarity/echo/templates/synthetic"
)

var (
	Infos = []adapter.Info{
		{
			Name:        "noop",
			Impl:        "github.com/covarity/echo/adapter/noop",
			Description: "Does nothing (useful for testing)",
			SupportedTemplates: []string{
				synthetic.TemplateName,
			},
			DefaultConfig: &tcp.Params{Timeout: 10 * time.Second},
		},
		{
			Name:        "tcp",
			Impl:        "github.com/covarity/echo/adapter/tcp",
			Description: "TCP based interactions",
			SupportedTemplates: []string{
				synthetic.TemplateName,
			},
		},
		{
			Name:        "http",
			Impl:        "github.com/covarity/echo/adapter/http",
			Description: "HTTP based interactions",
			SupportedTemplates: []string{
				synthetic.TemplateName,
			},
			DefaultConfig: &http.Params{Timeout: 10 * time.Second},
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
