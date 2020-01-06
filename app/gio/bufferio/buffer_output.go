package bufferio

import (
	"fmt"

	"github.com/goatcms/goatcore/app"
)

// BufferOutput represent system output
type BufferOutput struct {
	parent app.Output
	buffer *Buffer
}

// NewBufferOutput returns a new BufferOutput.
func NewBufferOutput(buffer *Buffer) app.Output {
	return &BufferOutput{
		buffer: buffer,
	}
}

// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
func (out *BufferOutput) Printf(format string, a ...interface{}) error {
	return out.buffer.WriteString(fmt.Sprintf(format, a...))
}

// Write data to output
func (out *BufferOutput) Write(p []byte) (n int, err error) {
	return out.buffer.Write(p)
}
