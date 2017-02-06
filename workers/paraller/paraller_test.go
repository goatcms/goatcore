package paraller

import (
	"runtime"
	"sync"
	"testing"
)

type TestCounterBody struct {
	mu      sync.Mutex
	Counter int
}

func (t *TestCounterBody) Step() (bool, error) {
	t.mu.Lock()
	t.Counter++
	t.mu.Unlock()
	return false, nil
}

func TestCountingThread(t *testing.T) {
	body := &TestCounterBody{
		Counter: 0,
	}
	job := NewParaller(body)
	if err := job.Run(); err != nil {
		t.Error(err)
		return
	}
	if err := job.Wait(); err != nil {
		t.Error(err)
		return
	}
	numCPU := runtime.NumCPU()
	// test result
	if body.Counter != numCPU {
		t.Errorf("body.Counter must be equals to NumCPU %v (and it is %v)", numCPU, body.Counter)
	}
}

func TestDefer(t *testing.T) {
	var wg sync.WaitGroup
	var deferExec bool
	wg.Add(1)
	body := &TestCounterBody{
		Counter: 0,
	}
	job := NewParaller(body)
	if err := job.Run(); err != nil {
		t.Error(err)
		return
	}
	job.Defer(func() error {
		deferExec = true
		wg.Done()
		return nil
	})
	wg.Wait()
	if !deferExec {
		t.Errorf("defer function must be executed")
	}
}
