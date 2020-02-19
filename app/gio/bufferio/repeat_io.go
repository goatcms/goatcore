package bufferio

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/gio"
)

// NewRepeatIO returns a new reapat IO instance. Repeat IO repeat input data to output
func NewRepeatIO(params gio.IOParams) (io app.IO) {
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
	buffer := NewBuffer()
	return gio.NewIO(gio.IOParams{
		In:  NewBufferInput(params.In, buffer),
		Out: NewRepeatOutput(params.Out, buffer),
		Err: NewRepeatOutput(params.Err, buffer),
		CWD: params.CWD,
	})
}
