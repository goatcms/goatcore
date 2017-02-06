package goaterr

// Errors represent error collection
type Errors interface {
	error
	Errors() []error
}
