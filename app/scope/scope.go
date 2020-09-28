package scope

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/injector"
)

// ChildParams describe child scope
type ChildParams struct {
	DataScope  app.DataScope
	EventScope app.EventScope
	Injectors  []app.Injector
}

// Params describe scope
type Params struct {
	DataScope  app.DataScope
	EventScope app.EventScope
	Injectors  []app.Injector
	SyncScope  app.SyncScope
	Tag        string
}

// Scope is global scope interface
type Scope struct {
	app.Injector
	app.EventScope
	app.DataScope
	app.SyncScope
}

// NewScope create new instance of scope
func NewScope(params Params) app.Scope {
	var scopeInjector app.Injector
	if params.DataScope == nil {
		params.DataScope = NewDataScope(make(map[string]interface{}))
	}
	if params.Tag != "" {
		params.Injectors = append(params.Injectors, NewScopeInjector(params.Tag, params.DataScope))
	}
	if len(params.Injectors) == 1 {
		scopeInjector = params.Injectors[0]
	} else if len(params.Injectors) > 1 {
		scopeInjector = injector.NewMultiInjector(params.Injectors)
	} else {
		scopeInjector = injector.NewNilInjector()
	}
	if params.EventScope == nil {
		params.EventScope = NewEventScope()
	}
	if params.SyncScope == nil {
		params.SyncScope = NewSyncScope(nil)
	}
	return &Scope{
		EventScope: params.EventScope,
		DataScope:  params.DataScope,
		Injector:   scopeInjector,
		SyncScope:  params.SyncScope,
	}
}

// Close scope
func (scp *Scope) Close() (err error) {
	if err = scp.Wait(); err != nil {
		scp.Kill()
		scp.AppendError(scp.Trigger(app.RollbackEvent, scp))
		scp.destroy()
		return scp.ToError()
	}
	scp.Kill()
	scp.AppendError(scp.Trigger(app.CommitEvent, scp))
	scp.destroy()
	return scp.ToError()
}

// Kill scope
func (scp *Scope) Kill() {
	scp.SyncScope.Kill()
	if err := scp.Trigger(app.KillEvent, nil); err != nil {
		scp.AppendError(err)
	}
}

func (scp *Scope) destroy() {
	scp.EventScope = nil
	scp.DataScope = nil
	scp.Injector = nil
}
