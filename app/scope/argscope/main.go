package argscope

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/scope"
	"github.com/goatcms/goatcore/varutil/r"
)

// ArgScope is application arguments scope
type ArgScope struct {
	app.EventScope
	data    map[string]string
	tagname string
}

// NewScope create new ArgScope instance
func NewScope(args []string, tagname string) (app.Scope, error) {
	scope := &ArgScope{
		EventScope: scope.NewEventScope(),
		tagname:    tagname,
		data:       map[string]string{},
	}
	for i, value := range args {
		// position keys
		ikey := "$" + strconv.Itoa(i)
		scope.data[ikey] = value
		// reduce prefixes
		if strings.HasPrefix(value, "--") {
			value = value[2:]
		} else if strings.HasPrefix(value, "-") {
			value = value[1:]
		}
		// key:value
		index := strings.Index(value, "=")
		if index != -1 {
			key := value[:index]
			value = value[index+1:]
			scope.data[key] = value
		} else {
			scope.data[value] = "true"
		}
	}
	return scope, nil
}

// InjectTo inject arguments to object
func (scope *ArgScope) InjectTo(obj interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()
	for i := 0; i < structValue.NumField(); i++ {
		var isRequired = true
		valueField := structValue.Field(i)
		structField := structValue.Type().Field(i)
		key := structField.Tag.Get(scope.tagname)
		if key == "" {
			continue
		}
		if strings.HasPrefix(key, "?") {
			isRequired = false
			key = key[1:]
		}
		switch valueField.Interface().(type) {
		case string:
			value, ok := scope.data[key]
			if !ok {
				if !isRequired {
					continue
				}
				return fmt.Errorf("Don't contains a argument for %v", key)
			}
			valueField.Set(reflect.ValueOf(value))
		default:
			value, ok := scope.data[key]
			if !ok {
				if !isRequired {
					continue
				}
				return fmt.Errorf("Don't contains a string for %v", key)
			}
			r.SetValueFromString(valueField, value)
		}
	}
	return nil
}

// Set new scope value
func (scope *ArgScope) Set(key string, v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("%v is not a string", v)
	}
	scope.data[key] = str
	return nil
}

// Get get value from context
func (scope *ArgScope) Get(key string) (interface{}, error) {
	return scope.data[key], nil
}

// Keys get map data
func (scope *ArgScope) Keys() ([]string, error) {
	keys := make([]string, len(scope.data))
	i := 0
	for key := range scope.data {
		keys[i] = key
		i++
	}
	return keys, nil
}
