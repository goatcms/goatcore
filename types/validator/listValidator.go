package validator

import (
	"github.com/goatcms/goat-core/types"
	"github.com/goatcms/goat-core/types/typemsg"
)

// ListValidator represent list of a validator for a field
type ListValidator struct {
	Validators []types.TypeValidator
}

// AddValid add new valid result to a list
func (lv ListValidator) AddValid(ival interface{}, basekey string, mm types.MessageMap) error {
	for _, validator := range lv.Validators {
		if err := validator.AddValid(ival, basekey, mm); err != nil {
			return err
		}
	}
	return nil
}

// Valid valid value for a list
func (lv ListValidator) Valid(ival interface{}) (types.MessageMap, error) {
	mm := typemsg.NewMessageMap()
	if err := lv.AddValid(ival, "", mm); err != nil {
		return nil, err
	}
	return mm, nil
}
