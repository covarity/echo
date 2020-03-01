package pool

import (
	"sync"
)

// References
// 1. https://www.ardanlabs.com/blog/2013/09/pool-go-routines-to-process-task.html

// WorkFunc is a function that will be called by a worker in the pool
type WorkFunc func(param interface{})

// GoroutinePool represents a set of reusable goroutines onto which work can be scheduled.
type GoroutinePool struct {
	queue          chan work      // Channel providing the work that needs to be executed
	wg             sync.WaitGroup // Used to block shutdown until all workers complete
	singleThreaded bool           // Whether to actually use goroutines or not
}

type work struct {
	fn    WorkFunc
	param interface{}
}

// New creates a new pool of goroutines to schedule async work.
func New(queueDepth int, singleThreaded bool) *GoroutinePool {
	return &GoroutinePool{
		queue:          make(chan work, queueDepth),
		singleThreaded: singleThreaded,
	}
}

// Close waits for all goroutines to terminate (and implements io.Closer).
func (gp *GoroutinePool) Close() error {
	if !gp.singleThreaded {
		close(gp.queue)
		gp.wg.Wait()
	}
	return nil
}

// ScheduleWork registers the given function to be executed at some point. The given param will
// be supplied to the function during execution.
func (gp *GoroutinePool) ScheduleWork(fn WorkFunc, param interface{}) {
	if gp.singleThreaded {
		fn(param)
	} else {
		gp.queue <- work{fn: fn, param: param}
	}
}

// AddWorkers introduces more goroutines in the worker pool, increasing potential parallelism.
func (gp *GoroutinePool) AddWorkers(numWorkers int) {
	if !gp.singleThreaded {
		gp.wg.Add(numWorkers)
		for i := 0; i < numWorkers; i++ {
			go func() {
				for work := range gp.queue {
					work.fn(work.param)
				}
				gp.wg.Done()
			}()
		}
	}
}
