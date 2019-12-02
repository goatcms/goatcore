package goaterr

import (
	"runtime/debug"
)

// ErrorWrapper contains other error
type ErrorWrapper struct {
	msg string
	err error
}

// NewErrorWrapper wrap a error
func NewErrorWrapper(msg string, err error) error {
	return &ErrorWrapper{
		msg: msg + "    wrap " + err.Error() + "    at " + string(debug.Stack()),
		err: err,
	}
}

// Error return error message
func (e ErrorWrapper) Error() string {
	return e.msg
}

// Unwrap return wrapped error
func (e ErrorWrapper) Unwrap() error {
	return e.err
}
