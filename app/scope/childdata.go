package scope

import (
	"sync"

	"github.com/goatcms/goatcore/app"
)

// DataChildScope represent child data scope. It contains all parent data and its own.
type DataChildScope struct {
	parent app.DataScope
	data   map[interface{}]interface{}
	mu     sync.RWMutex
}

// NewChildDataScope create new instance of child data scope
func NewChildDataScope(parent app.DataScope, data map[interface{}]interface{}) app.DataScope {
	return app.DataScope(&DataChildScope{
		parent: parent,
		data:   data,
	})
}

// Set new scope value
func (scp *DataChildScope) SetValue(key interface{}, v interface{}) {
	scp.mu.Lock()
	scp.data[key] = v
	scp.mu.Unlock()
}

// Get get value from context
func (scp *DataChildScope) Value(key interface{}) (value interface{}) {
	var (
		ok bool
	)
	scp.mu.RLock()
	if value, ok = scp.data[key]; ok {
		scp.mu.RUnlock()
		return value
	}
	scp.mu.RUnlock()
	return scp.parent.Value(key)
}

// Keys get map data
func (scp *DataChildScope) Keys() (keys []interface{}) {
	scp.mu.RLock()
	keys = make([]interface{}, len(scp.data))
	i := 0
	for key := range scp.data {
		keys[i] = key
		i++
	}
	scp.mu.RUnlock()
	return keys
}

// LockData return new data locker
func (scp *DataChildScope) LockData() (locker app.DataScopeLocker) {
	scp.mu.Lock()
	return newDataLocker(scp.data, scp.mu.Unlock, scp.parent)
}
