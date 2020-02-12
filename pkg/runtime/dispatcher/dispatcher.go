package dispatcher

import (
	"context"
	"sync"
	"time"

	"github.com/covarity/echo/pkg/adapter"
	"github.com/covarity/echo/pkg/pool"
	"github.com/covarity/echo/pkg/runtime/routing"
)

var (
	protocols = []string{"TCP", "UDP", "HTTP", "GRPC"}
)

// Dispatcher dispatches incoming API calls to configured adapters.
type Dispatcher interface {

	// Check dispatches to the set of adapters associated with the Check API method
	Check(ctx context.Context, destination string) (adapter.CheckResult, error)

	// GetRequester get an interface where reports are buffered.
	GetRequester(ctx context.Context) Requester
}

// Impl is the runtime implementation of the Dispatcher interface.
type Impl struct {

	// Current routing context.
	rc *RoutingContext
	// the reader-writer lock for accessing or changing the context.
	rcLock sync.RWMutex

	// pool of sessions
	sessionPool sync.Pool

	// pool of dispatch states
	statePool sync.Pool

	// pool of reporters
	requesterPool sync.Pool

	// pool of goroutines
	gp *pool.GoroutinePool
}

var _ Dispatcher = &Impl{}

// New returns a new Impl instance. The Impl instance is initialized with an empty routing table.
func New(handlerGP *pool.GoroutinePool) *Impl {
	d := &Impl{
		gp: handlerGP,
		rc: &RoutingContext{
			Routes: routing.Empty(),
		},
	}

	d.sessionPool.New = func() interface{} { return &session{} }
	d.statePool.New = func() interface{} { return &dispatchState{} }
	d.requesterPool.New = func() interface{} { return &requester{state: &dispatchState{}} }
	return d
}

const (
	defaultValidDuration = 1 * time.Minute
	defaultValidUseCount = 10000
)

// Check implementation of runtime.Impl.
func (d *Impl) Check(ctx context.Context, destination string) (adapter.CheckResult, error) {
	s := d.getSession(ctx, destination)

	var r adapter.CheckResult
	err := s.dispatch()
	if err == nil {
		r = s.checkResult
		err = s.err

		if err == nil {
			// No adapters chimed in on this request, so we return a "good to go" value which can be cached
			// for up to a minute.
			//
			// TODO: make these fallback values configurable
			if r.IsDefault() {
				r = adapter.CheckResult{
					ValidUseCount: defaultValidUseCount,
					ValidDuration: defaultValidDuration,
				}
			}
		}
	}
	d.putSession(s)
	return r, err
}

// GetRequester ...
func (d *Impl) GetRequester(ctx context.Context) Requester {
	return d.getRequester(ctx)
}

// Session template variety is CHECK for output producing templates (CHECK_WITH_OUTPUT)
func (d *Impl) getSession(context context.Context, destination string) *session {
	s := d.sessionPool.Get().(*session)
	s.rc = d.acquireRoutingContext()
	s.destination = destination
	s.impl = d
	s.ctx = context

	return s
}

func (d *Impl) putSession(s *session) {
	s.clear()
	d.sessionPool.Put(s)
}

func (d *Impl) getDispatchState(context context.Context, destination *routing.Destination) *dispatchState {
	ds := d.statePool.Get().(*dispatchState)
	ds.destination = destination
	ds.ctx = context

	return ds
}

func (d *Impl) putDispatchState(ds *dispatchState) {
	ds.clear()
	d.statePool.Put(ds)
}

func (d *Impl) getRequester(context context.Context) *requester {
	r := d.requesterPool.Get().(*requester)

	r.impl = d
	r.ctx = context

	return r
}

func (d *Impl) putRequester(r *requester) {
	r.clear()
	d.requesterPool.Put(r)
}

func (d *Impl) acquireRoutingContext() *RoutingContext {
	d.rcLock.RLock()
	rc := d.rc
	rc.incRef()
	d.rcLock.RUnlock()

	return rc
}

// ChangeRoute changes the routing table on the Impl which, in turn, ends up creating a new RoutingContext.
func (d *Impl) ChangeRoute(newTable *routing.Table) *RoutingContext {
	newRC := &RoutingContext{
		Routes: newTable,
	}

	d.rcLock.Lock()
	old := d.rc
	d.rc = newRC
	d.rcLock.Unlock()

	return old
}
