package scope

import "errors"

// ErrDoned is the error returned by ContextScope.AddTasks when the context scope is done.
var ErrDoned = errors.New("context scope is done")
