package jobsync

import (
	"sync"
	"testing"
	"time"

	"github.com/goatcms/goat-core/workers"
)

type TestEnded struct {
	endedMu sync.Mutex
	ended   bool
}

func (e *TestEnded) End() {
	e.endedMu.Lock()
	e.ended = true
	e.endedMu.Unlock()
}

func (e *TestEnded) IsEnded() bool {
	e.endedMu.Lock()
	defer e.endedMu.Unlock()
	return e.ended
}

type TestTimeoutBody struct {
	pool      *Pool
	lifecycle *Lifecycle
	ended     *TestEnded
}

func (t *TestTimeoutBody) Loop() {
	defer t.pool.Done()
	finishPoint := time.Now().Add(time.Minute)
	for {
		if time.Now().Sub(finishPoint) > 0 {
			t.ended.End()
			return
		}
		if t.lifecycle.IsKilled() {
			return
		}
	}
}

func TestTimeout(t *testing.T) {
	for ti := 0; ti < workers.AsyncTestReapeat; ti++ {
		ended := &TestEnded{}
		lifecycle := NewLifecycle(time.Microsecond, true)
		pool := NewPool(workers.MaxJob)
		jobCounter := pool.Add(workers.MaxJob + 33)
		for i := 0; i < jobCounter; i++ {
			body := &TestTimeoutBody{
				lifecycle: lifecycle,
				pool:      pool,
				ended:     ended,
			}
			go body.Loop()
		}
		pool.Wait()
		if ended.IsEnded() != false {
			t.Errorf("must break before finished")
			return
		}
		if len(lifecycle.Errors()) == 0 {
			t.Errorf("must contains timeout error")
			return
		}
	}
}
