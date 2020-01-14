package tasks

import (
	"bytes"
	"sync"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/modules/commonm/commservices"
	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices"
)

// Task is single task object
type Task struct {
	logs   bytes.Buffer
	ctx    app.IOContext
	done   bool
	status string
	pip    pipservices.Pip
	wg     sync.WaitGroup
}

// NewTask create a Taks instance
func NewTask(ctx app.IOContext, pip pipservices.Pip) *Task {
	task := &Task{
		ctx: ctx,
		pip: pip,
	}
	task.wg.Add(1)
	return task
}

// newPipTaskWriter create a PipTaskWriter instance
func newPipTaskWriter(ctx app.IOContext, pip pipservices.Pip) pipservices.TaskWriter {
	return NewTask(ctx, pip)
}

// Name return task name
func (task *Task) Name() string {
	return task.pip.Name
}

// Logs return task result
func (task *Task) Logs() string {
	return task.logs.String()
}

// Done return true if task is finished
func (task *Task) Done() bool {
	return task.done
}

// IOContext return task IO context
func (task *Task) IOContext() (out app.IOContext) {
	return task.ctx
}

// Close mark task as done and close input data
func (task *Task) Close() (err error) {
	task.wg.Done()
	return task.ctx.Scope().Close()
}

// Wait for task finish
func (task *Task) Wait() error {
	task.wg.Wait()
	return nil
}

// WaitList return list of related tasks to wait for
func (task *Task) WaitList() []string {
	return task.pip.Wait
}

// LockMap return map described related resources to lock
func (task *Task) LockMap() commservices.LockMap {
	return task.pip.Lock
}

// Status return taks status
func (task *Task) Status() string {
	return task.status
}

// SetStatus return set taks status
func (task *Task) SetStatus(status string) {
	task.status = status
}

// Errors return task errors (or nil)
func (task *Task) Errors() []error {
	return task.ctx.Scope().Errors()
}
