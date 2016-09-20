package validator

import (
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/goatcms/goat-core/types"
)

const (
	// InvalidEmail is key of invalid email message
	InvalidEmail = "validator.email"
)

// EmailValidator is email valid object
type EmailValidator struct{}

// Valid check email
func (ev EmailValidator) Valid(ival interface{}, basekey string, mm types.MessageMap) error {
	value, ok := ival.(string)
	if !ok {
		return fmt.Errorf("EmailValidator support only string as input")
	}
	if govalidator.IsEmail(value) {
		mm.Add(basekey, InvalidEmail)
		return nil
	}
	return nil
}
