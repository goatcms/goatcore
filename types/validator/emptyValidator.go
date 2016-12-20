package validator

import (
	"github.com/goatcms/goat-core/types"
	"github.com/goatcms/goat-core/types/typemsg"
)

var (
	instance *NoValidator
)

// NoValidator represent not valid field
type NoValidator struct{}

// NewNoValidator return instance of no validate object
func NewNoValidator() *NoValidator {
	if instance == nil {
		instance = &NoValidator{}
	}
	return instance
}

// AddValid is empty validator
func (v NoValidator) AddValid(ival interface{}, basekey string, mm types.MessageMap) error {
	return nil
}

// Valid valid a value and return valid result list
func (v NoValidator) Valid(ival interface{}) (types.MessageMap, error) {
	mm := typemsg.NewMessageMap()
	if err := v.AddValid(ival, "", mm); err != nil {
		return nil, err
	}
	return mm, nil
}
