package gio

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

// IOContext represent task context
type IOContext struct {
	scope app.Scope
	io    app.IO
}

// NewIOContext returns a new IOContext instance.
func NewIOContext(scope app.Scope, io app.IO) (ioc app.IOContext, err error) {
	if scope == nil {
		return nil, goaterr.Errorf("gio.IOContext: Scope is required")
	}
	if io == nil {
		return nil, goaterr.Errorf("gio.IOContext: IO is required")
	}
	return IOContext{
		scope: scope,
		io:    io,
	}, nil
}

// Scope return task context scope
func (ioc IOContext) Scope() app.Scope {
	return ioc.scope
}

// IO return task context io
func (ioc IOContext) IO() app.IO {
	return ioc.io
}
