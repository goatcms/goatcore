package workers

import "github.com/goatcms/goat-core/varutil/goaterr"

// DeferFunc is a function run on stop runner (kill or finish)
type DeferFunc func() error

// JobBody contains code executor
type JobBody interface {
	Step() (bool, error)
}

// Job is async code api
type Job interface {
	Run() error
	Kill()
	KillSlot(interface{}) error
	Wait() goaterr.Errors
	Defer(DeferFunc)
}
