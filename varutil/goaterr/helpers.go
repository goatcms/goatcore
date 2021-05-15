package goaterr

import (
	"fmt"
)

// Errorf formats according to a format specifier and returns the string
// as a value that satisfies error.
func Errorf(format string, a ...interface{}) error {
	return NewError(fmt.Sprintf(format, a...))
}

// Wrapf formats according to a format specifier and returns the string
// as a value that satisfies error. And wrap a based error
func Wrapf(err error, format string, a ...interface{}) error {
	return Wrap(fmt.Sprintf(format, a...), err)
}

// AppendError append error to error collection
func AppendError(errs []error, newerrs ...error) []error {
	for _, err := range newerrs {
		if err == nil {
			continue
		}
		errs = append(errs, err)
	}
	return errs
}
