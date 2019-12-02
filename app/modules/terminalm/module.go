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
		}
		io         app.IO
		strictMode bool
	)
	if err = a.DependencyProvider().InjectTo(&deps); err != nil {
		return err
	}
	strictMode = strings.ToLower(deps.StrictMode) == "true"
	args := os.Args[1:]
	if len(args) == 0 {
		return deps.Terminal.RunCommand(a.IOContext(), []string{"help"})
	} else if args[0] != "terminal" {
		return deps.Terminal.RunCommand(a.IOContext(), args)
	}
	for {
		if err = deps.Terminal.RunLoop(a.IOContext()); err == nil {
			return nil
		}
		if strictMode {
			return err
		}
		io.Err().Printf("ERROR: %v", err)
	}
}
