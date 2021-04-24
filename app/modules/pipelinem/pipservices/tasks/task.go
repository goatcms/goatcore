package tasks

import (
	"sync"

	"github.com/goatcms/goatcore/app/gio"
	"github.com/goatcms/goatcore/app/gio/bufferio"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/modules/commonm/commservices"
	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

// Task is single task object
type Task struct {
	ctx             app.IOContext
	done            bool
	status          string
	pip             pipservices.Pip
	fullName        string
	wg              sync.WaitGroup
	oBroadcast      app.BufferedBroadcast
	ioBroadcast     app.BufferedBroadcast
	statusBroadcast app.Broadcast
	closeCB         func()
}

// NewTask create a Taks instance
func NewTask(ctx app.IOContext, pip pipservices.Pip, statusBroadcast app.Broadcast, closeCB func()) *Task {
	if ctx == nil {
		panic(goaterr.Errorf("context is required"))
	}
	oBroadcast := bufferio.NewBroadcast(nil, nil)
	ioBroadcast := bufferio.NewBroadcast(nil, nil)
	ctxIO := ctx.IO()
	ioBroadcastIO := gio.NewRepeatIO(gio.IOParams{
		In:  ctxIO.In(),
		Out: ioBroadcast,
		Err: ioBroadcast,
		CWD: ctxIO.CWD(),
	})
	childIO := gio.NewIO(gio.IOParams{
		In:  ioBroadcastIO.In(),
		Out: gio.NewMultiOutput([]app.Output{oBroadcast, ioBroadcastIO.Out(), ctxIO.Out()}),
		Err: gio.NewMultiOutput([]app.Output{oBroadcast, ioBroadcastIO.Err(), ctxIO.Err()}),
		CWD: ctxIO.CWD(),
	})
	task := &Task{
		ctx:             gio.NewIOContext(ctx.Scope(), childIO),
		pip:             pip,
		closeCB:         closeCB,
		oBroadcast:      oBroadcast,
		ioBroadcast:     ioBroadcast,
		statusBroadcast: statusBroadcast,
	}
	ns := pip.Namespaces.Task()
	if ns != "" {
		task.fullName = ns + ":" + task.pip.Name
	} else {
		task.fullName = pip.Name
	}
	task.wg.Add(1)
	if task.pip.Description != "" {
		statusBroadcast.Printf("\n [%s] %s... started", task.FullName(), task.Description())
	} else {
		statusBroadcast.Printf("\n [%s]... started", task.FullName())
	}
	return task
}

// Name return task name
func (task *Task) Name() string {
	return task.pip.Name
}

// FullName return task name
func (task *Task) FullName() string {
	return task.fullName
}

// Description return task description
func (task *Task) Description() string {
	return task.pip.Description
}

// OBroadcast return task output broadcast
func (task *Task) OBroadcast() app.BufferedBroadcast {
	return task.oBroadcast
}

// IOBroadcast return task input and output broadcast
func (task *Task) IOBroadcast() app.BufferedBroadcast {
	return task.ioBroadcast
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
	if task.closeCB != nil {
		task.closeCB()
	}
	if task.pip.Description != "" {
		task.statusBroadcast.Printf("\n [%s] %s... %s", task.FullName(), task.pip.Description, task.status)
	} else {
		task.statusBroadcast.Printf("\n [%s]... %s", task.FullName(), task.status)
	}
	return task.ctx.Scope().Close()
}

// Wait for task finish
func (task *Task) Wait() error {
	task.wg.Wait()
	return task.ctx.Scope().Err()
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
