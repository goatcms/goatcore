package containersb

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/bootstrap"
	"github.com/goatcms/goatcore/app/goatapp"
	"github.com/goatcms/goatcore/app/modules/commonm"
	"github.com/goatcms/goatcore/app/modules/commonm/commservices"
	"github.com/goatcms/goatcore/app/modules/ocm"
	"github.com/goatcms/goatcore/app/modules/ocm/ocservices"
	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices"
	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices/namespaces"
	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices/runner"
	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices/sandboxes"
	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices/sandboxes/selfsb"
	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices/tasks"
	"github.com/goatcms/goatcore/app/modules/terminalm"
	"github.com/goatcms/goatcore/app/modules/terminalm/termservices"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

func newApp() (mapp app.App, err error) {
	if mapp, err = goatapp.NewMockupApp(goatapp.Params{}); err != nil {
		return nil, err
	}
	dp := mapp.DependencyProvider()
	if err = goaterr.ToError(goaterr.AppendError(nil,
		dp.AddDefaultFactory(pipservices.SandboxesManagerService, sandboxes.ManagerFactory),
		dp.AddDefaultFactory(pipservices.NamespacesUnitService, namespaces.UnitFactory),
		dp.AddDefaultFactory(pipservices.TasksUnitService, tasks.UnitFactory),
		dp.AddDefaultFactory(pipservices.RunnerService, runner.Factory),
	)); err != nil {
		return nil, err
	}
	bootstraper := bootstrap.NewBootstrap(mapp)
	return mapp, goaterr.ToError(goaterr.AppendError(nil,
		bootstraper.Register(terminalm.NewModule()),
		bootstraper.Register(commonm.NewModule()),
		bootstraper.Register(ocm.NewModule()),
		bootstraper.Init(),
		initDependencies(mapp),
	))
}

func initDependencies(a app.App) (err error) {
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
	deps.Manager.Add(NewContainerSandboxBuilder(deps.EnvironmentsUnit, deps.OCManager))
	return nil
}
