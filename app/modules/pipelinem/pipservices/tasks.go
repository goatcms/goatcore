package pipservices

import (
	"github.com/goatcms/goatcore/app"
	commservices "github.com/goatcms/goatcore/app/modules/commonm/commservices"
)

// Task contains single task data
type Task interface {
	// Name return task name
	Name() string
	// FullName return name with task namespace as prefix
	FullName() string
	// Description return task description
	Description() string
	// OBroadcast write task output (and error output) to broadcast
	OBroadcast() app.BufferedBroadcast
	// IOBroadcast write input, output and error output to broadcast
	IOBroadcast() app.BufferedBroadcast
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
	// OBroadcast write all tasks output logs to broadcast
	OBroadcast() app.BufferedBroadcast
	// StatusBroadcast write tasks statuses changes to broadcast
	StatusBroadcast() app.BufferedBroadcast
	// Summary write summary log to output
	Summary(out app.Output) (err error)
	// Names return tasks names
	Names() []string
	// Get task by name
	Get(name string) (task Task, ok bool)
	// Create new task from Pip
	Create(pip Pip) (task TaskWriter, err error)
	// Wait
	Wait() (err error)
}

// TasksUnit menage pipeline tasks
type TasksUnit interface {
	FromScope(scp app.Scope) (tasks TasksManager, err error)
	BindScope(scp app.Scope, tasks TasksManager) (err error)
	Clear(scp app.Scope) (err error)
}
