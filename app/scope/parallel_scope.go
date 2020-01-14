package scope

import (
	"context"
	"sync"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

// ParallelScope represent sub-scope
type ParallelScope struct {
	app.DataScope
	app.EventScope
	parent    app.Scope
	errors    []error
	errorsMU  sync.Mutex
	waitGroup sync.WaitGroup
	injectors []app.Injector
}

// NewParallelScope create new instance of scope
func NewParallelScope(parent app.Scope, params Params) app.Scope {
	if params.DataScope == nil {
		params.DataScope = parent
	}
	if params.EventScope == nil {
		params.EventScope = parent
	}
	return &ParallelScope{
		parent:     parent,
		DataScope:  params.DataScope,
		EventScope: params.EventScope,
		injectors:  params.Injectors,
	}
}

// Context return shared context
func (scp *ParallelScope) Context() context.Context {
	return scp.parent.Context()
}

// Kill shared context
func (scp *ParallelScope) Kill() {
	scp.parent.Kill()
}

// IsKilled return true if shared context is killed/ended
func (scp *ParallelScope) IsKilled() bool {
	return scp.parent.IsKilled()
}

// Wait for end of all tasks in child scope
func (scp *ParallelScope) Wait() (err error) {
	scp.waitGroup.Wait()
	return goaterr.ToErrors(scp.errors)
}

// AddTasks tasks to child scope
func (scp *ParallelScope) AddTasks(delta int) (err error) {
	scp.waitGroup.Add(delta)
	return nil
}

// DoneTask done one child scope task
func (scp *ParallelScope) DoneTask() {
	scp.waitGroup.Done()
}

// Errors return child scope errors
func (scp *ParallelScope) Errors() []error {
	return scp.errors
}

// ToError return error if child scope contains any error
func (scp *ParallelScope) ToError() error {
	return goaterr.ToErrors(scp.errors)
}

// AppendError add error to child and parent scope
func (scp *ParallelScope) AppendError(err error) {
	scp.errorsMU.Lock()
	defer scp.errorsMU.Unlock()
	scp.Kill()
	scp.parent.AppendError(err)
	scp.errors = append(scp.errors, err)
}

// AppendErrors add errors to child and parent scope
func (scp *ParallelScope) AppendErrors(errs ...error) {
	scp.errorsMU.Lock()
	defer scp.errorsMU.Unlock()
	scp.Kill()
	scp.parent.AppendErrors(errs...)
	scp.errors = append(scp.errors, errs...)
}

// InjectTo insert data to object
func (scp *ParallelScope) InjectTo(obj interface{}) (err error) {
	for _, scpInjector := range scp.injectors {
		if err = scpInjector.InjectTo(obj); err != nil {
			return err
		}
	}
	return scp.parent.InjectTo(obj)
}

// Close child scope
func (scp *ParallelScope) Close() (err error) {
	return scp.Wait()
}
