package gio

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/scope"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

// ChildIOContextParams describe child context
type ChildIOContextParams struct {
	Scope scope.ChildParams
	IO    IOParams
}

// IOContextParams describe context
type IOContextParams struct {
	Scope scope.Params
	IO    IOParams
}

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
func NewChildIOContext(parent app.IOContext, params ChildIOContextParams) (ioc app.IOContext) {
	var (
		parentIO app.IO
		childIO  app.IO
	)
	parentIO = parent.IO()
	if params.IO.In == nil && params.IO.Out == nil && params.IO.Err == nil && params.IO.CWD == nil {
		childIO = parentIO
	} else {
		if params.IO.In == nil {
			params.IO.In = parentIO.In()
		}
		if params.IO.Out == nil {
			params.IO.Out = parentIO.Out()
		}
		if params.IO.Err == nil {
			params.IO.Err = parentIO.Err()
		}
		if params.IO.CWD == nil {
			params.IO.CWD = parentIO.CWD()
		}
		childIO = NewIO(params.IO)
	}
	return IOContext{
		scope: scope.NewChildScope(parent.Scope(), params.Scope),
		io:    childIO,
	}
}

// NewParallelIOContext create new parallel io context related to current.
func NewParallelIOContext(parent app.IOContext, params IOContextParams) (ioc app.IOContext) {
	var (
		parentIO app.IO
		childIO  app.IO
	)
	parentIO = parent.IO()
	if params.IO.In == nil && params.IO.Out == nil && params.IO.Err == nil && params.IO.CWD == nil {
		childIO = parentIO
	} else {
		if params.IO.In == nil {
			params.IO.In = parentIO.In()
		}
		if params.IO.Out == nil {
			params.IO.Out = parentIO.Out()
		}
		if params.IO.Err == nil {
			params.IO.Err = parentIO.Err()
		}
		if params.IO.CWD == nil {
			params.IO.CWD = parentIO.CWD()
		}
		childIO = NewIO(params.IO)
	}
	return IOContext{
		scope: scope.NewParallelScope(parent.Scope(), params.Scope),
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

// Close elements (like scope)
func (ioc IOContext) Close() (err error) {
	return ioc.scope.Close()
}
