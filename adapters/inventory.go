package adapters

import (
	tcp "github.com/covarity/echo/adapters/tcp"
	"github.com/covarity/echo/adapters/noop"
	"github.com/covarity/echo/pkg/adapter"
)

func Inventory() []adapter.InfoFn {
	return []adapter.InfoFn{
		tcp.GetInfo,
		noop.GetInfo,
	}
}