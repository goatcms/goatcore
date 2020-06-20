package goaterr

// MessageError represent interface for error with message
type MessageError interface {
	ErrorMessage() string
}

// JSONError represent interface for error serialized to json
type JSONError interface {
	ErrorJSON() string
}

// CodeError represent interface for error with code
type CodeError interface {
	ErrorCode() int
}

// TrackedError is described tracked errors interface
type TrackedError interface {
	Stack() string
}

// ErrorsWrapper represent interface for error wrapper
type ErrorsWrapper interface {
	UnwrapAll() []error
}

// ErrorWrapper represent interface for error wrapper
type ErrorWrapper interface {
	Unwrap() error
}

// GoatError represent interface for goat error
type GoatError interface {
	JSONError
	CodeError
	TrackedError
	ErrorsWrapper
}
