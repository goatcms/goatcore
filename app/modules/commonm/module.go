package commonm

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/modules/commonm/commservices"
	"github.com/goatcms/goatcore/app/modules/commonm/commservices/args"
	"github.com/goatcms/goatcore/app/modules/commonm/commservices/envs"
	"github.com/goatcms/goatcore/app/modules/commonm/commservices/infos"
	"github.com/goatcms/goatcore/app/modules/commonm/commservices/mutex"
	"github.com/goatcms/goatcore/app/modules/commonm/commservices/waits"
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
	return goaterr.ToError(goaterr.AppendError(nil,
		dp.AddDefaultFactory(commservices.ArgumentsServcies, args.ArgumentsFactory),
		dp.AddDefaultFactory(commservices.EnvironmentsUnitService, envs.UnitFactory),
		dp.AddDefaultFactory(commservices.InfoService, infos.InfoFactory),
		dp.AddDefaultFactory(commservices.SharedMutexService, mutex.SharedMutexFactory),
		dp.AddDefaultFactory(commservices.WaitManagerService, waits.WaitManagerFactory),
	))
}

// InitDependencies is init callback to inject dependencies inside module
func (m *Module) InitDependencies(a app.App) (err error) {
	return nil
}

// Run start command line loop
func (m *Module) Run(a app.App) (err error) {
	return nil
}
