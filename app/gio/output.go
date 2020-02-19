package gio

import (
	"github.com/goatcms/goatcore/app"
)

var (
	nilOutputInstance app.Output = NilOutput{}
)

// NilOutput represent empty output
type NilOutput struct{}

// NewNilOutput returns a new NilOutput
func NewNilOutput() app.Output {
	return nilOutputInstance
}

// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
func (out NilOutput) Printf(format string, a ...interface{}) error {
	return nil
}

// Write data to output
func (out NilOutput) Write(p []byte) (n int, err error) {
	return len(p), nil
}
