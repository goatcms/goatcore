package terminalm

import (
	"io"
	"strings"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/gio"
	"github.com/goatcms/goatcore/app/modules"
	"github.com/goatcms/goatcore/app/scope"
	"github.com/goatcms/goatcore/app/scope/argscope"
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
func (terminal *IOTerminal) RunLoop(ctx app.IOContext, prompt string) (err error) {
	var (
		args []string
		eof  = false
		io   = ctx.IO()
	)
	for !eof {
		if ctx.Scope().IsKilled() {
			return ctx.Scope().ToError()
		}
		if prompt != "" {
			io.Out().Printf(prompt)
		}
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
	return err
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
	// find command
	if commandName == "" {
		return HelpComamnd(terminal.deps.App, ctx)
	}
	if cbIns, err = commandScope.Get("command." + commandName); err != nil || cbIns == nil {
		return goaterr.Errorf("Error: unknown command %s", commandName)
	}
	cb = cbIns.(app.CommandCallback)
	//prepare command child context
	argsData := &scope.DataScope{
		Data: make(map[string]interface{}),
	}
	if err = argscope.InjectArgsToScope(args, argsData); err != nil {
		return err
	}
	baseScope := ctx.Scope()
	injectableScope := scope.NewScope(scope.Params{
		DataScope:  baseScope,
		EventScope: baseScope,
		SyncScope:  baseScope,
		Injectors: []app.Injector{
			argsData.Injector("command"),
		},
	})
	commandContext = gio.NewIOContext(injectableScope, ctx.IO())
	// run
	return cb(terminal.deps.App, commandContext)
}
