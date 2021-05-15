package terminalm

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/gio"
	"github.com/goatcms/goatcore/app/modules/commonm/commservices"
	"github.com/goatcms/goatcore/app/modules/terminalm/termcommands/termc"
	"github.com/goatcms/goatcore/app/modules/terminalm/termservices"
	"github.com/goatcms/goatcore/app/modules/terminalm/termservices/terminals"
	"github.com/goatcms/goatcore/app/scope"
	"github.com/goatcms/goatcore/app/scope/contextscope"
)

// Module is command unit
type Module struct{}

// NewModule create new command module instance
func NewModule() app.Module {
	return &Module{}
}

// RegisterDependencies is init callback to register module dependencies
func (m *Module) RegisterDependencies(a app.App) error {
	dp := a.DependencyProvider()
	return dp.AddDefaultFactory(termservices.TerminalService, terminals.IOTerminalFactory)
}

// InitDependencies is init callback to inject dependencies inside module
func (m *Module) InitDependencies(a app.App) error {
	return nil
}

// Run start command line loop
func (m *Module) Run(a app.App) (err error) {
	var (
		deps struct {
			Arguments commservices.Arguments `dependency:"CommonArguments"`
			Terminal  termservices.Terminal  `dependency:"TerminalService"`
		}
		exTerm     termservices.Terminal
		io         = a.IOContext().IO()
		strictMode bool
	)
	if err = a.DependencyProvider().InjectTo(&deps); err != nil {
		return err
	}
	strictMode = deps.Arguments.StrictMode()
	exTerm = deps.Terminal.Extends(termc.Commands()...)
	args := a.Arguments()
	if len(args) < 2 {
		return exTerm.RunCommand(a.IOContext(), []string{"help"})
	} else if args[1] != "terminal" {
		return exTerm.RunCommand(a.IOContext(), args[1:])
	}
	for {
		if a.Scopes().App().IsDone() {
			return
		}
		if err = m.runLoop(a.IOContext(), deps.Terminal); err == nil {
			return nil
		}
		if strictMode {
			a.IOContext().Scope().AppendError(err)
			return err
		}
		io.Err().Printf("ERROR: %v\n", err)
	}
}

func (m *Module) runLoop(parentCtx app.IOContext, terminal termservices.Terminal) (err error) {
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
	if err = terminal.RunLoop(isolatedCtx, ">"); err != nil {
		return err
	}
	return isolatedCtx.Scope().Wait()
}
