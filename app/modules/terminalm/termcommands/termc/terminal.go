package termc

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/gio"
	"github.com/goatcms/goatcore/app/modules/commonm/commservices"
	"github.com/goatcms/goatcore/app/modules/terminalm/termcommands"
	"github.com/goatcms/goatcore/app/modules/terminalm/termcommands/helpc"
	"github.com/goatcms/goatcore/app/modules/terminalm/termservices"
	"github.com/goatcms/goatcore/app/scope"
	"github.com/goatcms/goatcore/app/scope/contextscope"
)

// RunTerminal run interactive terminal
func RunTerminal(a app.App, ctx app.IOContext) (err error) {
	var (
		deps struct {
			Arguments commservices.Arguments `dependency:"CommonArguments"`
			Terminal  termservices.Terminal  `dependency:"TerminalService"`
		}
		strictMode bool
		io         = a.IOContext().IO()
		scp        = a.IOContext().Scope()
	)
	if err = a.DependencyProvider().InjectTo(&deps); err != nil {
		return err
	}
	strictMode = deps.Arguments.StrictMode()
	exTerm := deps.Terminal.Extends(append(
		Commands(),
		helpc.Commands()...,
	)...)
	for {
		if a.Scopes().App().IsDone() {
			return
		}
		if err = runLoop(a.IOContext(), exTerm); err == nil {
			return nil
		}
		if strictMode {
			scp.AppendError(err)
			return err
		}
		io.Err().Printf("ERROR: %v\n", err)
	}
}

func runLoop(parentCtx app.IOContext, terminal termservices.Terminal) (err error) {
	var (
		isolatedScope app.Scope
		isolatedCtx   app.IOContext
		parentScope   = parentCtx.Scope()
	)
	// Related scope is scope for terminal loop.
	// It is isolated child scope (It contains isolated app.ContextScope).
	// Kill the scope doesn't kill application scope by default.
	// The scope share data with application scope.
	isolatedScope = scope.NewChild(parentScope, scope.ChildParams{
		ContextScope: contextscope.NewIsolated(parentScope),
		DataScope:    parentScope.BaseDataScope(),
		//EventScope: parentScope, <- event scope is not shared to prevent memory leaks
	})
	isolatedCtx = gio.NewIOContext(isolatedScope, parentCtx.IO())
	defer isolatedCtx.Close()
	if err = terminal.RunLoop(isolatedCtx, termcommands.Prompt); err != nil {
		return err
	}
	return isolatedCtx.Scope().Wait()
}
