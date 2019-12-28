package terminalm

import (
	"io"
	"strings"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/modules"
	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcore/varutil"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

// IOTerminal is user communication interface
type IOTerminal struct {
	deps struct {
		App app.App `dependency:"App"`
	}
}

// IOTerminalFactory create new IOTerminal instance
func IOTerminalFactory(dp dependency.Provider) (in interface{}, err error) {
	instance := &IOTerminal{}
	if err = dp.InjectTo(&instance.deps); err != nil {
		return nil, err
	}
	return modules.Terminal(instance), nil
}

// RunLoop run terminal loop
func (terminal *IOTerminal) RunLoop(ctx app.IOContext) (err error) {
	var (
		args []string
		eof  = false
		io   = ctx.IO()
	)
	for !eof {
		io.Out().Printf("\n>")
		if args, eof, err = varutil.ReadArguments(io.In()); err != nil {
			return err
		}
		if len(args) == 0 {
			continue
		}
		if err = terminal.RunCommand(ctx, args); err != nil {
			return err
		}
	}
	return nil
}

// RunString execute single command
func (terminal *IOTerminal) RunString(ctx app.IOContext, s string) (err error) {
	var args []string
	if args, _, err = varutil.SplitArguments(s); err != nil {
		return err
	}
	return terminal.RunCommand(ctx, args)
}

// RunCommandFromReader execute single command from io.Reader
func (terminal *IOTerminal) RunCommandFromReader(ctx app.IOContext, reader io.Reader) (eof bool, err error) {
	var args []string
	if args, eof, err = varutil.ReadArguments(reader); err != nil {
		return eof, err
	}
	return eof, terminal.RunCommand(ctx, args)
}

// RunCommand execute single command
func (terminal *IOTerminal) RunCommand(ctx app.IOContext, args []string) (err error) {
	var (
		commandName    string
		cbIns          interface{}
		cb             app.CommandCallback
		commandScope   = terminal.deps.App.CommandScope()
		commandContext app.IOContext
	)
	if len(args) != 0 {
		commandName = strings.ToLower(args[0])
	}
	if commandName == "" {
		return HelpComamnd(terminal.deps.App, ctx)
	}
	if cbIns, err = commandScope.Get("command." + commandName); err != nil || cbIns == nil {
		return goaterr.Errorf("Error: unknown command %s", commandName)
	}
	cb = cbIns.(app.CommandCallback)
	if commandContext, err = newCommandContext(ctx, args); err != nil {
		return err
	}
	return cb(terminal.deps.App, commandContext)
}
