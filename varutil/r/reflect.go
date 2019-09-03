package r

import (
	"reflect"

	"github.com/goatcms/goatcore/varutil/goaterr"
)

// UnpackValue skip pointers, interfaces etc and return a finish structure or a simple value
func UnpackValue(v reflect.Value) (reflect.Value, error) {
	if v.IsNil() {
		return v, nil
	}
	k := v.Type().Kind()
	if k == reflect.Map {
		return v, nil
	} else if k == reflect.Ptr {
		return UnpackValue(v.Elem())
	} else if k == reflect.Interface {
		return UnpackValue(v.Elem())
	}
	return v, goaterr.Errorf("Unsupported type")
}

/*
func IsNilable(v reflect.Value) bool {
	if v.IsNil() {
		return true
	}
	k := v.Type().Kind()
	if k == reflect.Map {
		return v, nil
	} else if k == reflect.Ptr {
		return unpackValue(v.Elem())
	} else if k == reflect.Interface {
		return unpackValue(v.Elem())
	} else {
		return v, goaterr.Errorf("Unsupported type")
	}
}*/
