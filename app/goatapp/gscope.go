package goatapp

import (
	"github.com/goatcms/goat-core/app"
	"github.com/goatcms/goat-core/app/injector"
	"github.com/goatcms/goat-core/app/scope"
)

// GlobalScope represent dependency scope lvl
type GlobalScope struct {
	app.EventScope
	app.DataScope
	app.Injector

	scopes []app.Scope
}

// NewGlobalScope create new dependency scope
func NewGlobalScope(tagname string, scopes []app.Scope) app.Scope {
	dataScope := &scope.DataScope{
		Data: make(map[string]interface{}),
	}
	injectors := make([]app.Injector, len(scopes)+1)
	for i, scope := range scopes {
		injectors[i] = scope
	}
	injectors[len(injectors)-1] = dataScope.Injector(tagname)
	globalScope := GlobalScope{
		EventScope: scope.NewEventScope(),
		DataScope:  dataScope,
		Injector:   injector.NewMultiInjector(injectors),
		scopes:     scopes,
	}
	return globalScope
}

func (gs GlobalScope) Trigger(key int, data interface{}) error {
	if err := gs.EventScope.Trigger(key, data); err != nil {
		return err
	}
	for _, scope := range gs.scopes {
		if err := scope.Trigger(key, data); err != nil {
			return err
		}
	}
	return nil
}
