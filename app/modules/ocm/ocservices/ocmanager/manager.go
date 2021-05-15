package ocmanager

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/modules/ocm/ocservices"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

// Manager provide OC engines
type Manager struct {
	engines       map[string]ocservices.Engine
	defaultEngine ocservices.Engine
}

// ManagerFactory create an environment variables manager instance
func ManagerFactory(dp app.DependencyProvider) (ins interface{}, error error) {
	return ocservices.Manager(&Manager{
		engines: make(map[string]ocservices.Engine),
	}), nil
}

// AddEngine register new engine
func (manager *Manager) AddEngine(id string, engine ocservices.Engine) (err error) {
	if _, ok := manager.engines[id]; ok {
		return goaterr.Errorf("Engine %s already exists", id)
	}
	manager.engines[id] = engine
	return nil
}

// Run container
func (manager *Manager) Run(container ocservices.Container) (err error) {
	var (
		engine ocservices.Engine
		ok     bool
	)
	if container.Engine == "" {
		return manager.defaultEngine.Run(container)
	}
	if engine, ok = manager.engines[container.Engine]; !ok {
		return goaterr.Errorf("Unknow engine %s", container.Engine)
	}
	return engine.Run(container)
}

// SetDefaultEngine set default engine
func (manager *Manager) SetDefaultEngine(engine ocservices.Engine) {
	manager.defaultEngine = engine
}
