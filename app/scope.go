package app

import (
	"context"

	"github.com/goatcms/goatcore/dependency"
)

// DataScope provide data provider
type DataScope interface {
	Set(string, interface{}) error
	Get(string) (interface{}, error)
	Keys() ([]string, error)
	LockData() (transaction DataScopeLocker)
}

// DataScopeLocker provide data scope commitable interface
type DataScopeLocker interface {
	DataScope
	Commit() (err error)
}

// EventScope provide event interface
type EventScope interface {
	Trigger(int, interface{}) error
	On(int, EventCallback)
}

// ErrorScope provide error interface
type ErrorScope interface {
	Errors() []error
	ToError() error
	AppendError(err error)
	AppendErrors(err ...error)
}

// SyncScope provide sync interface
type SyncScope interface {
	ErrorScope

	Context() context.Context
	Kill()
	IsKilled() bool
	Wait() error
	AddTasks(delta int) (err error)
	DoneTask()
}

// Scope is global scope interface
type Scope interface {
	DataScope
	EventScope
	SyncScope
	dependency.Injector

	Close() (err error)
}
