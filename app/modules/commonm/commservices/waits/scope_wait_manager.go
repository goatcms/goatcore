package waits

import (
	"sync"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/modules/commonm/commservices"
)

// ScopeWaitManager is wait group manager
type ScopeWaitManager struct {
	waitGroupsMU sync.Mutex
	waitGroups   map[string]*sync.WaitGroup
	ctxScope     app.Scope
}

// NewScopeWaitManager create new ScopeWaitManager instance
func NewScopeWaitManager(ctxScope app.Scope) commservices.ScopeWaitManager {
	return &ScopeWaitManager{
		waitGroups: map[string]*sync.WaitGroup{},
		ctxScope:   ctxScope,
	}
}

// Add task counter to waitgroup
func (manager *ScopeWaitManager) Add(name string, delta int) {
	wg := manager.get(name)
	wg.Add(delta)
	manager.ctxScope.AddTasks(delta)
}

// Done decrements the WaitGroup counter by one.
func (manager *ScopeWaitManager) Done(name string) {
	manager.get(name).Done()
	manager.ctxScope.DoneTask()
}

// Wait blocks until the WaitGroup counter is zero.
func (manager *ScopeWaitManager) Wait(name string) {
	manager.get(name).Wait()
}

func (manager *ScopeWaitManager) get(name string) (wg *sync.WaitGroup) {
	manager.waitGroupsMU.Lock()
	defer manager.waitGroupsMU.Unlock()
	if wg = manager.waitGroups[name]; wg == nil {
		wg = &sync.WaitGroup{}
		manager.waitGroups[name] = wg
	}
	return wg
}
