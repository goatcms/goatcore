package scope

import (
	"fmt"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/injector"
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
func (ds *DataScope) Get(key string) (value interface{}, err error) {
	var ok bool
	if value, ok = ds.Data[key]; !ok {
		return nil, fmt.Errorf("Unknow value for key %v", key)
	}
	return value, nil
}

// Keys get map data
func (ds *DataScope) Keys() ([]string, error) {
	keys := make([]string, len(ds.Data))
	i := 0
	for key, _ := range ds.Data {
		keys[i] = key
		i++
	}
	return keys, nil
}

// Injector create new injector for the data scope
func (ds *DataScope) Injector(tagname string) app.Injector {
	return injector.NewMapInjector(tagname, ds.Data)
}
