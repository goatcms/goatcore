package args

import (
	"flag"
	"fmt"
	"reflect"
	"strings"

	"github.com/goatcms/goat-core/app"
)

// Injector is map data injector with conversion from string
type Injector struct {
	tagname string
}

// NewInjector create new map string conversion injector instance
func NewInjector(tagname string) app.Injector {
	return app.Injector(Injector{
		tagname: tagname,
	})
}

// InjectTo inject data from all injectors
func (injector Injector) InjectTo(obj interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()
	for i := 0; i < structValue.NumField(); i++ {
		valueField := structValue.Field(i)
		structField := structValue.Type().Field(i)
		key := structField.Tag.Get(injector.tagname)
		if key == "" {
			continue
		}
		if !valueField.IsValid() {
			return fmt.Errorf("goatcore/app/args/injector.InjectTo: %s is not valid", structField.Name)
		}
		if !valueField.CanSet() {
			return fmt.Errorf("goatcore/app/args/injector.InjectTo: Cannot set %s field value", structField.Name)
		}
		newValue, err := decodeRefArg(key, structField.Type.Name())
		if err != nil {
			return err
		}
		refValue := reflect.ValueOf(newValue)
		valueField.Set(refValue)
	}
	return nil
}

func decodeRefArg(name, typeName string) (interface{}, error) {
	switch strings.ToLower(typeName) {
	case "*bool":
		newValue := flag.Bool(name, app.DefaultBoolValue, name)
		return newValue, nil
	case "bool":
		newValue := flag.Bool(name, app.DefaultBoolValue, name)
		return *newValue, nil

	case "*string":
		newValue := flag.String(name, app.DefaultStringValue, name)
		return *newValue, nil
	case "string":
		newValue := flag.String(name, app.DefaultStringValue, name)
		return *newValue, nil

	case "*time.duration":
		newValue := flag.Duration(name, app.DefaultDurationValue, name)
		return *newValue, nil
	case "time.duration":
		newValue := flag.Duration(name, app.DefaultDurationValue, name)
		return *newValue, nil

	case "*float64":
		newValue := flag.Float64(name, app.DefaultFloat64Value, name)
		return *newValue, nil
	case "float64":
		newValue := flag.Float64(name, app.DefaultFloat64Value, name)
		return *newValue, nil

	case "*int":
		newValue := flag.Int(name, app.DefaultIntValue, name)
		return *newValue, nil
	case "int":
		newValue := flag.Int(name, app.DefaultIntValue, name)
		return *newValue, nil

	case "*int64":
		newValue := flag.Int64(name, app.DefaultInt64Value, name)
		return *newValue, nil
	case "int64":
		newValue := flag.Int64(name, app.DefaultInt64Value, name)
		return *newValue, nil

	case "*uint":
		newValue := flag.Uint(name, app.DefaultUIntValue, name)
		return *newValue, nil
	case "uint":
		newValue := flag.Uint(name, app.DefaultUIntValue, name)
		return *newValue, nil

	case "*uint64":
		newValue := flag.Uint64(name, app.DefaultUInt64Value, name)
		return *newValue, nil
	case "uint64":
		newValue := flag.Uint64(name, app.DefaultUInt64Value, name)
		return *newValue, nil
	}
	return nil, fmt.Errorf("unknow type %s for argument %s", typeName, name)
}
