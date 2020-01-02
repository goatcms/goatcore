package bufferio

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/gio"
	"github.com/goatcms/goatcore/filesystem"
)

// NewRepeatIO returns a new reapat IO instance. Repeat IO repeat input data to output
func NewRepeatIO(in app.Input, out app.Output, eout app.Output, cwd filesystem.Filespace) (io app.IO) {
	buffer := NewBuffer()
	rin := NewBufferInput(in, buffer)
	rout := NewRepeatOutput(out, buffer)
	reout := NewRepeatOutput(eout, buffer)
	return gio.NewIO(rin, rout, reout, cwd)
}
