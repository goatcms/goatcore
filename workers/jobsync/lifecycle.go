package jobsync

import (
	"sync"
	"time"

	"github.com/goatcms/goatcore/workers"
)

// Lifecycle is execution lifecycle controller object
type Lifecycle struct {
	mutex      sync.Mutex
	killFlag   bool
	strictMode bool
	errors     []error
	deadpoint  time.Time
	muStep     sync.RWMutex
	step       int
}

// NewLifecycle create new Lifecycle instance
func NewLifecycle(lifetime time.Duration, strictMode bool) *Lifecycle {
	deadpoint := time.Now().Add(lifetime)
	return &Lifecycle{
		killFlag:   false,
		strictMode: strictMode,
		errors:     []error{},
		deadpoint:  deadpoint,
	}
}

// Kill set lifecycle kill flag to true. It is signal to stop related goroutines
func (lifecycle *Lifecycle) Kill() {
	lifecycle.mutex.Lock()
	lifecycle.killFlag = true
	lifecycle.mutex.Unlock()
}

// IsKilled check if lifecycle is kill or there was a timeout
func (lifecycle *Lifecycle) IsKilled() bool {
	if lifecycle.killFlag {
		return true
	}
	timediff := lifecycle.deadpoint.Sub(time.Now())
	if timediff < 0 {
		lifecycle.Error(workers.ErrTimeout)
		lifecycle.Kill()
		return true
	}
	return false
}

// Error append errors to lifecycle and kill it.
func (lifecycle *Lifecycle) Error(e ...error) {
	lifecycle.mutex.Lock()
	lifecycle.errors = append(lifecycle.errors, e...)
	if lifecycle.strictMode {
		lifecycle.killFlag = true
	}
	lifecycle.mutex.Unlock()
}

// Errors return lifecycle error array
func (lifecycle *Lifecycle) Errors() []error {
	return lifecycle.errors
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
