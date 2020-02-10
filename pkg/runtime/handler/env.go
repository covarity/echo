package handler

import (
	"fmt"

	"github.com/covarity/echo/pkg/adapter"
	"github.com/covarity/echo/pkg/pool"
)

type env struct {
	// logger           adapter.Logger
	gp *pool.GoroutinePool
	// monitoringCtx    context.Context
	daemons, workers *int64
}

// NewEnv returns a new environment instance.
func NewEnv(name string, gp *pool.GoroutinePool) adapter.Env {
	return env{
		// logger:        newLogger(name),
		gp:      gp,
		daemons: new(int64),
		workers: new(int64),
	}
}

// Logger from adapter.Env.
// func (e *env) Logger() adapter.Logger {
// return e.logger
// }

// ScheduleWork from adapter.Env.
func (e env) ScheduleWork(fn adapter.WorkFunc) {

	e.gp.ScheduleWork(func(ifn interface{}) {
		reachedEnd := false

		defer func() {

			if !reachedEnd {
				r := recover()
				fmt.Printf("Adapter worker failed: %v", r)
				// _ = e.Logger().Errorf("Adapter worker failed: %v", r) //
			}
		}()

		ifn.(adapter.WorkFunc)()
		reachedEnd = true
	}, fn)
}

// ScheduleDaemon from adapter.Env.
func (e env) ScheduleDaemon(fn adapter.DaemonFunc) {

	go func() {
		reachedEnd := false

		defer func() {

			if !reachedEnd {
				r := recover()
				fmt.Printf("Adapter daemon failed: %v", r)
				// _ = e.Logger().Errorf("Adapter worker failed: %v", r) //
			}
		}()

		fn()
		reachedEnd = true
	}()
}
