package gio

import (
	"bytes"
	"testing"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/filesystem/filespace/memfs"
)

func TestIO(t *testing.T) {
	t.Parallel()
	var (
		in   = NewInput(new(bytes.Buffer))
		out  = NewOutput(new(bytes.Buffer))
		eout = NewOutput(new(bytes.Buffer))
		cwd  filesystem.Filespace
		io   app.IO
		err  error
	)
	if cwd, err = memfs.NewFilespace(); err != nil {
		t.Error(err)
		return
	}
	if io, err = NewIO(in, out, eout, cwd); err != nil {
		t.Error(err)
		return
	}
	if io.In() != in {
		t.Errorf("Expected input from constructor")
	}
	if io.Out() != out {
		t.Errorf("Expected output from constructor")
	}
	if io.Err() != eout {
		t.Errorf("Expected error output from constructor")
	}
	if io.CWD() != cwd {
		t.Errorf("Expected Current Working Directory from constructor")
	}
}

func TestIORequireAllAttributes(t *testing.T) {
	t.Parallel()
	var (
		in   = NewInput(new(bytes.Buffer))
		out  = NewOutput(new(bytes.Buffer))
		eout = NewOutput(new(bytes.Buffer))
		cwd  filesystem.Filespace
		err  error
	)
	if cwd, err = memfs.NewFilespace(); err != nil {
		t.Error(err)
		return
	}
	if _, err = NewIO(in, out, eout, cwd); err != nil {
		t.Error(err)
		return
	}
	if _, err = NewIO(nil, out, eout, cwd); err == nil {
		t.Errorf("Input is required")
	}
	if _, err = NewIO(in, nil, eout, cwd); err == nil {
		t.Errorf("Output is required")
	}
	if _, err = NewIO(in, out, nil, cwd); err == nil {
		t.Errorf("Error output is required")
	}
	if _, err = NewIO(in, out, eout, nil); err == nil {
		t.Errorf("Current Working Directory is required")
	}
}
