package goaterr

import (
	"fmt"
	"runtime/debug"
)

// ErrorCollection is a basic Errors interface implementation
type ErrorCollection struct {
	errs       []error
	stackTrace string
}

// NewErrors create a new errors instance
func NewErrors(errs []error) Errors {
	return ErrorCollection{
		errs:       errs,
		stackTrace: string(debug.Stack()),
	}
}

// ToErrors return error object if error list is not empty or nil.
// Otherwise return error object.
func ToErrors(errs []error) error {
	if errs == nil || len(errs) == 0 {
		return nil
	}
	if len(errs) == 1 {
		return errs[0]
	}
	return NewErrors(errs)
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

// Error return error message
func (e ErrorCollection) Error() string {
	return e.String()
}

// String method convert object to string
func (e ErrorCollection) String() (out string) {
	for _, err := range e.errs {
		out += fmt.Sprintf("   %s\n", err.Error())
	}
	return fmt.Sprintf("\n%s {\n%s\n}\n", e.stackTrace, out)
}

// Errors return sub error collection
func (e ErrorCollection) Errors() []error {
	return e.errs
}
