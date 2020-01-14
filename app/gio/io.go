package gio

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

// IOParams describe io
type IOParams struct {
	In  app.Input
	Out app.Output
	Err app.Output
	CWD filesystem.Filespace
}

// IO represent task context input and output
type IO struct {
	in   app.Input
	out  app.Output
	eout app.Output
	cwd  filesystem.Filespace
}

// NewIO returns a new IO instance.
func NewIO(params IOParams) (io app.IO) {
	if params.In == nil {
		panic(goaterr.Errorf("gio.IO: Input is required"))
	}
	if params.Out == nil {
		panic(goaterr.Errorf("gio.IO: Output is required"))
	}
	if params.Err == nil {
		panic(goaterr.Errorf("gio.IO: Error output is required"))
	}
	if params.CWD == nil {
		panic(goaterr.Errorf("gio.IO: CWD (Current Working Directory) is required"))
	}
	return IO{
		in:   params.In,
		out:  params.Out,
		eout: params.Err,
		cwd:  params.CWD,
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
