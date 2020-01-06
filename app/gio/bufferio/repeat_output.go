package bufferio

import (
	"fmt"

	"github.com/goatcms/goatcore/app"
)

// RepeatOutput represent system output
type RepeatOutput struct {
	parent app.Output
	buffer *Buffer
}

// NewRepeatOutput returns a new RepeatOutput.
func NewRepeatOutput(parent app.Output, buffer *Buffer) app.Output {
	return &RepeatOutput{
		parent: parent,
		buffer: buffer,
	}
}

func newAppRepeatOutput(parent app.Output, buffer *Buffer) app.Output {
	return NewRepeatOutput(parent, buffer)
}

// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
func (out *RepeatOutput) Printf(format string, a ...interface{}) error {
	prompted := out.Prompt(fmt.Sprintf(format, a...))
	return out.parent.Printf(prompted)
}

// Write data to output
func (out *RepeatOutput) Write(p []byte) (n int, err error) {
	prompted := out.Prompt(string(p))
	return out.parent.Write([]byte(prompted))
}

// Prompt add prompt to output
func (out *RepeatOutput) Prompt(s string) (result string) {
	prompt := out.buffer.ReadAndClean()
	if prompt == "" {
		return fmt.Sprintf("%s", s)
	}
	return fmt.Sprintf("> %s : \n%s", prompt, s)
}
