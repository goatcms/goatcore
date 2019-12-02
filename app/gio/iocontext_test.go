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
		ioScope = scope.NewScope("tag")
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
	if io, err = NewIO(in, out, eout, cwd); err != nil {
		t.Error(err)
		return
	}
	if ioc, err = NewIOContext(ioScope, io); err != nil {
		t.Error(err)
		return
	}
	if ioc.Scope() != ioScope {
		t.Errorf("Expected input from constructor")
	}
	if ioc.IO() != io {
		t.Errorf("Expected output from constructor")
	}
}

func TestIOContextRequireAllAttributes(t *testing.T) {
	t.Parallel()
	var (
		ioScope = scope.NewScope("tag")
		in      = NewInput(new(bytes.Buffer))
		out     = NewOutput(new(bytes.Buffer))
		eout    = NewOutput(new(bytes.Buffer))
		cwd     filesystem.Filespace
		io      app.IO
		err     error
	)
	if cwd, err = memfs.NewFilespace(); err != nil {
		t.Error(err)
		return
	}
	if io, err = NewIO(in, out, eout, cwd); err != nil {
		t.Error(err)
		return
	}
	if _, err = NewIOContext(nil, io); err == nil {
		t.Errorf("Scope is required")
	}
	if _, err = NewIOContext(ioScope, nil); err == nil {
		t.Errorf("IO is required")
	}
}
