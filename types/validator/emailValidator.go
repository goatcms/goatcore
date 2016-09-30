package validator

import (
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/goatcms/goat-core/types"
	"github.com/goatcms/goat-core/types/typemsg"
)

const (
	// InvalidEmail is key of invalid email message
	InvalidEmail = "validator.email"
)

// EmailValidator is email valid object
type EmailValidator struct{}

// AddValid check email
func (ev EmailValidator) AddValid(ival interface{}, basekey string, mm types.MessageMap) error {
	value, ok := ival.(string)
	if !ok {
		return fmt.Errorf("EmailValidator support only string as input")
	}
	if !govalidator.IsEmail(value) {
		mm.Add(basekey, InvalidEmail)
		return nil
	}
	return nil
}

// Valid valid a value and return valid result list
func (ev EmailValidator) Valid(ival interface{}) (types.MessageMap, error) {
	mm := typemsg.NewMessageMap()
	if err := ev.AddValid(ival, "", mm); err != nil {
		return nil, err
	}
	return mm, nil
}
