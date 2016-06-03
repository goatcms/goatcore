package definition

import (
	"fmt"
	"reflect"
)

func unpackValue(v reflect.Value) (reflect.Value, error) {
	if v.IsNil() {
		return v, nil
	}
	k := v.Type().Kind()
	if k == reflect.Map {
		return v, nil
	} else if k == reflect.Ptr {
		return unpackValue(v.Elem())
	} else if k == reflect.Interface {
		return unpackValue(v.Elem())
	} else {
		return v, fmt.Errorf("Unsupported type")
	}
}
