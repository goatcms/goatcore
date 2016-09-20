package validator

import "github.com/goatcms/goat-core/types"

// ListValidator represent list of a validator for a field
type ListValidator struct {
	Validators []types.Validator
}

// Valid valid value for a list
func (lv *ListValidator) Valid(ival interface{}, basekey string, mm types.MessageMap) error {
	for _, validator := range lv.Validators {
		if err := validator(ival, basekey, mm); err != nil {
			return err
		}
	}
	return nil
}
