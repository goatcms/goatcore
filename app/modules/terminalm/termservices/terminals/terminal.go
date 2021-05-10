package terminals

import (
	"io"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/modules/terminalm/termservices"
	"github.com/goatcms/goatcore/app/terminal"
	"github.com/goatcms/goatcore/app/terminal/termexec"
	"github.com/goatcms/goatcore/dependency"
)

// IOTerminal is user communication interface
type IOTerminal struct {
	app      app.App
	commands app.TerminalCommands
}

// NewIOTerminal create new IOTerminal instance for specified app and command scope
func NewIOTerminal(a app.App, commands app.TerminalCommands) termservices.Terminal {
	return termservices.Terminal(&IOTerminal{
		app:      a,
		commands: commands,
	})
}

// IOTerminalFactory create new IOTerminal instance
func IOTerminalFactory(dp dependency.Provider) (in interface{}, err error) {
	var deps struct {
		App app.App `dependency:"App"`
	}
	if err = dp.InjectTo(&deps); err != nil {
		return nil, err
	}
	return termservices.Terminal(&IOTerminal{
		app:      deps.App,
		commands: deps.App.Terminal(),
	}), nil
}

// RunLoop run terminal loop
func (ioterm *IOTerminal) RunLoop(ctx app.IOContext, prompt string) (err error) {
	return termexec.RunLoop(termexec.NewRunCtx(termexec.RunCtxParams{
		Application: ioterm.app,
		Ctx:         ctx,
		Commands:    ioterm.commands,
	}), prompt)
}

// RunString execute single command
func (ioterm *IOTerminal) RunString(ctx app.IOContext, s string) (err error) {
	return termexec.RunString(termexec.NewRunCtx(termexec.RunCtxParams{
		Application: ioterm.app,
		Ctx:         ctx,
		Commands:    ioterm.commands,
	}), s)
}

// RunCommandFromReader execute single command from io.Reader
func (ioterm *IOTerminal) RunCommandFromReader(ctx app.IOContext, reader io.Reader) (eof bool, err error) {
	return termexec.RunCommandFromReader(termexec.NewRunCtx(termexec.RunCtxParams{
		Application: ioterm.app,
		Ctx:         ctx,
		Commands:    ioterm.commands,
	}), reader)
}

// RunCommand execute single command
func (ioterm *IOTerminal) RunCommand(ctx app.IOContext, args []string) (err error) {
	return termexec.RunCommand(termexec.NewRunCtx(termexec.RunCtxParams{
		Application: ioterm.app,
		Ctx:         ctx,
		Commands:    ioterm.commands,
	}), args)
}

// Extends create new terminal with extended command set
func (ioterm *IOTerminal) Extends(exts ...app.TerminalCommand) termservices.Terminal {
	commands := terminal.MergeCommands(ioterm.commands, exts...)
	return NewIOTerminal(ioterm.app, commands)
}
