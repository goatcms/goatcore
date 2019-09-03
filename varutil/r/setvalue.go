package r

import (
	"reflect"

	"github.com/goatcms/goatcore/varutil/goaterr"
	"github.com/goatcms/goatcore/varutil/totype"
)

// SetValueFromString convert string to reflect.Value type and set it.
func SetValueFromString(valueField reflect.Value, value string) error {
	if !valueField.IsValid() {
		return goaterr.Errorf("r.SetValueFromString: %s is not valid", valueField)
	}
	if !valueField.CanSet() {
		return goaterr.Errorf("r.SetValueFromString: Cannot set %s field value", valueField)
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
	case bool:
		val, err := totype.StringToBool(value)
		if err != nil {
			return err
		}
		refValue := reflect.ValueOf(val)
		valueField.Set(refValue)
	case *bool:
		val, err := totype.StringToBool(value)
		if err != nil {
			return err
		}
		refValue := reflect.ValueOf(&val)
		valueField.Set(refValue)
	default:
		return goaterr.Errorf("unsupported value type %v", valueField.Type().String())
	}
	return nil
}
