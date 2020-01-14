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
	io = NewIO(IOParams{
		In:  in,
		Out: out,
		Err: eout,
		CWD: cwd,
	})
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
	NewIO(IOParams{
		In:  nil,
		Out: baseIO.Out(),
		Err: baseIO.Err(),
		CWD: baseIO.CWD(),
	})
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
	NewIO(IOParams{
		In:  baseIO.In(),
		Out: nil,
		Err: baseIO.Err(),
		CWD: baseIO.CWD(),
	})
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
	NewIO(IOParams{
		In:  baseIO.In(),
		Out: baseIO.Out(),
		Err: nil,
		CWD: baseIO.CWD(),
	})
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
	NewIO(IOParams{
		In:  baseIO.In(),
		Out: baseIO.Out(),
		Err: baseIO.Err(),
		CWD: nil,
	})
}
