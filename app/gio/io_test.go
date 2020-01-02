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
	io = NewIO(in, out, eout, cwd)
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

func TestIORequireInput(t *testing.T) {
	t.Parallel()
	var (
		baseIO = newEmptyIO()
	)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	NewIO(nil, baseIO.Out(), baseIO.Err(), baseIO.CWD())
}

func TestIORequireOutput(t *testing.T) {
	t.Parallel()
	var (
		baseIO = newEmptyIO()
	)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	NewIO(baseIO.In(), nil, baseIO.Err(), baseIO.CWD())
}

func TestIORequireErrorOutput(t *testing.T) {
	t.Parallel()
	var (
		baseIO = newEmptyIO()
	)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	NewIO(baseIO.In(), baseIO.Out(), nil, baseIO.CWD())
}

func TestIORequireCWD(t *testing.T) {
	t.Parallel()
	var (
		baseIO app.IO
	)
	baseIO = newEmptyIO()
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	NewIO(baseIO.In(), baseIO.Out(), baseIO.Err(), nil)
}
