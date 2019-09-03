package workers

import (
	"runtime"
	"time"

	"github.com/goatcms/goatcore/varutil/goaterr"
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
	ErrTimeout = goaterr.Errorf("JobGroup timeout")
	// ErrKill is job/goroutines kill error
	ErrKill = goaterr.Errorf("JobGroup was killed")
)
