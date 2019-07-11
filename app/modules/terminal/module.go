package terminal

import (
	"os"
	"strings"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/modules"
)

// Module is command unit
type Module struct{}

// NewModule create new command module instance
func NewModule() app.Module {
	return &Module{}
}

// RegisterDependencies is init callback to register module dependencies
func (m *Module) RegisterDependencies(a app.App) error {
	commandScope := a.CommandScope()
	commandScope.Set("command.h", app.CommandCallback(HelpComamnd))
	commandScope.Set("command.help", app.CommandCallback(HelpComamnd))
	a.DependencyProvider().AddDefaultFactory(modules.TerminalService, IOTerminalFactory)
	return nil
}

// InitDependencies is init callback to inject dependencies inside module
func (m *Module) InitDependencies(a app.App) error {
	return nil
}

// Run start command line loop
func (m *Module) Run(a app.App) (err error) {
	var deps struct {
		Output     app.Output       `dependency:"OutputService"`
		Terminal   modules.Terminal `dependency:"TerminalService"`
		StrictMode string           `argument:"?strict"`
	}
	if err = a.DependencyProvider().InjectTo(&deps); err != nil {
		return err
	}
	deps.StrictMode = strings.ToLower(deps.StrictMode)
	args := os.Args[1:]
	if len(args) != 0 && args[0] == "terminal" {
		return m.runLoop(deps.Output, deps.Terminal, strings.ToLower(deps.StrictMode) == "true")
	}
	if err = deps.Terminal.RunCommand(args); err != nil {
		return err
	}
	return app.CloseApp(a)
}

func (m *Module) runLoop(out app.Output, terminal modules.Terminal, strict bool) (err error) {
	for {
		err = terminal.RunLoop()
		if err == nil {
			return nil
		}
		if strict {
			return err
		}
		out.Printf("CRITIC ERROR: %v", err)
	}
}
