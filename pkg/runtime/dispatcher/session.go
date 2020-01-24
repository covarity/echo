package dispatcher

import (
	"context"

	"github.com/covarity/echo/pkg/adapter"
)

// session represents a call session to the Impl. It contains all the mutable state needed for handling the
// call. It is used as temporary memory location to keep ephemeral state, thus avoiding garbage creation.
type session struct {
	// owner
	impl *Impl
	// input parameters that was collected as part of the call.
	ctx context.Context
	// output parameters that get collected / accumulated as results.
	checkResult adapter.CheckResult
	// The current number of activeDispatches handler dispatches.
	activeDispatches int
	requestStates *dispatchState
	err error

	// channel for collecting states of completed dispatches.
	completed chan *dispatchState
}


func (s *session) clear() {
	s.impl = nil
	s.ctx = nil
	s.checkResult = adapter.CheckResult{}
	s.err = nil
	s.activeDispatches = 0
	exit := false
	for !exit {
		select {
		case <-s.completed:
			// log.Warn("Leaked dispatch state discovered!")
			continue
		
		default: 
			exit = true
		}
	}

}


