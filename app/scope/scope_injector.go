package scope

import (
	"reflect"
	"strings"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

// ScopeInjector is map data injector
type ScopeInjector struct {
	data    app.DataScope
	tagname string
}

// NewScopeInjector create new map injector instance
func NewScopeInjector(tagname string, ds app.DataScope) app.Injector {
	return app.Injector(ScopeInjector{
		tagname: tagname,
		data:    ds,
	})
}

// InjectTo inject data from all injectors
func (ds ScopeInjector) InjectTo(obj interface{}) (err error) {
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
			return goaterr.Errorf("ScopeInjector.InjectTo: %s is not valid", structField.Name)
		}
		if !valueField.CanSet() {
			return goaterr.Errorf("ScopeInjector.InjectTo: Cannot set %s field value", structField.Name)
		}
		if newValue, err = ds.data.Get(key); err != nil {
			return err
		}
		if newValue == nil {
			if !isRequired {
				continue
			}
			return goaterr.Errorf("value for %s is unknown", key)
		}
		if newValue == nil {
			return goaterr.Errorf("ScopeInjector.InjectTo: dependency instance can not be nil (%s)", key)
		}
		refValue := reflect.ValueOf(newValue)
		valueField.Set(refValue)
	}
	return nil
}
