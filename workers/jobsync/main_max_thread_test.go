package jobsync

import (
	"sync"
	"testing"

	"github.com/goatcms/goatcore/workers"
)

type TestThreadCounter struct {
	pool      *Pool
	lifecycle *Lifecycle
	muCount   sync.Mutex
	count     int
}

func (t *TestThreadCounter) Loop() {
	defer t.pool.Done()
	t.muCount.Lock()
	defer t.muCount.Unlock()
	t.count++
}

func TestMaxThread3(t *testing.T) {
	expectedCounter := 3
	for ti := 0; ti < workers.AsyncTestReapeat; ti++ {
		lifecycle := NewLifecycle(workers.DefaultTestTimeout, true)
		pool := NewPool(expectedCounter)
		jobCounter := pool.Add(40)
		body := &TestThreadCounter{
			lifecycle: lifecycle,
			pool:      pool,
			count:     0,
		}
		for i := 0; i < jobCounter; i++ {
			go body.Loop()
		}
		pool.Wait()
		if body.count != expectedCounter {
			t.Errorf("expect %v threads and it is %v threads", expectedCounter, body.count)
			return
		}
		if len(lifecycle.Errors()) != 0 {
			t.Errorf("%v", lifecycle.Errors())
			return
		}
	}
}
