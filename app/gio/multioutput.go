package gio

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

// MultiOutput is helper to brodcast call to many writers
type MultiOutput struct {
	outs []app.Output
}

// NewMultiOutput return new MultiOutput instance
func NewMultiOutput(outs []app.Output) app.Output {
	return &MultiOutput{
		outs: outs,
	}
}

// Writer is the interface that wraps the basic Write method.
func (broadcast *MultiOutput) Write(p []byte) (n int, err error) {
	for _, out := range broadcast.outs {
		if n, err = out.Write(p); err != nil {
			return n, err
		}
		if n != len(p) {
			return n, goaterr.Errorf("Can not write %d bytes (%d bytes writen)", len(p), n)
		}
	}
	return n, err
}

// Printf print to multiple outputs.
func (broadcast *MultiOutput) Printf(format string, a ...interface{}) (err error) {
	for _, out := range broadcast.outs {
		if err = out.Printf(format, a...); err != nil {
			return err
		}
	}
	return nil
}
