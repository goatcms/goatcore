package terminal

import "errors"

var (
	emptyCommands = newCommands()
)

var (
	ErrArgumentNameIsRequired     = errors.New("Argument name is required")
	ErrArgumentTypeIsIncorrect    = errors.New("Argument type is incorrect")
	ErrArgumentCommandsIsRequired = errors.New("Argument commands is required")
	ErrCommandNameIsRequired      = errors.New("Command name is required")
	ErrCommandCallbackIsRequired  = errors.New("Command callback is required")
	ErrCommandDuplicated          = errors.New("Command is duplicated")
	ErrArgumentDuplicated         = errors.New("Arguement is duplicated")
	ErrHealtDuplicated            = errors.New("HealthCheckerCallback is duplicated")
)
