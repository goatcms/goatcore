package gio

import (
	"io"

	"github.com/goatcms/goatcore/app"
)

var (
	nilInputInstance app.Input = NilInput{}
)

// NilInput represent empty input
type NilInput struct{}

// NewNilInput returns a new NilInput
func NewNilInput() app.Input {
	return nilInputInstance
}

// Read return EOF
func (in NilInput) Read(p []byte) (n int, err error) {
	return 0, io.EOF
}

// ReadWord return EOF
func (in NilInput) ReadWord() (string, error) {
	return "", io.EOF
}

// ReadLine return EOF
func (in NilInput) ReadLine() (string, error) {
	return "", io.EOF
}
