package adapters

import (
	http "github.com/covarity/echo/adapters/http"
	"github.com/covarity/echo/adapters/noop"
	tcp "github.com/covarity/echo/adapters/tcp"
	"github.com/covarity/echo/pkg/adapter"
)

func Inventory() []adapter.InfoFn {
	return []adapter.InfoFn{
		tcp.GetInfo,
		http.GetInfo,
		noop.GetInfo,
	}
}
