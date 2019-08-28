package scope

import (
	"context"
	"sync"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/varutil/goaterr"
	"github.com/goatcms/goatcore/workers/jobsync"
)

// SyncScope is default error scope
type SyncScope struct {
	lifecycle *jobsync.Lifecycle
	waitGroup *sync.WaitGroup
}

// NewSyncScope create new instance of error scope
func NewSyncScope() app.SyncScope {
	return &SyncScope{
		lifecycle: jobsync.NewLifecycle(app.DefaultDeadline, true),
		waitGroup: &sync.WaitGroup{},
	}
}

// Context return golang context
func (s *SyncScope) Context() context.Context {
	return s.lifecycle.Context()
}

// Kill scope
func (s *SyncScope) Kill() {
	s.lifecycle.Kill()
}

// IsKilled check if scope is killed
func (s *SyncScope) IsKilled() bool {
	return s.lifecycle.IsKilled()
}

// ToError return scope error object or nil if does't contains a error
func (s *SyncScope) ToError() error {
	return goaterr.ToErrors(s.Errors())
}

// Errors return scope errors
func (s *SyncScope) Errors() []error {
	return s.lifecycle.Errors()
}

// AppendError append error to scope (skip nil error)
func (s *SyncScope) AppendError(err error) {
	if err == nil {
		return
	}
	s.lifecycle.Error(err)
}

// AppendErrors append many errors to scope (skip nil errors)
func (s *SyncScope) AppendErrors(errs ...error) {
	for _, err := range errs {
		s.AppendError(err)
	}
}

// Wait for finish of context execution
func (s *SyncScope) Wait() error {
	s.waitGroup.Wait()
	return s.ToError()
}

// AddTask add task to execute
func (s *SyncScope) AddTask(delta int) {
	s.waitGroup.Add(delta)
}

// AddTasks add task to wait group
func (s *SyncScope) AddTasks(delta int) {
	s.waitGroup.Add(delta)
}

// DoneTask mark single task as done
func (s *SyncScope) DoneTask() {
	s.waitGroup.Done()
}
