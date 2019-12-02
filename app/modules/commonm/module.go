package commonm

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/modules/commonm/services/namespaces"
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
		dp.AddDefaultFactory("Namespaces", namespaces.NamasepacesFactory),
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
