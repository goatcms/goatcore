package contextscope

import (
	"context"
	"sync"
	"time"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

// The isolated scope is a isolated context scope depended from parent scope. The isolated scope doesn't affect parent.
type Isolated struct {
	done     chan struct{}
	errors   []error
	errorsMU sync.Mutex
	parent   app.ContextScope
}

// NewIsolated create new isolated context scope instance
func NewIsolated(parent app.ContextScope) app.ContextScope {
	isolated := &Isolated{
		parent: parent,
		done:   make(chan struct{}),
	}
	go func() {
		select {
		case <-parent.Done():
			// the gorutine kill isolated context if parent die.
			if len(parent.Errors()) != 0 {
				isolated.Kill()
				return
			}
			isolated.Stop()
		case <-isolated.done:
			// stop gorutine if isolated context die (prevent memory leaks)
			return
		}
	}()
	return isolated
}

// Deadline returns the time when work done on behalf of this context
// should be canceled. Deadline returns ok==false when no deadline is
// set. Successive calls to Deadline return the same results.
func (scp *Isolated) Deadline() (deadline time.Time, ok bool) {
	return deadline, false
}

// Done is close when the scope context is done (kill or stop)
func (scp *Isolated) Done() <-chan struct{} {
	return scp.done
}

// IsDone check if the scope context is done (kill or stop)
func (scp *Isolated) IsDone() bool {
	select {
	case <-scp.done:
		return true
	default:
	}
	return false
}

// Kill scope
func (scp *Isolated) Kill() {
	scp.AppendError(context.Canceled)
}

// Stop stop the scope context without error
func (scp *Isolated) Stop() {
	if !scp.IsDone() {
		close(scp.done)
	}
}

// Err return cumulative error if the scope context contains any error
func (scp *Isolated) Err() error {
	return goaterr.ToError(scp.errors)
}

// Errors return scope errors
func (scp *Isolated) Errors() []error {
	return scp.errors
}

// AppendErrors append many errors to scope (skip nil errors)
func (scp *Isolated) AppendError(errs ...error) {
	var i = 0
	if len(errs) == 0 {
		return
	}
	scp.errorsMU.Lock()
	for _, err := range errs {
		if err == nil {
			continue
		}
		scp.errors = append(scp.errors, err)
		i++
	}
	scp.errorsMU.Unlock()
	if i != 0 {
		scp.Stop()
	}
}
