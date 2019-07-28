package jobsync

import (
	"testing"

	"github.com/goatcms/goatcore/workers"
)

type TestCounterBody struct {
	pool      *Pool
	lifecycle *Lifecycle
	acc       *TestAccumulator
	max       int
}

func (tc *TestCounterBody) Loop() {
	defer tc.pool.Done()
	for i := 0; i < tc.max; i++ {
		tc.acc.Add(1)
		if tc.lifecycle.IsKilled() {
			return
		}
	}
}

func TestParall(t *testing.T) {
	maxCounter := 1000
	expectValue := maxCounter * workers.MaxJob
	for ti := 0; ti < workers.AsyncTestReapeat; ti++ {
		accumulator := &TestAccumulator{}
		lifecycle := NewLifecycle(workers.DefaultTestTimeout, true)
		pool := NewPool(workers.MaxJob)
		jobCounter := pool.Add(workers.MaxJob)
		for i := 0; i < jobCounter; i++ {
			body := &TestCounterBody{
				lifecycle: lifecycle,
				pool:      pool,
				acc:       accumulator,
				max:       maxCounter,
			}
			go body.Loop()
		}
		pool.Wait()
		if accumulator.Value() != expectValue {
			t.Errorf("body.Counter must be equals to maxValue %v (and it is %v)", expectValue, accumulator.Value())
			return
		}
	}
}
