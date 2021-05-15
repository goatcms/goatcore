package scope

import (
	"fmt"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/scope/datascope"
	"github.com/goatcms/goatcore/app/scope/eventscope"
	"github.com/goatcms/goatcore/varutil/idutil"
)

// ChildParams describe child scope
type ChildParams struct {
	ContextScope app.ContextScope
	DataScope    app.DataScope
	EventScope   app.EventScope
	Injector     app.Injector
	Name         string
	CID          string
}

// NewChild create new instance of child scope
func NewChild(parent app.Scope, params ChildParams) app.Scope {
	var sid string
	parent.AddTasks(1)
	if params.ContextScope == nil {
		params.ContextScope = parent.BaseContextScope()
	}
	if params.DataScope == nil {
		params.DataScope = datascope.NewChild(parent.BaseDataScope(), make(map[interface{}]interface{}))
	}
	if params.EventScope == nil {
		params.EventScope = eventscope.NewChild(parent.BaseEventScope())
	}
	if params.Injector == nil {
		params.Injector = parent.BaseInjector()
	}
	if params.Name != "" {
		sid = fmt.Sprintf("%s-%s(%s)", parent.SID(), idutil.StringID(), params.Name)
	} else {
		sid = fmt.Sprintf("%s-%s", parent.SID(), idutil.StringID())
	}
	if params.CID == "" {
		params.CID = parent.CID()
	}
	return &Scope{
		parent:       parent,
		sid:          sid,
		cid:          params.CID,
		ContextScope: params.ContextScope,
		DataScope:    params.DataScope,
		EventScope:   params.EventScope,
		Injector:     params.Injector,
	}
}
