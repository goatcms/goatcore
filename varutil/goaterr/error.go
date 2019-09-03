package goaterr

import (
	"errors"
	"runtime/debug"
)

// NewError formats according to a format specifier and returns the string
// as a value that satisfies error.
func NewError(msg string) error {
	return errors.New(msg + " at " + string(debug.Stack()))
}
