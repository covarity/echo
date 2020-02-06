package runtime

import (
	"sync"

	"github.com/covarity/echo/pkg/pool"
	"github.com/covarity/echo/pkg/runtime/dispatcher"
	"github.com/covarity/echo/pkg/runtime/handler"
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
	// adapters map[string]*adapter.Info,
	executorPool *pool.GoroutinePool,
	handlerPool *pool.GoroutinePool) *Runtime {

	rt := &Runtime{
		handlers:    handler.Empty(),
		dispatcher:  dispatcher.New(executorPool),
		handlerPool: handlerPool,
	}

	return rt
}

// Dispatcher returns the dispatcher.Dispatcher that is implemented by this runtime package.
func (c *Runtime) Dispatcher() dispatcher.Dispatcher {
	return c.dispatcher
}
