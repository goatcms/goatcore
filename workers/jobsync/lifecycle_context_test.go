package jobsync

import (
	"testing"
	"time"

	"github.com/goatcms/goatcore/workers"
)

func TestCancelContextWhenKill(t *testing.T) {
	for ti := 0; ti < workers.AsyncTestReapeat; ti++ {
		lifecycle := NewLifecycle(time.Minute, true)
		lifecycle.Kill()
		select {
		case _, ok := <-lifecycle.ctx.Done():
			if ok {
				t.Errorf("expected closed context after kill")
				return
			}
		default:
			// it is ok
		}
	}
}
