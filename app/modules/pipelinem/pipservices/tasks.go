package pipservices

import (
	"github.com/goatcms/goatcore/app"
	commservices "github.com/goatcms/goatcore/app/modules/commonm/commservices"
)

// Task contains single task data
type Task interface {
	// Name return task name
	Name() string
	// Logs return task result
	Logs() string
	// Done return true if task is finished
	Done() bool
	// Status return task status description
	Status() string
	// Wait for task finish
	Wait() error
	// WaitList return list of related tasks to wait for
	WaitList() []string
	// WaitLockMapList return map described related resources to lock
	LockMap() commservices.LockMap
	// Errors return task errors (or nil)
	Errors() []error
}

// TaskWriter write data to task
type TaskWriter interface {
	Task
	// IOContext return task IOContext
	IOContext() app.IOContext
	// SetStatus set task status
	SetStatus(status string)
	// Close make task done
	Close() (err error)
}

// TasksManager contains tasks data
type TasksManager interface {
	Logs() string
	Names() []string
	Get(name string) (task Task)
	Create(pip Pip) (task TaskWriter, err error)
}

// TasksUnit menage pipeline tasks
type TasksUnit interface {
	FromScope(scp app.Scope) (tasks TasksManager, err error)
	BindScope(scp app.Scope, tasks TasksManager) (err error)
	Clear(scp app.Scope) (err error)
}
