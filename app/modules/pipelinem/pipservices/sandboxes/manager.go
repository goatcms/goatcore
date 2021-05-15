package sandboxes

import (
	"sync"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

// Manager SandboxesManager is a tool to menage sandboxes.
type Manager struct {
	factoriesMu sync.RWMutex
	factories   []pipservices.SandboxBuilder
}

// ManagerFactory create a session manager instance
func ManagerFactory(dp app.DependencyProvider) (ins interface{}, error error) {
	return pipservices.SandboxesManager(&Manager{}), nil
}

// Add new sandbox factory to SandboxesManager
func (manager *Manager) Add(sandboxBuilder pipservices.SandboxBuilder) {
	manager.factoriesMu.Lock()
	defer manager.factoriesMu.Unlock()
	manager.factories = append(manager.factories, sandboxBuilder)
}

// Get return sandbox by name (create if required)
func (manager *Manager) Get(name string) (sandbox pipservices.Sandbox, err error) {
	manager.factoriesMu.RLock()
	defer manager.factoriesMu.RUnlock()
	for _, factory := range manager.factories {
		if factory.Is(name) {
			return factory.Build(name)
		}
	}
	return nil, goaterr.Errorf("Can not find sandbox factory for '%s'", name)
}
