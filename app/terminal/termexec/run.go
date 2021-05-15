package termexec

import (
	"fmt"
	"io"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/gio"
	"github.com/goatcms/goatcore/app/injector"
	"github.com/goatcms/goatcore/app/scope"
	"github.com/goatcms/goatcore/app/scope/argscope"
	"github.com/goatcms/goatcore/app/scope/datascope"
	"github.com/goatcms/goatcore/varutil"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

type RunCtxParams struct {
	Application app.App
	Ctx         app.IOContext
	Commands    app.TerminalCommands
}

type RunCtx struct {
	application app.App
	ctx         app.IOContext
	commands    app.TerminalCommands
}

func NewRunCtx(params RunCtxParams) (ctx RunCtx) {
	if params.Application == nil {
		panic("Application is required")
	}
	if params.Ctx == nil {
		panic("Ctx is required")
	}
	if params.Commands == nil {
		panic("Commands is required")
	}
	return RunCtx{
		application: params.Application,
		ctx:         params.Ctx,
		commands:    params.Commands,
	}
}

func RunLoop(rctx RunCtx, prompt string) (err error) {
	var (
		in      = rctx.ctx.IO().In()
		out     = rctx.ctx.IO().Out()
		argChan = make(chan []string, 1)
		next    = make(chan struct{}, 1)
	)
	go func() {
		for {
			select {
			case <-rctx.ctx.Scope().Done():
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
						rctx.ctx.Scope().AppendError(err)
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
				argChan <- args
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
		case <-rctx.ctx.Scope().Done():
			return
		case args, more := <-argChan:
			if !more {
				return
			}
			if err = RunCommand(rctx, args); err != nil {
				rctx.ctx.Scope().AppendError(err)
				return
			}
		}
	}
}

func RunString(rctx RunCtx, s string) (err error) {
	var args []string
	if args, _, err = varutil.SplitArguments(s); err != nil {
		return err
	}
	return RunCommand(rctx, args)
}

func RunCommandFromReader(rctx RunCtx, reader io.Reader) (eof bool, err error) {
	var args []string
	if args, eof, err = varutil.ReadArguments(reader); err != nil {
		return eof, err
	}
	return eof, RunCommand(rctx, args)
}

func RunCommand(rctx RunCtx, args []string) (err error) {
	var (
		commandName    string
		command        app.TerminalCommand
		commandContext app.IOContext
	)
	if len(args) != 0 {
		commandName = args[0]
	}
	// find command
	if commandName == "" {
		return goaterr.Errorf("Expected a command")
	}
	if command = rctx.commands.Command(commandName); command == nil {
		return goaterr.Errorf("Error: unknown command %s", commandName)
	}
	//prepare command child context
	argsData := datascope.New(make(map[interface{}]interface{}))
	if err = argscope.InjectArgs(argsData, args...); err != nil {
		return err
	}
	parentScope := rctx.ctx.Scope()
	childScope := scope.NewChild(parentScope, scope.ChildParams{
		DataScope:  parentScope.BaseDataScope(),
		EventScope: parentScope.BaseEventScope(),
		Injector: injector.NewMultiInjector([]app.Injector{
			// reset injectors to application lvl
			rctx.application,
			// add command injector
			datascope.NewInjector("command", argsData),
		}),
		Name: fmt.Sprintf("command:%s", commandName),
	})
	commandContext = gio.NewIOContext(childScope, rctx.ctx.IO())
	// run
	if err = command.Callback()(rctx.application, commandContext); err != nil {
		return goaterr.ToError(goaterr.AppendError([]error{err}, childScope.Close()))
	}
	return childScope.Close()
}
