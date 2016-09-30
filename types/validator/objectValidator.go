package validator

import (
	"github.com/goatcms/goat-core/types"
	"github.com/goatcms/goat-core/types/typemsg"
	"github.com/goatcms/goat-core/varutil"
)

// ObjectValidator is default validator for all object field
type ObjectValidator struct {
	Types map[string]types.CustomType
}

// AddValid add new valid result to a list
func (v ObjectValidator) AddValid(ival interface{}, basekey string, mm types.MessageMap) error {
	var (
		err      error
		fieldVal interface{}
	)
	if basekey != "" {
		basekey = basekey + "."
	}
	for name, validator := range v.Types {
		if fieldVal, err = varutil.GetField(ival, name); err != nil {
			return err
		}
		if err = validator.AddValid(fieldVal, basekey+name, mm); err != nil {
			return err
		}
	}
	return nil
}

// Valid valid value and return valid messages
func (v ObjectValidator) Valid(ival interface{}) (types.MessageMap, error) {
	mm := typemsg.NewMessageMap()
	if err := v.AddValid(ival, "", mm); err != nil {
		return nil, err
	}
	return mm, nil
}
