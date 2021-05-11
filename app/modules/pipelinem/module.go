package pipelinem

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/modules/commonm/commservices"
	"github.com/goatcms/goatcore/app/modules/ocm/ocservices"
	"github.com/goatcms/goatcore/app/modules/pipelinem/pipcommands/pipc"
	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices"
	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices/namespaces"
	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices/runner"
	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices/sandboxes"
	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices/sandboxes/containersb"
	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices/sandboxes/selfsb"
	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices/sandboxes/sshsb"
	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices/tasks"
	"github.com/goatcms/goatcore/app/modules/terminalm/termservices"
	"github.com/goatcms/goatcore/varutil/goaterr"
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
	term := a.Terminal()
	term.SetCommand(pipc.Commands()...)
	return goaterr.ToError(goaterr.AppendError(nil,
		dp.AddDefaultFactory(pipservices.SandboxesManagerService, sandboxes.ManagerFactory),
		dp.AddDefaultFactory(pipservices.NamespacesUnitService, namespaces.UnitFactory),
		dp.AddDefaultFactory(pipservices.RunnerService, runner.Factory),
		dp.AddDefaultFactory(pipservices.TasksUnitService, tasks.UnitFactory),
	))
}

// InitDependencies is init callback to inject dependencies inside module
func (m *Module) InitDependencies(a app.App) (err error) {
	var (
		deps struct {
			Manager          pipservices.SandboxesManager  `dependency:"PipSandboxesManager"`
			Terminal         termservices.Terminal         `dependency:"TerminalService"`
			EnvironmentsUnit commservices.EnvironmentsUnit `dependency:"CommonEnvironmentsUnit"`
			OCManager        ocservices.Manager            `dependency:"OCManager"`
		}
		builder pipservices.SandboxBuilder
	)
	if err = a.DependencyProvider().InjectTo(&deps); err != nil {
		return err
	}
	if builder, err = selfsb.NewSandboxBuilder(deps.Terminal); err != nil {
		return err
	}
	deps.Manager.Add(builder)
	deps.Manager.Add(containersb.NewContainerSandboxBuilder(deps.EnvironmentsUnit, deps.OCManager))
	deps.Manager.Add(sshsb.NewSSHSandboxBuilder(deps.EnvironmentsUnit))
	return nil
}

// Run start command line loop
func (m *Module) Run(a app.App) (err error) {
	return nil
}
