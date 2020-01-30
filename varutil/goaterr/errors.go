package goaterr

import (
	"encoding/json"
	"runtime/debug"
)

// ErrorCollection is a basic Errors interface implementation
type ErrorCollection struct {
	errs       []error
	Errs       []string `json:"errors"`
	Childs     []Errors `json:"childs"`
	StackTrace string   `json:"stackTrace"`
}

// ErrorCollectionJSON represent errors json structure
type ErrorCollectionJSON struct {
}

// NewErrorCollection create a new errors instance
func NewErrorCollection(errs []error) Errors {
	e := &ErrorCollection{
		errs:       errs,
		Errs:       []string{},
		Childs:     []Errors{},
		StackTrace: string(debug.Stack()),
	}
	for _, err := range errs {
		switch v := err.(type) {
		case Errors:
			e.Childs = append(e.Childs, v)
		default:
			e.Errs = append(e.Errs, v.Error())
		}
	}
	return e
}

// ToError return error object if error list is not empty or nil.
// Otherwise return error object.
func ToError(errs []error) error {
	if errs == nil || len(errs) == 0 {
		return nil
	}
	if len(errs) == 1 {
		return errs[0]
	}
	return NewErrorCollection(errs)
}

// Errors return sub error collection
func (e *ErrorCollection) Errors() []error {
	return e.errs
}

// Error return error message
func (e *ErrorCollection) Error() string {
	return e.JSON()
}

// String method convert object to string
func (e *ErrorCollection) String() (out string) {
	return e.JSON()
}

// JSON method convert object to JSON
func (e *ErrorCollection) JSON() (out string) {
	var (
		jsonv []byte
		err   error
	)
	if jsonv, err = json.MarshalIndent(e, "", "  "); err != nil {
		panic(err)
	}
	return string(jsonv)
}
