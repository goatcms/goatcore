package gio

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/scope"
	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

// IOContext represent task context
type IOContext struct {
	scope app.Scope
	io    app.IO
}

// NewIOContext returns a new IOContext instance.
func NewIOContext(scope app.Scope, io app.IO) (ioc app.IOContext) {
	if scope == nil {
		panic(goaterr.Errorf("gio.IOContext: Scope is required"))
	}
	if io == nil {
		panic(goaterr.Errorf("gio.IOContext: IO is required"))
	}
	return IOContext{
		scope: scope,
		io:    io,
	}
}

// NewChildIOContext extends exist IOContext.
func NewChildIOContext(parent app.IOContext, in app.Input, out app.Output, eout app.Output, cwd filesystem.Filespace) (ioc app.IOContext) {
	var (
		parentIO    app.IO
		childIO     app.IO
		parentScope app.Scope
		childScope  app.Scope
	)
	parentIO = parent.IO()
	if in == nil && out == nil && eout == nil && cwd == nil {
		childIO = parentIO
	} else {
		if in == nil {
			in = parentIO.In()
		}
		if out == nil {
			out = parentIO.Out()
		}
		if eout == nil {
			eout = parentIO.Err()
		}
		if cwd == nil {
			cwd = parentIO.CWD()
		}
		childIO = NewIO(in, out, eout, cwd)
	}
	parentScope = parent.Scope()
	childScope = scope.NewChildScope(parentScope, parentScope, parentScope)
	return IOContext{
		scope: childScope,
		io:    childIO,
	}
}

// Scope return task context scope
func (ioc IOContext) Scope() app.Scope {
	return ioc.scope
}

// IO return task context io
func (ioc IOContext) IO() app.IO {
	return ioc.io
}
