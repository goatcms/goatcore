package workers

import (
	"fmt"
	"runtime"
	"time"
)

const (
	// AsyncTestReapeat specify the number of async test loop iteration
	AsyncTestReapeat = 1000
	// DefaultTestTimeout specify test timeout
	DefaultTestTimeout = time.Minute
	// DefaultTimeout specify default timeout
	DefaultTimeout = 2 * time.Minute
)

var (
	// MaxJob is default max goroutines for task
	MaxJob = runtime.NumCPU()
	// ErrTimeout is timeout error
	ErrTimeout = fmt.Errorf("JobGroup timeout")
	// ErrKill is job/goroutines kill error
	ErrKill = fmt.Errorf("JobGroup was killed")
)
