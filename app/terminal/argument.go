package terminal

import (
	"github.com/goatcms/goatcore/app"
)

type ArgumentParams struct {
	Name     string
	Help     string
	Required bool
	Type     byte
	Commands app.TerminalCommands
}

// Argument structure implements app.TerminalArgument interface
type Argument struct {
	app.TerminalCommands
	params ArgumentParams
}

func NewArgument(params ArgumentParams) app.TerminalArgument {
	if params.Name == "" {
		panic(ErrArgumentNameIsRequired)
	}
	if params.Type == 0 {
		panic(ErrArgumentTypeIsIncorrect)
	}
	if params.Commands == nil {
		params.Commands = emptyCommands
	}
	return Argument{
		TerminalCommands: params.Commands,
		params:           params,
	}
}

func (reader Argument) Name() string {
	return reader.params.Name
}

func (reader Argument) Help() string {
	return reader.params.Help
}

func (reader Argument) Required() bool {
	return reader.params.Required
}

func (reader Argument) Type() byte {
	return reader.params.Type
}
func (reader Argument) Commands() app.TerminalCommands {
	return reader.params.Commands
}
