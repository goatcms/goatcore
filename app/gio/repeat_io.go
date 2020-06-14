package gio

import (
	"github.com/goatcms/goatcore/app"
)

// NewRepeatIO returns a new reapat IO instance. Repeat IO repeat input data to output
func NewRepeatIO(params IOParams) (io app.IO) {
	if params.In == nil {
		panic("Input is required")
	}
	if params.Out == nil {
		panic("Output is required")
	}
	if params.Err == nil {
		panic("Error putput is required")
	}
	if params.CWD == nil {
		panic("CWD is required")
	}
	repeater := NewRepeater(params.Out, params.Err, params.In)
	return NewIO(IOParams{
		In:  repeater,
		Out: repeater,
		Err: repeater.Err(),
		CWD: params.CWD,
	})
}
