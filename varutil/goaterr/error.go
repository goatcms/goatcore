package goaterr

import (
	"runtime/debug"
	"strings"
)

// Error is goat error object
type Error struct {
	msg        string
	stack      string
	wraps      []error
	wrapsError error
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
	var msgs []string
	if len(errs) == 0 {
		return nil
	}
	if len(errs) == 1 {
		return errs[0]
	}
	for _, err := range errs {
		if messageError, ok := err.(MessageError); ok {
			msgs = append(msgs, messageError.Message())
		} else {
			msgs = append(msgs, err.Error())
		}
	}
	return &Error{
		msg:   strings.Join(msgs, " && "),
		stack: string(debug.Stack()),
		wraps: errs,
	}
}

// Error return goat error object with title and stacktrace
func (err *Error) Error() string {
	return print(err)
}

// Unwrap return wrapped error
func (err *Error) Unwrap() error {
	if err.wrapsError == nil {
		err.wrapsError = ToError(err.wraps)
	}
	return err.wrapsError
}

// UnwrapAll return wrapped errors
func (err *Error) UnwrapAll() []error {
	return err.wraps
}

// Stack return error stack trace
func (err *Error) Stack() string {
	return err.stack
}

// Message return error message
func (err *Error) Message() string {
	return err.msg
}
