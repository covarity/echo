package dispatcher

import (
	"sync"

	"github.com/covarity/echo/pkg/pool"
)

var (
	protocols = []string{"TCP", "UDP", "HTTP", "GRPC"}
)

// Impl is the runtime implementation of the Dispatcher interface.
type Impl struct {

	// the reader-writer lock for accessing or changing the context.
	rcLock sync.RWMutex

	// pool of sessions
	sessionPool sync.Pool

	// pool of dispatch states
	statePool sync.Pool

	// pool of reporters
	reporterPool sync.Pool

	// pool of goroutines
	gp *pool.GoroutinePool
}

// New returns a new Impl instance. The Impl instance is initialized with an empty routing table.
func New(handlerGP *pool.GoroutinePool) *Impl {
	d := &Impl{
		gp: handlerGP,
	}

	d.sessionPool.New = func() interface{} { return &session{} }
	d.statePool.New = func() interface{} { return &dispatchState{} }
	return d
}

// Mux - Responsible for routing new tasks into the corresponding work queue and surfacing a consumable interface for handlers/adapters
// type Mux struct {
// 	Protocols []string
// 	events    chan queue.Item
// 	queues    map[string]*queue.Queue
// }

// // New returns an initialized Mux instance
// func New() *Mux {
// 	return new(Mux).Init()
// }

// // Init initializes Mux instance
// func (m *Mux) Init() *Mux {
// 	m.Protocols = protocols
// 	m.queues = make(map[string]*queue.Queue)
// 	for _, protocol := range m.Protocols {
// 		m.queues[protocol] = queue.New()
// 	}
// 	return m
// }

// func (m *Mux) AddTask(x interface{}, protocol string) {
// 	fmt.Printf("adding task:protocol:%s:x:%v", protocol, x)
// 	m.queues[protocol].Enqueue(queue.Item{Value: x})
// }

// func (m *Mux) ListTasks(protocol string) string {
// 	return m.queues[protocol].String()
// }
