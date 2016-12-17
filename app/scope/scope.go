package scope

import "github.com/goatcms/goat-core/app"

// Scope is global scope interface
type Scope struct {
	app.Injector
	app.EventScope
	app.DataScope
}

// NewScope create new instance of scope
func NewScope(tagname string) app.Scope {
	ds := &DataScope{
		Data: make(map[string]interface{}),
	}
	return &Scope{
		EventScope: NewEventScope(),
		DataScope:  app.DataScope(ds),
		Injector:   ds.Injector(tagname),
	}
}
