package jobsync

import "sync"

type TestAccumulator struct {
	mu    sync.Mutex
	value int
}

func (ta *TestAccumulator) Add(v int) {
	ta.mu.Lock()
	defer ta.mu.Unlock()
	ta.value += v
}

func (ta *TestAccumulator) Value() int {
	return ta.value
}
