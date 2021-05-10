package tasks

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/bootstrap"
	"github.com/goatcms/goatcore/app/goatapp"
	"github.com/goatcms/goatcore/app/modules/commonm"
	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices"
	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices/namespaces"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

func newApp() (mapp app.App, err error) {
	if mapp, err = goatapp.NewMockupApp(goatapp.Params{}); err != nil {
		return nil, err
	}
	dp := mapp.DependencyProvider()
	if err = goaterr.ToError(goaterr.AppendError(nil,
		dp.AddDefaultFactory(pipservices.NamespacesUnitService, namespaces.UnitFactory),
		dp.AddDefaultFactory(pipservices.TasksUnitService, UnitFactory),
	)); err != nil {
		return nil, err
	}
	bootstraper := bootstrap.NewBootstrap(mapp)
	if err = bootstraper.Register(commonm.NewModule()); err != nil {
		return nil, err
	}
	if err = bootstraper.Init(); err != nil {
		return nil, err
	}
	return mapp, nil
}
