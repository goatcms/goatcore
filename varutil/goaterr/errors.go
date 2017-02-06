package goaterr

import "fmt"

// ErrorCollection is a basic Errors interface implementation
type ErrorCollection struct {
	errs []error
}

// NewErrors create a new errors instance
func NewErrors(errs []error) Errors {
	return ErrorCollection{
		errs: errs,
	}
}

// Error return error message
func (e ErrorCollection) Error() string {
	return e.String()
}

// String method convert object to string
func (e ErrorCollection) String() string {
	return fmt.Sprintf("errors: %v", e.errs)
}

// Errors return sub error collection
func (e ErrorCollection) Errors() []error {
	return e.errs
}
