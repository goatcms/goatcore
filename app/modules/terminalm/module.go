package terminalm

import (
	"strings"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/gio"
	"github.com/goatcms/goatcore/app/modules"
	"github.com/goatcms/goatcore/app/scope"
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
	dp.AddDefaultFactory(modules.TerminalService, IOTerminalFactory)
	app.RegisterCommand(a, "health", HealthComamnd, "chack and show application health")
	app.RegisterCommand(a, "help", HelpComamnd, "Show help")
	return nil
}

// InitDependencies is init callback to inject dependencies inside module
func (m *Module) InitDependencies(a app.App) error {
	return nil
}

// Run start command line loop
func (m *Module) Run(a app.App) (err error) {
	var (
		deps struct {
			Terminal   modules.Terminal `dependency:"TerminalService"`
			StrictMode string           `argument:"?strict"`
			ctx        app.IOContext
		}
		io         = a.IOContext().IO()
		strictMode bool
	)
	if err = a.DependencyProvider().InjectTo(&deps); err != nil {
		return err
	}
	strictMode = strings.ToLower(deps.StrictMode) == "true"
	args := a.Arguments()
	if len(args) < 2 {
		return deps.Terminal.RunCommand(a.IOContext(), []string{"help"})
	} else if args[1] != "terminal" {
		return deps.Terminal.RunCommand(a.IOContext(), args[1:])
	}
	for {
		if a.AppScope().IsKilled() {
			return
		}
		if err = m.runLoop(a.IOContext(), deps.Terminal); err == nil {
			return nil
		}
		if strictMode {
			a.IOContext().Scope().AppendError(err)
			return err
		}
		io.Err().Printf("ERROR: %v", err)
	}
}

func (m *Module) runLoop(parentCtx app.IOContext, terminal modules.Terminal) (err error) {
	var (
		relatedScope app.Scope
		relatedCtx   app.IOContext
		parentScope  = parentCtx.Scope()
	)
	// Related scope is scope for terminal loop.
	// It is not child scope (It contains separated app.ScopeSync).
	// Kill the scope doesn't kill application scope by default.
	// The scope share data with application scope.
	relatedScope = scope.NewScope(scope.Params{
		DataScope: parentScope,
		//EventScope: parentScope, <- event scope is not shared to prevent memory leaks
	})
	relatedCtx = gio.NewIOContext(relatedScope, parentCtx.IO())
	go func() {
		select {
		case <-parentCtx.Scope().Context().Done():
			// the gorutine kill related context if parent die.
			relatedCtx.Scope().Kill()
		case <-relatedCtx.Scope().Context().Done():
			// stop if related context die (prevent memory leaks)
			return
		}
	}()
	defer relatedCtx.Close()
	if err = terminal.RunLoop(relatedCtx, "\n>"); err != nil {
		return err
	}
	return relatedCtx.Scope().Wait()
}
