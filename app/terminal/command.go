package terminal

import (
	"github.com/goatcms/goatcore/app"
)

type CommandParams struct {
	Name      string
	Help      string
	Callback  app.CommandCallback
	Arguments app.TerminalArguments
}

// Command implements app.Command
type Command struct {
	app.TerminalArguments
	params CommandParams
}

// NewCommand create new command
func NewCommand(params CommandParams) (command app.TerminalCommand) {
	if params.Name == "" {
		panic(ErrCommandNameIsRequired)
	}
	if params.Callback == nil {
		panic(ErrCommandCallbackIsRequired)
	}
	return &Command{
		params:            params,
		TerminalArguments: params.Arguments,
	}
}

func (command *Command) Name() string {
	return command.params.Name
}

func (command *Command) Callback() app.CommandCallback {
	return command.params.Callback
}

func (command *Command) Help() string {
	return command.params.Help
}
