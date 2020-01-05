package mux

import (
	queue "github.com/covarity/echo/pkg/pqueue"
	"fmt"
)

var (
	protocols = []string{"TCP", "UDP", "HTTP", "GRPC"}
)

// Mux - Responsible for routing new tasks into the corresponding work queue and surfacing a consumable interface for handlers/adapters
type Mux struct {
	Protocols []string
	events    chan queue.Item
	queues    map[string]*queue.Queue
}

// New returns an initialized Mux instance
func New() *Mux {
	return new(Mux).Init()
}

// Init initializes Mux instance
func (m *Mux) Init() *Mux {
	m.Protocols = protocols
	m.queues = make(map[string]*queue.Queue)
	for _, protocol := range m.Protocols {
		m.queues[protocol] = queue.New()
	}
	return m
}

func (m *Mux) AddTask(x interface{}, protocol string) {
	fmt.Printf("adding task:protocol:%s:x:%v",protocol, x)
	m.queues[protocol].Enqueue(queue.Item{Value: x})
}

func (m *Mux) ListTasks(protocol string) string {
	return m.queues[protocol].String()
}
