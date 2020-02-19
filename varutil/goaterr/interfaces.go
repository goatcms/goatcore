package goaterr

// MessageError represent interface for error with message
type MessageError interface {
	Message() string
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
