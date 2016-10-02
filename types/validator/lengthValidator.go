package validator

import (
	"fmt"

	"github.com/goatcms/goat-core/types"
	"github.com/goatcms/goat-core/types/typemsg"
)

const (
	// InvalidMinLength is key for too short strings
	InvalidMinLength = "validator.length.min"
	// InvalidMaxLength is key for too length strings
	InvalidMaxLength = "validator.length.max"
)

// LengthValidator is string length validator
type LengthValidator struct {
	min int
	max int
}

// NewLengthValidator create new length validator
func NewLengthValidator(min, max int) LengthValidator {
	return LengthValidator{
		min: min,
		max: max,
	}
}

// AddValid check length
func (lv LengthValidator) AddValid(ival interface{}, basekey string, mm types.MessageMap) error {
	value, ok := ival.(string)
	if !ok {
		return fmt.Errorf("LengthValidator support only string as input")
	}
	if len(value) < lv.min {
		mm.Add(basekey, InvalidMinLength)
		return nil
	}
	if len(value) > lv.max {
		mm.Add(basekey, InvalidMaxLength)
		return nil
	}
	return nil
}

// Valid valid a value and return valid result list
func (lv LengthValidator) Valid(ival interface{}) (types.MessageMap, error) {
	mm := typemsg.NewMessageMap()
	if err := lv.AddValid(ival, "", mm); err != nil {
		return nil, err
	}
	return mm, nil
}
