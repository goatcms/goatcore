package jobsync

import (
	"context"
	"sync"
	"time"

	"github.com/goatcms/goatcore/varutil/goaterr"
)

// Lifecycle is execution lifecycle controller object
type Lifecycle struct {
	mutex sync.Mutex
	//killFlag   bool
	ctx        context.Context
	cancel     context.CancelFunc
	strictMode bool
	errors     []error
	// deadline   time.Time
	muStep sync.RWMutex
	step   int
}

// NewLifecycle create new Lifecycle instance
func NewLifecycle(lifetime time.Duration, strictMode bool) (lifecycle *Lifecycle) {
	deadline := time.Now().Add(lifetime)
	lifecycle = &Lifecycle{
		// killFlag:   false,
		strictMode: strictMode,
		errors:     []error{},
		// deadline:   deadline,
	}
	lifecycle.ctx, lifecycle.cancel = context.WithDeadline(context.Background(), deadline)
	return lifecycle
}

// Context return lifecycle context.Context
func (lifecycle *Lifecycle) Context() context.Context {
	return lifecycle.ctx
}

// Kill set lifecycle kill flag to true. It is signal to stop related goroutines
func (lifecycle *Lifecycle) Kill() {
	// lifecycle.mutex.Lock()
	//lifecycle.killFlag = true
	lifecycle.cancel()
	// lifecycle.mutex.Unlock()
}

// IsKilled check if lifecycle is kill or there was a timeout
func (lifecycle *Lifecycle) IsKilled() bool {
	/*if lifecycle.killFlag {
		return true
	}
	timediff := lifecycle.deadline.Sub(time.Now())
	if timediff < 0 {
		lifecycle.Error(workers.ErrTimeout)
		lifecycle.Kill()
		return true
	}
	return false*/
	select {
	case <-lifecycle.ctx.Done():
		return true
	default:
	}
	return false
}

// Error append errors to lifecycle and kill it.
func (lifecycle *Lifecycle) Error(e ...error) {
	lifecycle.mutex.Lock()
	lifecycle.errors = append(lifecycle.errors, e...)
	if lifecycle.strictMode {
		// lifecycle.killFlag = true
		lifecycle.Kill()
	}
	lifecycle.mutex.Unlock()
}

// Errors return lifecycle error array
func (lifecycle *Lifecycle) Errors() []error {
	return goaterr.AppendError(lifecycle.errors, lifecycle.ctx.Err())
}

// Step return lifecycle step
func (lifecycle *Lifecycle) Step() int {
	lifecycle.muStep.RLock()
	defer lifecycle.muStep.RUnlock()
	return lifecycle.step
}

// NextStep set new lifecycle step.
func (lifecycle *Lifecycle) NextStep(step int) {
	lifecycle.muStep.Lock()
	lifecycle.step = step
	lifecycle.muStep.Unlock()
}
