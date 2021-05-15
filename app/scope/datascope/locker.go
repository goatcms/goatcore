package datascope

import (
	"sync"

	"github.com/goatcms/goatcore/app"
)

type dataLockerUnlocker func()

// DataLocker represent the data scope transaction
type DataLocker struct {
	data     map[interface{}]interface{}
	mu       sync.RWMutex
	unlockCB dataLockerUnlocker
	parent   app.DataScope
}

// Commit unlock parent scope and close locker
func newDataLocker(data map[interface{}]interface{}, unlockCB dataLockerUnlocker, parent app.DataScope) (locker app.DataScopeLocker) {
	return &DataLocker{
		parent:   parent,
		data:     data,
		unlockCB: unlockCB,
	}
}

func (locker *DataLocker) SetValue(key interface{}, v interface{}) {
	locker.mu.Lock()
	locker.data[key] = v
	locker.mu.Unlock()
}

func (locker *DataLocker) Value(key interface{}) (value interface{}) {
	var ok bool
	locker.mu.RLock()
	value, ok = locker.data[key]
	locker.mu.RUnlock()
	if ok {
		return value
	}
	if locker.parent != nil {
		return locker.parent.Value(key)
	}
	return nil
}

// Keys get map data
func (locker *DataLocker) Keys() (keys []interface{}) {
	locker.mu.RLock()
	defer locker.mu.RUnlock()
	keys = make([]interface{}, len(locker.data))
	i := 0
	for key := range locker.data {
		keys[i] = key
		i++
	}
	return
}

// LockData return new data locker
func (locker *DataLocker) LockData() app.DataScopeLocker {
	locker.mu.Lock()
	return newDataLocker(locker.data, locker.mu.Unlock, locker.parent)
}

// Commit unlock parent scope and close locker
func (locker *DataLocker) Commit() (err error) {
	locker.unlockCB()
	locker.data = nil
	return nil
}
