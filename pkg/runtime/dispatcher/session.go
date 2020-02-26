package dispatcher

import (
	"context"

	adptTmpl "github.com/covarity/echo/api/adapter/model/v1"
	"github.com/covarity/echo/pkg/adapter"
)

const queueAllocSize = 64

// session represents a call session to the Impl. It contains all the mutable state needed for handling the
// call. It is used as temporary memory location to keep ephemeral state, thus avoiding garbage creation.
type session struct {
	// owner
	impl *Impl

	// routing context for the life of this session
	rc *RoutingContext
	// input parameters that was collected as part of the call.
	ctx context.Context
	// output parameters that get collected / accumulated as results.
	checkResult adapter.CheckResult
	// The current number of activeDispatches handler dispatches.
	activeDispatches int
	destination      string
	requestState     *dispatchState
	err              error

	// The variety of the operation that is being performed.
	variety adptTmpl.TemplateVariety

	// channel for collecting states of completed dispatches.
	completed chan *dispatchState
}

func (s *session) clear() {
	s.impl = nil
	s.rc = nil
	s.ctx = nil
	s.variety = 0
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

func (s *session) ensureParallelism(minParallelism int) {
	// Resize the channel to accommodate the parallelism, if necessary.
	if cap(s.completed) < minParallelism {
		allocSize := ((minParallelism / queueAllocSize) + 1) * queueAllocSize
		s.completed = make(chan *dispatchState, allocSize)
	}
}

func (s *session) dispatch() error {

	var state *dispatchState
	destination := s.rc.Routes.GetDestination(s.variety, s.destination)
	//TODO: make this dynamic for performance or batch request scenarios
	s.ensureParallelism(10)
	state = s.impl.getDispatchState(s.ctx, destination)
	s.requestState = state

	s.dispatchToHandler(state)

	s.waitForDispatched()
	return nil
}

func (s *session) dispatchToHandler(ds *dispatchState) {
	s.activeDispatches++
	ds.session = s
	s.impl.gp.ScheduleWork(ds.invokeHandler, nil)
}

func (s *session) waitForDispatched() {
	for s.activeDispatches > 0 {
		state := <-s.completed
		s.activeDispatches--
		if state.err != nil {
			print("error occured wih dispatch %s", state.err)
		}

		s.impl.putDispatchState(state)
	}
}
