package scope

import (
	"github.com/goatcms/goat-core/app"
	"github.com/goatcms/goat-core/app/injector"
)

// DataScope represent scope data
type DataScope struct {
	Data map[string]interface{}
}

// NewDataScope create new instance of data scope
func NewDataScope(data map[string]interface{}) app.DataScope {
	return app.DataScope(&DataScope{
		Data: make(map[string]interface{}),
	})
}

// Set new scope value
func (ds *DataScope) Set(key string, v interface{}) error {
	ds.Data[key] = v
	return nil
}

// Get get value from context
func (ds *DataScope) Get(key string) interface{} {
	return ds.Data[key]
}

// Keys get map data
func (ds *DataScope) Keys() []string {
	keys := make([]string, 0, len(ds.Data))
	for k := range ds.Data {
		keys = append(keys, k)
	}
	return keys
}

// Injector create new injector for the data scope
func (ds *DataScope) Injector(tagname string) app.Injector {
	return injector.NewMapInjector(tagname, ds.Data)
}
