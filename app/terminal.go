package app

const (
	TerminalUndefinedArgument = 0
	TerminalIntArgument       = 1
	TerminalFloatArgument     = iota
	TerminalTextArgument      = iota
	TerminalBoolArgument      = iota
	TerminalPIPArgument       = iota
	TerminalOtherArgument     = iota
	TerminalURLArgument       = iota
	TerminalEmailrgument      = iota
)

func NilCommandCallback(App, IOContext) (err error) {
	return nil
}

// CommandCallback is function call to run user command
type CommandCallback func(App, IOContext) (err error)

// TerminalArgument is an argument read-only interface
type TerminalArgument interface {
	TerminalCommands
	Name() string
	Help() string
	Required() bool
	Type() byte
}

// TerminalCommand is a command read-only interface
type TerminalCommand interface {
	TerminalArguments
	Callback() CommandCallback
	Help() string
	MainArguments() []string
	Name() string
}

// ArgumentReader is interface provided arguments set
type TerminalArguments interface {
	ArgumentNames() []string
	Argument(name string) TerminalArgument
}

// TerminalCommands is interface provided commands reader
type TerminalCommands interface {
	CommandNames() []string
	Command(name string) TerminalCommand
}

// Terminal is interface represent terminal definition
type Terminal interface {
	TerminalArguments
	TerminalCommands
}

// TerminalManager is interface represent terminal definition
type TerminalManager interface {
	Terminal
	SetArgument(arg TerminalArgument)
	SetCommand(command ...TerminalCommand)
}
