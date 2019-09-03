package goaterr

import (
	"fmt"
)

// Errorf formats according to a format specifier and returns the string
// as a value that satisfies error.
func Errorf(format string, a []interface{}) error {
	return NewError(fmt.Sprintf(format, a...))
}
