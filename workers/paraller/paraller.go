package paraller

import (
	"runtime"
	"sync"

	"github.com/goatcms/goat-core/varutil/goaterr"
	"github.com/goatcms/goat-core/workers"
)

// Paraller provide paraller job
type Paraller struct {
	body     workers.JobBody
	defers   []workers.DeferFunc
	killed   bool
	errorsMu sync.Mutex
	errors   []error
	deferMu  sync.Mutex
	wg       sync.WaitGroup
}

// NewParaller create new paraller instance
func NewParaller(body workers.JobBody) *Paraller {
	return &Paraller{
		body:   body,
		errors: []error{},
	}
}

// Kill set job to kill asap
func (r *Paraller) Kill() {
	r.killed = true
}

// KillSlot is a kill event slot
func (r *Paraller) KillSlot(interface{}) error {
	r.killed = true
	return nil
}

// Run thread
func (r *Paraller) Run() error {
	numCPU := runtime.NumCPU()
	r.wg.Add(numCPU)
	for i := 0; i < numCPU; i++ {
		go r.run()
	}
	return nil
}

// Run thread
func (r *Paraller) run() {
	defer r.wg.Done()
	for {
		hasNext, err := r.body.Step()
		if err != nil {
			r.addError(err)
			r.killed = true
		}
		if r.killed || (hasNext == false) {
			return
		}
	}
}

// Wait frozen current thread to end of app execution
func (r *Paraller) Wait() goaterr.Errors {
	r.wg.Wait()
	if len(r.errors) > 0 {
		return goaterr.NewErrors(r.errors)
	}
	return nil
}

// Defer connect function with finish/or error action
func (r *Paraller) Defer(f workers.DeferFunc) {
	r.deferMu.Lock()
	r.defers = append(r.defers, f)
	if len(r.defers) == 1 {
		go r.deferInvoker()
	}
	r.deferMu.Unlock()
}

func (r *Paraller) deferInvoker() {
	r.wg.Wait()
	for _, def := range r.defers {
		if err := def(); err != nil {
			r.errors = append(r.errors, err)
		}
	}
}

func (r *Paraller) addError(err error) {
	r.errorsMu.Lock()
	r.errors = append(r.errors, err)
	r.errorsMu.Unlock()
}
