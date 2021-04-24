package terminalm

import (
	"io"
	"strings"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/gio"
	"github.com/goatcms/goatcore/app/injector"
	"github.com/goatcms/goatcore/app/modules"
	"github.com/goatcms/goatcore/app/scope"
	"github.com/goatcms/goatcore/app/scope/argscope"
	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcore/varutil"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

type arguments struct {
	args []string
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
		in      = ctx.IO().In()
		out     = ctx.IO().Out()
		argChan = make(chan arguments, 1)
		next    = make(chan struct{}, 1)
	)
	go func() {
		for {
			select {
			case <-ctx.Scope().Done():
				return
			case _, more := <-next:
				var (
					args []string
					eof  bool
					err  error
				)
				if !more {
					return
				}
				for {
					if prompt != "" {
						out.Printf(prompt)
					}
					if args, eof, err = varutil.ReadArguments(in); err != nil {
						ctx.Scope().AppendError(err)
						return
					}
					if len(args) != 0 {
						break
					}
					if eof {
						close(argChan)
						return
					}
				}
				argChan <- arguments{
					args: args,
				}
				if eof {
					close(argChan)
					return
				}
			}
		}
	}()
	defer func() {
		close(next)
	}()
	for {
		next <- struct{}{}
		select {
		case <-ctx.Scope().Done():
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
	if cbIns = terminal.commandScope.Value("command." + commandName); cbIns == nil {
		return goaterr.Errorf("Error: unknown command %s", commandName)
	}
	cb = cbIns.(app.CommandCallback)
	//prepare command child context
	argsData := &scope.DataScope{
		Data: make(map[interface{}]interface{}),
	}
	if err = argscope.InjectArgsToScope(args, argsData); err != nil {
		return err
	}
	baseScope := ctx.Scope()
	injectableScope := scope.NewScope(scope.Params{
		DataScope:    baseScope,
		EventScope:   baseScope,
		ContextScope: baseScope,
		Injector: injector.NewMultiInjector([]app.Injector{
			// reset injector to AppScope lvl
			terminal.app.AppScope(),
			// add command injector
			scope.NewScopeInjector("command", argsData),
		}),
	})
	commandContext = gio.NewIOContext(injectableScope, ctx.IO())
	// run
	return cb(terminal.app, commandContext)
}
