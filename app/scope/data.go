package scope

import (
	"sync"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/injector"
)

// DataScope represent scope data
type DataScope struct {
	Data map[string]interface{}
	mu   sync.RWMutex
}

// NewDataScope create new instance of data scope
func NewDataScope(data map[string]interface{}) app.DataScope {
	return app.DataScope(&DataScope{
		Data: data,
	})
}

// Set new scope value
func (ds *DataScope) Set(key string, v interface{}) error {
	ds.mu.Lock()
	defer ds.mu.Unlock()
	ds.Data[key] = v
	return nil
}

// Get get value from context
func (ds *DataScope) Get(key string) (value interface{}, err error) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()
	return ds.Data[key], nil
}

// Keys get map data
func (ds *DataScope) Keys() ([]string, error) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()
	keys := make([]string, len(ds.Data))
	i := 0
	for key := range ds.Data {
		keys[i] = key
		i++
	}
	return keys, nil
}

// Injector create new injector for the data scope
func (ds *DataScope) Injector(tagname string) app.Injector {
	return injector.NewMapInjector(tagname, ds.Data)
}

// LockData return new data locker
func (ds *DataScope) LockData() (locker app.DataScopeLocker) {
	ds.mu.Lock()
	return newDataLocker(ds.Data, ds.mu.Unlock, nil)
}
