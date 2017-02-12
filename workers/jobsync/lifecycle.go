package jobsync

import (
	"sync"
	"time"

	"github.com/goatcms/goat-core/workers"
)

type Lifecycle struct {
	mutex      sync.Mutex
	killFlag   bool
	strictMode bool
	errors     []error
	deadpoint  time.Time
}

func NewLifecycle(lifetime time.Duration, strictMode bool) *Lifecycle {
	deadpoint := time.Now().Add(lifetime)
	return &Lifecycle{
		killFlag:   false,
		strictMode: strictMode,
		errors:     []error{},
		deadpoint:  deadpoint,
	}
}

func (lifecycle *Lifecycle) Kill() {
	lifecycle.mutex.Lock()
	lifecycle.killFlag = true
	lifecycle.mutex.Unlock()
}

func (lifecycle *Lifecycle) IsKilled() bool {
	if lifecycle.killFlag {
		return true
	}
	timediff := lifecycle.deadpoint.Sub(time.Now())
	if timediff < 0 {
		lifecycle.Error(workers.TimeoutError)
		lifecycle.Kill()
		return true
	}
	return false
}

func (lifecycle *Lifecycle) Error(e ...error) {
	lifecycle.mutex.Lock()
	lifecycle.errors = append(lifecycle.errors, e...)
	if lifecycle.strictMode {
		lifecycle.killFlag = true
	}
	lifecycle.mutex.Unlock()
}

func (lifecycle *Lifecycle) Errors() []error {
	return lifecycle.errors
}
