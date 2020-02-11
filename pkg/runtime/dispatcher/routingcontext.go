package dispatcher

import (
	"sync/atomic"

	"github.com/covarity/echo/pkg/runtime/routing"
)

// RoutingContext is the currently active dispatching context, based on a config.
type RoutingContext struct {
	// the routing table of this context.
	Routes *routing.Table

	// the current reference count. Indicates how many calls are currently using this RoutingContext.
	refCount int32
}

// incRef increases the reference count on the RoutingContext.
func (rc *RoutingContext) incRef() {
	atomic.AddInt32(&rc.refCount, 1)
}

// decRef decreases the reference count on the RoutingContext.
func (rc *RoutingContext) decRef() {
	atomic.AddInt32(&rc.refCount, -1)
}

// GetRefs returns the current reference count on the dispatch context.
func (rc *RoutingContext) GetRefs() int32 {
	return atomic.LoadInt32(&rc.refCount)
}
