package runner

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/bootstrap"
	"github.com/goatcms/goatcore/app/mockupapp"
	"github.com/goatcms/goatcore/app/modules"
	"github.com/goatcms/goatcore/app/modules/commonm"
	"github.com/goatcms/goatcore/app/modules/commonm/commservices"
	"github.com/goatcms/goatcore/app/modules/ocm"
	"github.com/goatcms/goatcore/app/modules/ocm/ocservices"
	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices"
	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices/namespaces"
	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices/sandboxes"
	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices/sandboxes/containersb"
	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices/sandboxes/selfsb"
	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices/tasks"
	"github.com/goatcms/goatcore/app/modules/terminalm"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

func newApp() (mapp app.App, err error) {
	if mapp, err = mockupapp.NewApp(mockupapp.MockupOptions{}); err != nil {
		return nil, err
	}
	dp := mapp.DependencyProvider()
	if err = goaterr.ToError(goaterr.AppendError(nil,
		dp.AddDefaultFactory(pipservices.NamespacesUnitService, namespaces.UnitFactory),
		dp.AddDefaultFactory(pipservices.TasksUnitService, tasks.UnitFactory),
		dp.AddDefaultFactory(pipservices.SandboxesManagerService, sandboxes.ManagerFactory),
		dp.AddDefaultFactory(pipservices.RunnerService, Factory),
	)); err != nil {
		return nil, err
	}
	bootstraper := bootstrap.NewBootstrap(mapp)
	if err = bootstraper.Register(terminalm.NewModule()); err != nil {
		return nil, err
	}
	if err = bootstraper.Register(commonm.NewModule()); err != nil {
		return nil, err
	}
	if err = bootstraper.Register(ocm.NewModule()); err != nil {
		return nil, err
	}
	if err = bootstraper.Init(); err != nil {
		return nil, err
	}
	if err = initDependencies(mapp); err != nil {
		return nil, err
	}
	return mapp, nil
}

func initDependencies(a app.App) (err error) {
	var (
		deps struct {
			Manager          pipservices.SandboxesManager  `dependency:"PipSandboxesManager"`
			Terminal         modules.Terminal              `dependency:"TerminalService"`
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
	return nil
}
