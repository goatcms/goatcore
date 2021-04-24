package scope

import (
	"context"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/injector"
)

// ChildParams describe child scope
type ChildParams struct {
	DataScope  app.DataScope
	EventScope app.EventScope
	Injector   app.Injector
}

// Params describe scope
type Params struct {
	DataScope    app.DataScope
	EventScope   app.EventScope
	Injector     app.Injector
	ContextScope app.ContextScope
	Tag          string
}

// Scope is global scope interface
type Scope struct {
	app.Injector
	app.EventScope
	app.DataScope
	app.ContextScope
	parent app.Scope
}

// NewScope create new instance of scope
func NewScope(params Params) app.Scope {
	if params.DataScope == nil {
		params.DataScope = NewDataScope(make(map[interface{}]interface{}))
	}
	if params.Tag != "" {
		if params.Injector != nil {
			params.Injector = injector.NewMultiInjector([]app.Injector{
				params.Injector,
				NewScopeInjector(params.Tag, params.DataScope),
			})
		} else {
			params.Injector = NewScopeInjector(params.Tag, params.DataScope)
		}
	}
	if params.Injector == nil {
		params.Injector = injector.NewNilInjector()
	}
	if params.EventScope == nil {
		params.EventScope = NewEventScope()
	}
	if params.ContextScope == nil {
		params.ContextScope = NewContextScope()
	}
	return &Scope{
		EventScope:   params.EventScope,
		DataScope:    params.DataScope,
		Injector:     params.Injector,
		ContextScope: params.ContextScope,
	}
}

// Close scope
func (scp *Scope) Close() (err error) {
	scp.AppendError(scp.Trigger(app.BeforeCloseEvent, scp))
	if err = scp.Wait(); err != nil {
		scp.AppendError(scp.Trigger(app.BeforeRollbackEvent, scp))
		scp.AppendError(scp.Trigger(app.RollbackEvent, scp))
		scp.AppendError(scp.Trigger(app.AfterRollbackEvent, scp))
		scp.close()
		return scp.Err()
	}
	scp.AppendError(scp.Trigger(app.BeforeCommitEvent, scp))
	scp.AppendError(scp.Trigger(app.CommitEvent, scp))
	scp.AppendError(scp.Trigger(app.AfterCommitEvent, scp))
	scp.close()
	return scp.Err()
}

func (scp *Scope) close() {
	if scp.parent != nil {
		scp.parent.DoneTask()
	}
	scp.AppendError(scp.Trigger(app.AfterCloseEvent, scp))
	scp.EventScope = nil
	scp.DataScope = nil
	scp.Injector = nil
}

// Kill scope
func (scp *Scope) Kill() {
	scp.ContextScope.Kill()
	scp.AppendError(scp.Trigger(app.KillEvent, nil))
}

// Stop scope
func (scp *Scope) Stop() {
	scp.ContextScope.Stop()
	scp.AppendError(scp.Trigger(app.StopEvent, nil))
}

// AppendError add an error to the scope
func (scp *Scope) AppendError(err ...error) {
	scp.ContextScope.AppendError(err...)
	scp.ContextScope.AppendError(scp.Trigger(app.ErrorEvent, err))
}

// Kill scope
func (scp *Scope) GoContext() context.Context {
	return scp
}
