package mutex

import (
	"sort"
	"sync"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/modules/commonm/commservices"
)

// SharedMutex lock reseources
type SharedMutex struct {
	mutexesMU sync.RWMutex
	mutexes   map[string]*sync.RWMutex
}

// NewSharedMutex create new SharedMutex instance
func NewSharedMutex() *SharedMutex {
	return &SharedMutex{
		mutexes: map[string]*sync.RWMutex{},
	}
}

// SharedMutexFactory create a SharedMutex instance
func SharedMutexFactory(dp app.DependencyProvider) (ri interface{}, err error) {
	return commservices.SharedMutex(NewSharedMutex()), nil
}

// Lock resources
// You must use the function once for a separeted task. You must ulock the resources after finish.
// A lock order is very important. If you make sure that all locks are always
// taken in the same order by any thread, deadlocks cannot occur.
// more at "Deadlock solution 1: lock ordering": https://web.mit.edu/6.005/www/fa15/classes/23-locks/
// The function provide lock ordering inside.
func (sharedMutex *SharedMutex) Lock(resources commservices.LockMap) (handler commservices.UnlockHandler) {
	var (
		list = make([]mutexRow, len(resources))
		i    = 0
	)
	for name, value := range resources {
		list[i].Name = name
		list[i].Value = value
		i++
	}
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].Name < list[j].Name
	})
	for _, row := range list {
		mu := sharedMutex.get(row.Name)
		if row.Value == commservices.LockR {
			mu.RLock()
		} else {
			mu.Lock()
		}
	}
	return &unlockHandler{
		list:        list,
		sharedMutex: sharedMutex,
	}
}

// get return mutex for name (create if mutex does not exist)
func (sharedMutex *SharedMutex) get(name string) (mu *sync.RWMutex) {
	var ok bool
	sharedMutex.mutexesMU.RLock()
	if mu, ok = sharedMutex.mutexes[name]; !ok {
		sharedMutex.mutexesMU.RUnlock()
		sharedMutex.mutexesMU.Lock()
		defer sharedMutex.mutexesMU.Unlock()
		if mu, ok = sharedMutex.mutexes[name]; ok {
			return mu
		}
		mu = &sync.RWMutex{}
		sharedMutex.mutexes[name] = mu
		return mu
	}
	sharedMutex.mutexesMU.RUnlock()
	return mu
}
