package jobsync

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/goatcms/goatcore/workers"
)

type TestErrorBody struct {
	pool      *Pool
	lifecycle *Lifecycle
	acc       *TestAccumulator
	max       int
}

func (t *TestErrorBody) Loop() {
	defer t.pool.Done()
	t.lifecycle.Error(fmt.Errorf("test error"))
	for i := 0; i < t.max; i++ {
		t.acc.Add(1)
		if t.lifecycle.IsKilled() {
			return
		}
		runtime.Gosched()
	}
}

func TestStrictErr(t *testing.T) {
	maxCounter := 1000
	expectValue := maxCounter * workers.MaxJob
	for ti := 0; ti < workers.AsyncTestReapeat; ti++ {
		accumulator := &TestAccumulator{}
		lifecycle := NewLifecycle(workers.DefaultTestTimeout, true)
		pool := NewPool(workers.MaxJob)
		jobCounter := pool.Add(workers.MaxJob + 33)
		for i := 0; i < jobCounter; i++ {
			body := &TestErrorBody{
				lifecycle: lifecycle,
				pool:      pool,
				acc:       accumulator,
				max:       maxCounter,
			}
			go body.Loop()
		}
		pool.Wait()
		if accumulator.Value() == expectValue {
			t.Errorf("must break before finish")
			return
		}
		// jobCounter + 1 since it contains one error for each job and one context cancel error
		expectedErrCounter := jobCounter + 1
		if len(lifecycle.Errors()) != expectedErrCounter {
			t.Errorf("each job should return one error and it chould contains context cancel error (it return %v errors and it expect %v)", len(lifecycle.Errors()), expectedErrCounter)
			return
		}
	}
}

func TestErr(t *testing.T) {
	maxCounter := 1000
	expectValue := maxCounter * workers.MaxJob
	for ti := 0; ti < workers.AsyncTestReapeat; ti++ {
		accumulator := &TestAccumulator{}
		lifecycle := NewLifecycle(workers.DefaultTestTimeout, false)
		pool := NewPool(workers.MaxJob)
		jobCounter := pool.Add(workers.MaxJob)
		for i := 0; i < jobCounter; i++ {
			body := &TestErrorBody{
				lifecycle: lifecycle,
				pool:      pool,
				acc:       accumulator,
				max:       maxCounter,
			}
			go body.Loop()
		}
		pool.Wait()
		if accumulator.Value() < expectValue {
			t.Errorf("must ignore error (expected value was %v and it is %v)", expectValue, accumulator.Value())
			return
		}
		if len(lifecycle.Errors()) != jobCounter {
			t.Errorf("each job should return one error (it return %v errors and it expect %v)", len(lifecycle.Errors()), jobCounter)
			return
		}
	}
}
