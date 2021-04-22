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

type arguments struct {
	args []string
	eof  bool
}

// IOTerminal is user communication interface
type IOTerminal struct {
	app          app.App
	commandScope app.Scope
}

// NewIOTerminal create new IOTerminal instance for specified app and command scope
func NewIOTerminal(a app.App, commandScope app.Scope) (term modules.Terminal, err error) {
	return modules.Terminal(&IOTerminal{
		app:          a,
		commandScope: commandScope,
	}), nil
}

// IOTerminalFactory create new IOTerminal instance
func IOTerminalFactory(dp dependency.Provider) (in interface{}, err error) {
	var deps struct {
		App app.App `dependency:"App"`
	}
	if err = dp.InjectTo(&deps); err != nil {
		return nil, err
	}
	return modules.Terminal(&IOTerminal{
		app:          deps.App,
		commandScope: deps.App.CommandScope(),
	}), nil
}

// RunLoop run terminal loop
func (terminal *IOTerminal) RunLoop(ctx app.IOContext, prompt string) (err error) {
	var (
		in       = ctx.IO().In()
		out      = ctx.IO().Out()
		argChan  = make(chan arguments, 1)
		doneChan = make(chan bool, 1)
	)
	go func() {
		for {
			select {
			case <-ctx.Scope().Context().Done():
				return
			case <-doneChan:
				return
			default:
				if prompt != "" {
					out.Printf(prompt)
				}
				args, eof, err := varutil.ReadArguments(in)
				if err != nil {
					ctx.Scope().AppendError(err)
					return
				}
				if len(args) != 0 {
					argChan <- arguments{
						args: args,
						eof:  eof,
					}
				}
				if eof {
					close(argChan)
					return
				}
			}
		}
	}()
	defer func() {
		doneChan <- true
	}()
	for {
		select {
		case <-ctx.Scope().Context().Done():
			return
		case a, more := <-argChan:
			if !more {
				return
			}
			if err = terminal.RunCommand(ctx, a.args); err != nil {
				ctx.Scope().AppendError(err)
				return
			}
		}
	}
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
		commandContext app.IOContext
	)
	if len(args) != 0 {
		commandName = strings.ToLower(args[0])
	}
	// find command
	if commandName == "" {
		return HelpComamnd(terminal.app, ctx)
	}
	if cbIns, err = terminal.commandScope.Get("command." + commandName); err != nil || cbIns == nil {
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
	return cb(terminal.app, commandContext)
}
