package gio

import (
	"bytes"
	"testing"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/scope"
	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/filesystem/filespace/memfs"
)

func TestIOContext(t *testing.T) {
	t.Parallel()
	var (
		ioScope = scope.NewScope(scope.Params{})
		in      = NewInput(new(bytes.Buffer))
		out     = NewOutput(new(bytes.Buffer))
		eout    = NewOutput(new(bytes.Buffer))
		cwd     filesystem.Filespace
		io      app.IO
		ioc     app.IOContext
		err     error
	)
	if cwd, err = memfs.NewFilespace(); err != nil {
		t.Error(err)
		return
	}
	io = NewIO(IOParams{
		In:  in,
		Out: out,
		Err: eout,
		CWD: cwd,
	})
	ioc = NewIOContext(ioScope, io)
	if ioc.Scope() != ioScope {
		t.Errorf("Expected input from constructor")
	}
	if ioc.IO() != io {
		t.Errorf("Expected output from constructor")
	}
}

func TestIOContextRequireIO(t *testing.T) {
	t.Parallel()
	var (
		io = newEmptyIO()
	)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	NewIOContext(nil, io)
}

func TestIOContextRequireScope(t *testing.T) {
	t.Parallel()
	var (
		ioScope = scope.NewScope(scope.Params{})
	)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	NewIOContext(ioScope, nil)
}
