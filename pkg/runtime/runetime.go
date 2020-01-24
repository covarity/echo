package runtime


import (
	"errors"
	"sync"
	"time"
	"github.com/covarity/echo/pkg/adapter"
	"github.com/covarity/echo/pkg/runtime/dispatcher"
	"github.com/covarity/echo/pkg/runtime/handler"
	"github.com/covarity/echo/pkg/pool"
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