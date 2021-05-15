package datascope

import (
	"reflect"
	"strings"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

// Injector is map data injector
type Injector struct {
	data    app.DataScope
	tagname string
}

// NewInjector create new map injector instance
func NewInjector(tagname string, ds app.DataScope) app.Injector {
	return Injector{
		tagname: tagname,
		data:    ds,
	}
}

// InjectTo inject data from all injectors
func (ds Injector) InjectTo(obj interface{}) (err error) {
	var newValue interface{}
	structValue := reflect.ValueOf(obj).Elem()
	for i := 0; i < structValue.NumField(); i++ {
		var isRequired = true
		valueField := structValue.Field(i)
		structField := structValue.Type().Field(i)
		key := structField.Tag.Get(ds.tagname)
		if key == "" {
			continue
		}
		if strings.HasPrefix(key, "?") {
			isRequired = false
			key = key[1:]
		}
		if !valueField.IsValid() {
			return goaterr.Errorf("Injector.InjectTo: %s is not valid", structField.Name)
		}
		if !valueField.CanSet() {
			return goaterr.Errorf("Injector.InjectTo: Cannot set %s field value", structField.Name)
		}
		newValue = ds.data.Value(key)
		if newValue == nil {
			if !isRequired {
				continue
			}
			return goaterr.Errorf("value for %s is unknown", key)
		}
		if newValue == nil {
			return goaterr.Errorf("Injector.InjectTo: dependency instance can not be nil (%s)", key)
		}
		refValue := reflect.ValueOf(newValue)
		valueField.Set(refValue)
	}
	return nil
}
