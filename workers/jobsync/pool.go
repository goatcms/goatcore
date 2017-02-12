package jobsync

import "sync"

type Pool struct {
	mutex   sync.Mutex
	wg      sync.WaitGroup
	counter int
	max     int
}

func NewPool(max int) *Pool {
	return &Pool{
		counter: 0,
		max:     max,
	}
}

func (p *Pool) Add(v int) int {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	max := p.max - p.counter
	if max < v {
		v = max
	}
	p.wg.Add(v)
	p.counter += v
	return v
}

func (p *Pool) Done() {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.wg.Done()
	p.counter--
}

func (p *Pool) Wait() {
	p.wg.Wait()
}
