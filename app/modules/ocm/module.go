package ocm

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/modules/ocm/ocservices"
	"github.com/goatcms/goatcore/app/modules/ocm/ocservices/dcmd"
	"github.com/goatcms/goatcore/app/modules/ocm/ocservices/ocmanager"
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
		dp.AddDefaultFactory(ocservices.OCManagerService, ocmanager.ManagerFactory),
	))
}

// InitDependencies is init callback to inject dependencies inside module
func (m *Module) InitDependencies(a app.App) (err error) {
	var (
		deps struct {
			Manager       ocservices.Manager `dependency:"OCManager"`
			DefaultEngine string             `argument:"?oc.engime"`
		}
		dockerEngine ocservices.Engine
		podmanEngine ocservices.Engine
	)
	if err = a.DependencyProvider().InjectTo(&deps); err != nil {
		return err
	}
	if hasDocker {
		dockerEngine = dcmd.NewEngine("docker")
		if err = deps.Manager.AddEngine(ocservices.DockerEngine, dockerEngine); err != nil {
			return err
		}
		deps.Manager.SetDefaultEngine(dockerEngine)
	}
	if hasPodman {
		podmanEngine = dcmd.NewEngine("podman")
		if err = deps.Manager.AddEngine(ocservices.DockerEngine, podmanEngine); err != nil {
			return err
		}
		deps.Manager.SetDefaultEngine(podmanEngine)
	}
	if deps.DefaultEngine != "" {
		if deps.DefaultEngine == "docker" {
			if dockerEngine == nil {
				return goaterr.Errorf("Docker open container engine is unavailable on your device")
			}
			deps.Manager.SetDefaultEngine(dockerEngine)
		} else if deps.DefaultEngine == "podman" {
			if podmanEngine == nil {
				return goaterr.Errorf("Podman open container engine is unavailable on your device")
			}
			deps.Manager.SetDefaultEngine(podmanEngine)
		} else {
			return goaterr.Errorf("Unknow '%s' open container engine", deps.DefaultEngine)
		}
	}
	return nil
}

// Run is empty function
func (m *Module) Run(a app.App) (err error) {
	return nil
}
