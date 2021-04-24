package scope

import (
	"context"
	"sync"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

// ChildContextScope is default error scope
type ChildContextScope struct {
	app.ContextScope
	waitGroup sync.WaitGroup
	errorsMU  sync.Mutex
	errors    []error
}

// NewChildContextScope create new instance of error scope
func NewChildContextScope(parent app.ContextScope) app.ContextScope {
	return &ChildContextScope{
		ContextScope: parent,
	}
}

// Kill scope
func (s *ChildContextScope) Kill() {
	s.AppendError(context.Canceled)
}

// Err return cumulative error if the scope context contains any error
func (s *ChildContextScope) Err() error {
	return goaterr.ToError(s.errors)
}

// Errors return scope errors
func (s *ChildContextScope) Errors() []error {
	return s.errors
}

// AppendErrors append many errors to scope (skip nil errors)
func (s *ChildContextScope) AppendError(errs ...error) {
	var i = 0
	if len(errs) == 0 {
		return
	}
	s.errorsMU.Lock()
	for _, err := range errs {
		if err == nil {
			continue
		}
		s.errors = append(s.errors, err)
		i++
	}
	s.errorsMU.Unlock()
	if s.ContextScope != nil {
		s.ContextScope.AppendError(errs...)
	}
}

// Wait for finish of context execution
func (s *ChildContextScope) Wait() error {
	s.waitGroup.Wait()
	return s.Err()
}

// AddTasks add task to wait group
func (s *ChildContextScope) AddTasks(delta int) (err error) {
	if s.IsDone() {
		return ErrDoned
	}
	s.waitGroup.Add(delta)
	return nil
}

// DoneTask mark single task as done
func (s *ChildContextScope) DoneTask() {
	s.waitGroup.Done()
}
