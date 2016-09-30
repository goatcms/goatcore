package validator

import (
	"github.com/goatcms/goat-core/types"
	"github.com/goatcms/goat-core/types/typemsg"
)

// EmptyValidator represent not valid field
type EmptyValidator struct{}

// AddValid is empty validator
func (v EmptyValidator) AddValid(ival interface{}, basekey string, mm types.MessageMap) error {
	return nil
}

// Valid valid a value and return valid result list
func (v EmptyValidator) Valid(ival interface{}) (types.MessageMap, error) {
	mm := typemsg.NewMessageMap()
	if err := v.AddValid(ival, "", mm); err != nil {
		return nil, err
	}
	return mm, nil
}
