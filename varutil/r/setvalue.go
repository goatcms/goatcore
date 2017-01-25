package r

import (
	"fmt"
	"reflect"

	"github.com/goatcms/goat-core/varutil/totype"
)

func SetValueFromString(valueField reflect.Value, value string) error {
	if !valueField.IsValid() {
		return fmt.Errorf("r.SetValueFromString: %s is not valid", valueField)
	}
	if !valueField.CanSet() {
		return fmt.Errorf("r.SetValueFromString: Cannot set %s field value", valueField)
	}
	switch valueField.Interface().(type) {
	case string:
		valueField.SetString(value)
	case *string:
		refValue := reflect.ValueOf(&value)
		valueField.Set(refValue)
	case int:
		val, err := totype.StringToInt(value)
		if err != nil {
			return err
		}
		refValue := reflect.ValueOf(val)
		valueField.Set(refValue)
	case *int:
		val, err := totype.StringToInt(value)
		if err != nil {
			return err
		}
		refValue := reflect.ValueOf(&val)
		valueField.Set(refValue)
	case int16:
		val, err := totype.StringToInt16(value)
		if err != nil {
			return err
		}
		refValue := reflect.ValueOf(val)
		valueField.Set(refValue)
	case *int16:
		val, err := totype.StringToInt16(value)
		if err != nil {
			return err
		}
		refValue := reflect.ValueOf(&val)
		valueField.Set(refValue)
	case int32:
		val, err := totype.StringToInt32(value)
		if err != nil {
			return err
		}
		refValue := reflect.ValueOf(val)
		valueField.Set(refValue)
	case *int32:
		val, err := totype.StringToInt32(value)
		if err != nil {
			return err
		}
		refValue := reflect.ValueOf(&val)
		valueField.Set(refValue)
	case int64:
		val, err := totype.StringToInt64(value)
		if err != nil {
			return err
		}
		refValue := reflect.ValueOf(val)
		valueField.Set(refValue)
	case *int64:
		val, err := totype.StringToInt64(value)
		if err != nil {
			return err
		}
		refValue := reflect.ValueOf(&val)
		valueField.Set(refValue)
	default:
		return fmt.Errorf("unsupported value type %v", valueField.Type().String())
	}
	return nil
}
