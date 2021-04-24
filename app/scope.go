package app

import (
	"context"
	"time"

	"github.com/goatcms/goatcore/dependency"
)

// DataScope provide data provider
type DataScope interface {
	// SetValue set the value associated with this context for key
	SetValue(key interface{}, value interface{})
	// Value returns the value associated with this context for key, or nil
	Value(key interface{}) interface{}
	// Keys return a keys to all scope values
	Keys() []interface{}
	// LockData lock the data scope and return new DataScopeLocker
	LockData() (transaction DataScopeLocker)
}

// DataScopeLocker provide data scope commitable interface
type DataScopeLocker interface {
	DataScope
	// Commit save a data and unlock the data scope
	Commit() (err error)
}

// EventScope provide event interface
type EventScope interface {
	// Trigger an event
	Trigger(id interface{}, data interface{}) error
	// On add an event listener
	On(id interface{}, callback EventCallback)
}

// ErrorScope provide error interface
type ErrorScope interface {
	// Errors return the scope errors as an array
	Errors() []error
	// Err return cumulative error if the scope context contains any error
	Err() error
	// AppendError add an error to the scope
	AppendError(err ...error)
}

// ContextScope provide sync interface
type ContextScope interface {
	ErrorScope
	// Deadline returns the time when work done on behalf of this context
	// should be canceled. Deadline returns ok==false when no deadline is
	// set. Successive calls to Deadline return the same results.
	Deadline() (deadline time.Time, ok bool)
	// Done is close when the scope context is done (kill or stop)
	Done() <-chan struct{}
	// IsDone check if the scope context is done (kill or stop)
	IsDone() bool
	// Kill stop the scope context with error
	Kill()
	// Stop stop the scope context without error
	Stop()
	// Wait until the scope context is done and return error
	Wait() error
	// AddTasks add a tasks and return an error if too many gorutines
	AddTasks(delta int) (err error)
	// DoneTask mark single task as done
	DoneTask()
}

// Scope is global scope interface
type Scope interface {
	DataScope
	EventScope
	ContextScope
	dependency.Injector

	// GoContext convert scope to a golang context (context.Context)
	GoContext() context.Context
	// Close the scope. Wait to finish and mark scope as done.
	Close() (err error)
}
