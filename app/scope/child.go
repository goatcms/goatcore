package scope

import (
	"github.com/goatcms/goatcore/app"
)

// NewChildScope create new instance of scope
func NewChildScope(parent app.Scope, params ChildParams) app.Scope {
	parent.AddTasks(1)
	if params.DataScope == nil {
		params.DataScope = NewChildDataScope(parent, make(map[interface{}]interface{}))
	}
	if params.EventScope == nil {
		params.EventScope = NewChildEventScope(parent)
	}
	if params.Injector == nil {
		params.Injector = parent
	}
	return &Scope{
		parent:       parent,
		ContextScope: NewChildContextScope(parent),
		DataScope:    params.DataScope,
		EventScope:   params.EventScope,
		Injector:     params.Injector,
	}
}
