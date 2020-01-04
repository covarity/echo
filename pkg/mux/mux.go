package mux

import (
	"github.com/covarity/echo/pkg/pqueue"
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
	for _, protocol := range m.Protocols {
		m.queues[protocol] = queue.New()
	}
	return m
}
