package workers

import (
	"sort"
	"sync"
)

// ResourcesManager is a resources allocation system
type ResourcesManager struct {
	rowsMU sync.Mutex
	rows   map[string]*sync.Mutex
}

// NewResourcesManager create a new resources object
func NewResourcesManager() *ResourcesManager {
	return &ResourcesManager{
		rows: map[string]*sync.Mutex{},
	}
}

// Run task by name
func (r *ResourcesManager) row(resourceName string) (mu *sync.Mutex) {
	var ok bool
	r.rowsMU.Lock()
	defer r.rowsMU.Unlock()
	if mu, ok = r.rows[resourceName]; !ok {
		mu = &sync.Mutex{}
		r.rows[resourceName] = mu
	}
	return mu
}

// Lock lock a resource
func (r *ResourcesManager) Lock(resourceName string) {
	r.row(resourceName).Lock()
}

// LockAll lock all resource
func (r *ResourcesManager) LockAll(resources []string) {
	// this sort is important. It resolve deadlocks problems
	// Always lock A -> B -> C
	// Something lockked sequence runed by other thread
	// like B -> C -> A can create deadlocks
	// but if we allocate allways in the same order. This problem doesn't exist.
	sort.Strings(resources)
	for _, row := range resources {
		r.row(row).Lock()
	}
}

// Unlock resource
func (r *ResourcesManager) Unlock(resourceName string) {
	r.row(resourceName).Unlock()
}

// UnlockAll unlock all resource
func (r *ResourcesManager) UnlockAll(resources []string) {
	for _, row := range resources {
		r.row(row).Unlock()
	}
}
