package gio

import (
	"fmt"
	"io"
	"sync"

	"github.com/goatcms/goatcore/app"
)

// Output represent system output
type Output struct {
	wd io.Writer // writer provided by the client
	mu sync.Mutex
}

// NewOutput returns a new Output.
func NewOutput(wd io.Writer) *Output {
	return &Output{
		wd: wd,
	}
}

// NewAppOutput returns a new app.Output.
func NewAppOutput(wd io.Writer) app.Output {
	return NewOutput(wd)
}

// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
func (out *Output) Printf(format string, a ...interface{}) error {
	out.mu.Lock()
	defer out.mu.Unlock()
	s := fmt.Sprintf(format, a...)
	_, err := out.wd.Write([]byte(s))
	return err
}

// Write data to output
func (out *Output) Write(p []byte) (n int, err error) {
	out.mu.Lock()
	defer out.mu.Unlock()
	return out.wd.Write(p)
}
