package runtime

import (
	"sync"

	"github.com/covarity/echo/pkg/adapter"
	"github.com/covarity/echo/pkg/pool"
	"github.com/covarity/echo/pkg/runtime/dispatcher"
	"github.com/covarity/echo/pkg/runtime/handler"
	"github.com/covarity/echo/pkg/runtime/routing"
	"github.com/covarity/echo/pkg/template"
)

// Runtime is the main entry point to the Mixer runtime environment. It listens to configuration, instantiates handler
// instances, creates the dispatch machinery and handles incoming requests.
type Runtime struct {
	handlers *handler.Table

	dispatcher *dispatcher.Impl

	handlerPool *pool.GoroutinePool

	stateLock            sync.Mutex
	shutdown             chan struct{}
	waitQuiesceListening sync.WaitGroup
}

// New returns a new instance of Runtime.
func New(
	templates map[string]*template.Info,
	adapters map[string]*adapter.Info,
	executorPool *pool.GoroutinePool,
	handlerPool *pool.GoroutinePool) *Runtime {
	handlers := handler.NewTable(adapters, handlerPool)
	rt := &Runtime{
		handlers:    handlers,
		dispatcher:  dispatcher.New(executorPool),
		handlerPool: handlerPool,
	}

	routes := routing.BuildTable(handlers, adapters, templates)

	rt.dispatcher.ChangeRoute(routes)

	return rt
}

// Dispatcher returns the dispatcher.Dispatcher that is implemented by this runtime package.
func (c *Runtime) Dispatcher() dispatcher.Dispatcher {
	return c.dispatcher
}
