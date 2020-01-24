// Package adapter defines the types consumed by adapter implementations to
// interface with Echo
package adapter

import (
	"github.com/gogo/protobuf/proto"
)


type (
	// Config represents a chunk of adapter configuration state
	Config proto.Message

	// WorkFunc represents a function to invoke.
	WorkFunc func()

	// DaemonFunc represents a function to invoke asynchronously to run a long-running background processing loop.
	DaemonFunc func()

	// Env defines the environment in which an adapter executes.
	Env interface {
		// Logger returns the logger for the adapter to use at runtime.
		//TODO: Logger() Logger

		// ScheduleWork records a function for execution.
		//
		// Under normal circumstances, this method executes the
		// function on a separate goroutine. But when Echo is
		// running in single-threaded mode, then the function
		// will be invoked synchronously on the same goroutine.
		//
		// Adapters should not spawn 'naked' goroutines, they should
		// use this method or ScheduleDaemon instead.
		ScheduleWork(fn WorkFunc)

		// ScheduleDaemon records a function for background execution.
		// Unlike ScheduleWork, this method guarantees execution on a
		// different goroutine. Use ScheduleDaemon for long-running
		// background operations, whereas ScheduleWork is for bursty
		// one-shot kind of things.
		//
		// Adapters should not spawn 'naked' goroutines, they should
		// use this method or ScheduleWork instead.
		ScheduleDaemon(fn DaemonFunc)
	}

	//TODO: logger interface 
	// Logger defines where adapters should output their log state to.
	//
	// This log is funneled to Echo which augments it with
	// desirable metadata and then routes it to the right place.
	// Logger interface {
	// }
)