package bufferio

import (
	"github.com/goatcms/goatcore/app"
)

// BufferInput cache input data
type BufferInput struct {
	buffer *Buffer
	parent app.Input
}

// NewBufferInput retur new chached input
func NewBufferInput(parent app.Input, buffer *Buffer) app.Input {
	return &BufferInput{
		parent: parent,
		buffer: buffer,
	}
}

// ReadWord return next word from input stream
func (input *BufferInput) ReadWord() (s string, err error) {
	if s, err = input.parent.ReadWord(); err == nil {
		input.buffer.WriteString(s)
	}
	return s, err
}

// ReadLine return next line from input stream
func (input *BufferInput) ReadLine() (s string, err error) {
	if s, err = input.parent.ReadWord(); err == nil {
		input.buffer.WriteString(s)
	}
	return s, err
}

func (input *BufferInput) Read(p []byte) (n int, err error) {
	if n, err = input.parent.Read(p); n > 0 {
		input.buffer.WriteString(string(p[:n]))
	}
	return n, err
}
