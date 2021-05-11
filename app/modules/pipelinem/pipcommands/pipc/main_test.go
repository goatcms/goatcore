package pipc

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
	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices/sandboxes/containersb"
	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices/sandboxes/selfsb"
	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices/tasks"
	"github.com/goatcms/goatcore/app/modules/terminalm"
	"github.com/goatcms/goatcore/app/modules/terminalm/termservices"
	"github.com/goatcms/goatcore/app/terminal"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

func newApp(params goatapp.Params) (mapp *goatapp.MockupApp, bootstraper app.Bootstrap, err error) {
	if mapp, err = goatapp.NewMockupApp(params); err != nil {
		return nil, nil, err
	}
	dp := mapp.DependencyProvider()
	if err = goaterr.ToError(goaterr.AppendError(nil,
		dp.AddDefaultFactory(pipservices.NamespacesUnitService, namespaces.UnitFactory),
		dp.AddDefaultFactory(pipservices.TasksUnitService, tasks.UnitFactory),
		dp.AddDefaultFactory(pipservices.SandboxesManagerService, sandboxes.ManagerFactory),
		dp.AddDefaultFactory(pipservices.RunnerService, runner.Factory),
	)); err != nil {
		return nil, nil, err
	}
	term := mapp.Terminal()
	term.SetCommand(
		RunCommand(),
		TryCommand(),
	)
	term.SetCommand(terminal.NewCommand(terminal.CommandParams{
		Name: "testCommand",
		Callback: func(a app.App, ctx app.IOContext) (err error) {
			return ctx.IO().Out().Printf("output")
		},
		Help: "print 'output'",
	}))
	bootstraper = bootstrap.NewBootstrap(mapp)
	if err = goaterr.ToError(goaterr.AppendError(nil,
		bootstraper.Register(terminalm.NewModule()),
		bootstraper.Register(commonm.NewModule()),
		bootstraper.Register(ocm.NewModule()),
		bootstraper.Init(),
		initDependencies(mapp),
	)); err != nil {
		return nil, nil, err
	}
	return mapp, bootstraper, nil
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
	deps.Manager.Add(containersb.NewContainerSandboxBuilder(deps.EnvironmentsUnit, deps.OCManager))
	return nil
}
