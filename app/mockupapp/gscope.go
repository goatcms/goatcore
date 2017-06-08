package mockupapp

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/scope"
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
	globalScope := GlobalScope{
		EventScope: scope.NewEventScope(),
		DataScope:  dataScope,
		Injector:   dataScope.Injector(tagname),
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
