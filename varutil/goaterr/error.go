package goaterr

import (
	"fmt"
	"runtime/debug"
)

// Error is goat error object
type Error struct {
	msg   string
	stack string
	wraps []error
}

// newGoatError return goat error object with title and stacktrace
func newGoatError(msg string, wraps []error) *Error {
	return &Error{
		msg:   msg,
		stack: string(debug.Stack()),
		wraps: wraps,
	}
}

// NewError formats according to a format specifier and returns the string
// as a value that satisfies error.
func NewError(msg string) error {
	return &Error{
		msg:   msg,
		stack: string(debug.Stack()),
	}
}

// Wrap function returns a new error that adds context to the original error by
// recording a stack trace at the point Wrap is called, together with the supplied message.
func Wrap(err error, msg string) error {
	return &Error{
		msg:   msg,
		stack: string(debug.Stack()),
		wraps: []error{err},
	}
}

// ToError return error object if error list is not empty or nil.
// Otherwise return error object.
func ToError(errs []error) error {
	if len(errs) == 0 {
		return nil
	}
	if len(errs) == 1 {
		return errs[0]
	}
	return &Error{
		msg:   fmt.Sprintf("Error wrapper (Contains %v errors).", len(errs)),
		stack: string(debug.Stack()),
		wraps: errs,
	}
}

// Error return goat error object with title and stacktrace
func (err *Error) Error() (s string) {
	s = err.msg
	if len(err.wraps) > 0 {
		s += fmt.Sprintf(" Wraps %v errors:", len(err.wraps))
		for _, e := range err.wraps {
			switch v := e.(type) {
			case MessageError:
				s += fmt.Sprintf("\n - %s", v.ErrorMessage())
			default:
				s += fmt.Sprintf("\n - %s", v.Error())
			}
		}
	}
	return
}

// Unwrap return wrapped error
func (err *Error) Unwrap() error {
	if len(err.wraps) < 1 {
		return nil
	}
	return err.wraps[0]
}

// UnwrapAll return wrapped errors
func (err *Error) UnwrapAll() []error {
	return err.wraps
}

// Stack return error stack trace
func (err *Error) Stack() string {
	return err.stack
}

// ErrorJSON return error json tree
func (err *Error) ErrorJSON() string {
	return printJSON(err)
}

// ErrorMessage return error json tree
func (err *Error) ErrorMessage() string {
	return err.msg
}
