package terminalm

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/modules/terminalm/termcommands/commonc"
	"github.com/goatcms/goatcore/app/modules/terminalm/termcommands/helpc"
	"github.com/goatcms/goatcore/app/modules/terminalm/termcommands/termc"
	"github.com/goatcms/goatcore/app/modules/terminalm/termservices"
	"github.com/goatcms/goatcore/app/modules/terminalm/termservices/terminals"
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
	a.Terminal().SetArgument(commonc.Arguments()...)
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
			Terminal termservices.Terminal `dependency:"TerminalService"`
		}
		exTerm termservices.Terminal
	)
	if err = a.DependencyProvider().InjectTo(&deps); err != nil {
		return err
	}
	exTerm = deps.Terminal.Extends(append(
		termc.Commands(),
		helpc.Commands()...,
	)...)
	args := a.Arguments()
	if len(args) < 2 {
		return exTerm.RunCommand(a.IOContext(), []string{"help"})
	}
	return exTerm.RunCommand(a.IOContext(), args[1:])
}
