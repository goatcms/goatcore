package scope

import (
	"sync"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/injector"
)

// DataChildScope represent child data scope. It contains all parent data and its own.
type DataChildScope struct {
	parent app.DataScope
	data   map[string]interface{}
	mu     sync.RWMutex
}

// NewChildDataScope create new instance of child data scope
func NewChildDataScope(parent app.DataScope, data map[string]interface{}) app.DataScope {
	return app.DataScope(&DataChildScope{
		parent: parent,
		data:   data,
	})
}

// Set new scope value
func (scp *DataChildScope) Set(key string, v interface{}) error {
	scp.mu.Lock()
	defer scp.mu.Unlock()
	scp.data[key] = v
	return nil
}

// Get get value from context
func (scp *DataChildScope) Get(key string) (value interface{}, err error) {
	var (
		ok bool
	)
	scp.mu.RLock()
	if value, ok = scp.data[key]; ok {
		scp.mu.RUnlock()
		return value, nil
	}
	scp.mu.RUnlock()
	return scp.parent.Get(key)
}

// Keys get map data
func (scp *DataChildScope) Keys() ([]string, error) {
	scp.mu.RLock()
	defer scp.mu.RUnlock()
	keys := make([]string, len(scp.data))
	i := 0
	for key := range scp.data {
		keys[i] = key
		i++
	}
	return keys, nil
}

// Injector create new injector for the data scope
func (scp *DataChildScope) Injector(tagname string) app.Injector {
	return injector.NewMapInjector(tagname, scp.data)
}

// LockData return new data locker
func (scp *DataChildScope) LockData() (locker app.DataScopeLocker) {
	scp.mu.Lock()
	return newDataLocker(scp.data, scp.mu.Unlock, scp.parent)
}
