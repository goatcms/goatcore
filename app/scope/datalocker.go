package scope

import (
	"sync"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/injector"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

type dataLockerUnlocker func()

// DataLocker represent scope data
type DataLocker struct {
	*DataScope
	mu       sync.RWMutex
	unlockCB dataLockerUnlocker
}

// Commit unlock parent scope and close locker
func newDataLocker(dataScope *DataScope, unlockCB dataLockerUnlocker) (locker app.DataScopeLocker) {
	return &DataLocker{
		DataScope: dataScope,
		unlockCB:  unlockCB,
	}
}

// Set new scope value
func (locker *DataLocker) Set(key string, v interface{}) error {
	locker.mu.Lock()
	defer locker.mu.Unlock()
	locker.DataScope.Data[key] = v
	return nil
}

// Get get value from context
func (locker *DataLocker) Get(key string) (value interface{}, err error) {
	locker.mu.RLock()
	defer locker.mu.RUnlock()
	var ok bool
	if value, ok = locker.DataScope.Data[key]; !ok {
		return nil, goaterr.Errorf("Unknow value for key %v", key)
	}
	return value, nil
}

// Keys get map data
func (locker *DataLocker) Keys() ([]string, error) {
	locker.mu.RLock()
	defer locker.mu.RUnlock()
	keys := make([]string, len(locker.DataScope.Data))
	i := 0
	for key := range locker.DataScope.Data {
		keys[i] = key
		i++
	}
	return keys, nil
}

// Injector create new injector for the data scope
func (locker *DataLocker) Injector(tagname string) app.Injector {
	return injector.NewMapInjector(tagname, locker.DataScope.Data)
}

// LockData return new data locker
func (locker *DataLocker) LockData() app.DataScopeLocker {
	locker.mu.Lock()
	return newDataLocker(locker.DataScope, locker.mu.Unlock)
}

// Commit unlock parent scope and close locker
func (locker *DataLocker) Commit() (err error) {
	locker.unlockCB()
	locker.DataScope = nil
	return nil
}
