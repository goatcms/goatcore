package jobsync

import (
	"testing"

	"github.com/goatcms/goat-core/workers"
)

type TestInfinityBody struct {
	pool      *Pool
	lifecycle *Lifecycle
}

func (t *TestInfinityBody) Loop() {
	defer t.pool.Done()
	for {
		// do something
		if t.lifecycle.IsKilled() {
			return
		}
	}
}

func TestKill(t *testing.T) {
	for ti := 0; ti < workers.AsyncTestReapeat; ti++ {
		lifecycle := NewLifecycle(workers.DefaultTestTimeout, true)
		pool := NewPool(workers.MaxJob)
		jobCounter := pool.Add(workers.MaxJob)
		for i := 0; i < jobCounter; i++ {
			body := &TestInfinityBody{
				lifecycle: lifecycle,
				pool:      pool,
			}
			go body.Loop()
		}
		lifecycle.Kill()
		pool.Wait()
	}
}
