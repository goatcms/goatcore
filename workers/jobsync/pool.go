package jobsync

import "sync"

// Pool is goroutines group quantity controller
type Pool struct {
	mutex   sync.Mutex
	wg      sync.WaitGroup
	counter int
	max     int
}

// NewPool create new instance of Poll
func NewPool(max int) *Pool {
	return &Pool{
		counter: 0,
		max:     max,
	}
}

// Add reserve a amount of goroutines. If amount is greater than max, reserve and return max available amount.
func (p *Pool) Add(amount int) int {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	max := p.max - p.counter
	if max < amount {
		amount = max
	}
	p.wg.Add(amount)
	p.counter += amount
	return amount
}

// Done is signal that mean a goroutine is finish.
func (p *Pool) Done() {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.wg.Done()
	p.counter--
}

// Wait is function stop current goroutine to finish all goroutine in the Pool.
func (p *Pool) Wait() {
	p.wg.Wait()
}
