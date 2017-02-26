package workers

import (
	"fmt"
	"runtime"
	"time"
)

const (
	AsyncTestReapeat   = 1000
	DefaultTestTimeout = time.Minute
	DefaultTimeout     = 2 * time.Minute
)

var (
	MaxJob = runtime.NumCPU()

	TimeoutError = fmt.Errorf("JobGroup timeout")
	KilledError  = fmt.Errorf("JobGroup was killed")
)
