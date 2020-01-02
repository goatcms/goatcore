package gio

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

// IO represent task context input and output
type IO struct {
	in   app.Input
	out  app.Output
	eout app.Output
	cwd  filesystem.Filespace
}

// NewIO returns a new IO instance.
func NewIO(in app.Input, out app.Output, eout app.Output, cwd filesystem.Filespace) (io app.IO) {
	if in == nil {
		panic(goaterr.Errorf("gio.IO: Input is required"))
	}
	if out == nil {
		panic(goaterr.Errorf("gio.IO: Output is required"))
	}
	if eout == nil {
		panic(goaterr.Errorf("gio.IO: Error output is required"))
	}
	if cwd == nil {
		panic(goaterr.Errorf("gio.IO: CWD (Current Working Directory) is required"))
	}
	return IO{
		in:   in,
		out:  out,
		eout: eout,
		cwd:  cwd,
	}
}

// In return default application input
func (io IO) In() app.Input {
	return io.in
}

// Out return default application output
func (io IO) Out() app.Output {
	return io.out
}

// Err return default application error output
func (io IO) Err() app.Output {
	return io.eout
}

// CWD return Current working directory
func (io IO) CWD() filesystem.Filespace {
	return io.cwd
}
