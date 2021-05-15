package datascope

import (
	"sync"

	"github.com/goatcms/goatcore/app"
)

// DataScope represent scope data
type DataScope struct {
	Data map[interface{}]interface{}
	mu   sync.RWMutex
}

// New create new instance of data scope
func New(data map[interface{}]interface{}) app.DataScope {
	return app.DataScope(&DataScope{
		Data: data,
	})
}

// Set new scope value
func (ds *DataScope) SetValue(key interface{}, v interface{}) {
	ds.mu.Lock()
	ds.Data[key] = v
	ds.mu.Unlock()
}

// Get get value from context
func (ds *DataScope) Value(key interface{}) (value interface{}) {
	ds.mu.RLock()
	value = ds.Data[key]
	ds.mu.RUnlock()
	return
}

// Keys get map data
func (ds *DataScope) Keys() (keys []interface{}) {
	ds.mu.RLock()
	keys = make([]interface{}, len(ds.Data))
	i := 0
	for key := range ds.Data {
		keys[i] = key
		i++
	}
	ds.mu.RUnlock()
	return
}

// LockData return new data locker
func (ds *DataScope) LockData() (locker app.DataScopeLocker) {
	ds.mu.Lock()
	return newDataLocker(ds.Data, ds.mu.Unlock, nil)
}
