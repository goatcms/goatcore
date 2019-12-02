package sandboxes

import (
	"sync"

	"github.com/goatcms/goatcore/app/modules/pipelinem/services"
	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

// SandboxsManager SandboxsManager is a tool to menage sandboxes.
type SandboxsManager struct {
	factoriesMu sync.RWMutex
	factories   []services.SandboxBuilder
}

// SandboxsManagerFactory create a session manager instance
func SandboxsManagerFactory(dp dependency.Provider) (ins interface{}, error error) {
	return services.SandboxsManager(&SandboxsManager{}), nil
}

// Add new sandbox factory to SandboxsManager
func (manager *SandboxsManager) Add(sandboxBuilder services.SandboxBuilder) {
	manager.factoriesMu.Lock()
	defer manager.factoriesMu.Unlock()
	manager.factories = append(manager.factories, sandboxBuilder)
}

// Get return sandbox by name (create if required)
func (manager *SandboxsManager) Get(name string) (sandbox services.Sandbox, err error) {
	manager.factoriesMu.RLock()
	defer manager.factoriesMu.RUnlock()
	for _, factory := range manager.factories {
		if factory.Is(name) {
			return factory.Build(name)
		}
	}
	return nil, goaterr.Errorf("Can not find sandbox factory for '%s'", name)
}
