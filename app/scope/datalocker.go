package scope

import (
	"sync"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/injector"
)

type dataLockerUnlocker func()

// DataLocker represent scope data
type DataLocker struct {
	data     map[string]interface{}
	mu       sync.RWMutex
	unlockCB dataLockerUnlocker
}

// Commit unlock parent scope and close locker
func newDataLocker(data map[string]interface{}, unlockCB dataLockerUnlocker) (locker app.DataScopeLocker) {
	return &DataLocker{
		data:     data,
		unlockCB: unlockCB,
	}
}

// Set new scope value
func (locker *DataLocker) Set(key string, v interface{}) error {
	locker.mu.Lock()
	defer locker.mu.Unlock()
	locker.data[key] = v
	return nil
}

// Get get value from context
func (locker *DataLocker) Get(key string) (value interface{}, err error) {
	locker.mu.RLock()
	defer locker.mu.RUnlock()
	return locker.data[key], nil
}

// Keys get map data
func (locker *DataLocker) Keys() ([]string, error) {
	locker.mu.RLock()
	defer locker.mu.RUnlock()
	keys := make([]string, len(locker.data))
	i := 0
	for key := range locker.data {
		keys[i] = key
		i++
	}
	return keys, nil
}

// Injector create new injector for the data scope
func (locker *DataLocker) Injector(tagname string) app.Injector {
	return injector.NewMapInjector(tagname, locker.data)
}

// LockData return new data locker
func (locker *DataLocker) LockData() app.DataScopeLocker {
	locker.mu.Lock()
	return newDataLocker(locker.data, locker.mu.Unlock)
}

// Commit unlock parent scope and close locker
func (locker *DataLocker) Commit() (err error) {
	locker.unlockCB()
	locker.data = nil
	return nil
}
