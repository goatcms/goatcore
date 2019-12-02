package pipelinem

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/modules"
	"github.com/goatcms/goatcore/app/modules/pipelinem/services"
	"github.com/goatcms/goatcore/app/modules/pipelinem/services/sandboxes"
	"github.com/goatcms/goatcore/app/modules/pipelinem/services/sandboxes/dockersb"
	"github.com/goatcms/goatcore/app/modules/pipelinem/services/sandboxes/terminalsb"
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
	return goaterr.ToErrors(goaterr.AppendError(nil,
		dp.AddDefaultFactory("SandboxsManager", sandboxes.SandboxsManagerFactory),
		app.RegisterHealthChecker(a, "docker", SandboxHealthChecker),
	))
}

// InitDependencies is init callback to inject dependencies inside module
func (m *Module) InitDependencies(a app.App) (err error) {
	var (
		deps struct {
			Manager  services.SandboxsManager `dependency:"SandboxsManager"`
			Terminal modules.Terminal         `dependency:"TerminalService"`
		}
		builder services.SandboxBuilder
	)
	if err = a.DependencyProvider().InjectTo(&deps); err != nil {
		return err
	}
	if builder, err = terminalsb.NewSandboxBuilder(deps.Terminal); err != nil {
		return err
	}
	deps.Manager.Add(builder)
	deps.Manager.Add(dockersb.NewDockerSandboxBuilder())
	return nil
}

// Run start command line loop
func (m *Module) Run(a app.App) (err error) {
	return nil
}
