package scope

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

// ErrDoned is the error returned by ContextScope.AddTasks when the context scope is done.
var ErrDoned = errors.New("context scope is done")

// ContextScope is default error scope
type ContextScope struct {
	done      chan struct{}
	errorsMU  sync.Mutex
	errors    []error
	waitGroup *sync.WaitGroup
}

// NewContextScope create new instance of error scope
func NewContextScope() app.ContextScope {
	return &ContextScope{
		done:      make(chan struct{}),
		waitGroup: &sync.WaitGroup{},
	}
}

// Deadline returns the time when work done on behalf of this context
// should be canceled. Deadline returns ok==false when no deadline is
// set. Successive calls to Deadline return the same results.
func (s *ContextScope) Deadline() (deadline time.Time, ok bool) {
	return deadline, false
}

// Done is close when the scope context is done (kill or stop)
func (s *ContextScope) Done() <-chan struct{} {
	return s.done
}

// IsDone check if the scope context is done (kill or stop)
func (s *ContextScope) IsDone() bool {
	select {
	case <-s.done:
		return true
	default:
	}
	return false
}

// Kill scope
func (s *ContextScope) Kill() {
	s.AppendError(context.Canceled)
}

// Stop stop the scope context without error
func (s *ContextScope) Stop() {
	if !s.IsDone() {
		close(s.done)
	}
}

// Err return cumulative error if the scope context contains any error
func (s *ContextScope) Err() error {
	return goaterr.ToError(s.errors)
}

// Errors return scope errors
func (s *ContextScope) Errors() []error {
	return s.errors
}

// AppendErrors append many errors to scope (skip nil errors)
func (s *ContextScope) AppendError(errs ...error) {
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
	if i != 0 {
		s.Stop()
	}
}

// Wait for finish of context execution
func (s *ContextScope) Wait() error {
	s.waitGroup.Wait()
	return s.Err()
}

// AddTasks add task to wait group
func (s *ContextScope) AddTasks(delta int) (err error) {
	if s.IsDone() {
		return ErrDoned
	}
	s.waitGroup.Add(delta)
	return nil
}

// DoneTask mark single task as done
func (s *ContextScope) DoneTask() {
	s.waitGroup.Done()
}
