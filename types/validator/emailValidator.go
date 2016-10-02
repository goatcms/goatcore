package validator

import (
	"fmt"
	"regexp"

	"github.com/goatcms/goat-core/types"
	"github.com/goatcms/goat-core/types/typemsg"
)

const (
	// InvalidEmail is key of invalid email message
	InvalidEmail = "validator.email"
)

var (
	// regexpValidator is regullar expression for email
	regexpValidator = regexp.MustCompile("^[a-zA-Z0-9.!#$%&’*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\\.[a-zA-Z0-9-]+)*$")
)

// EmailValidator is email valid object
type EmailValidator struct {
}

// NewEmailValidator create new email vaidator
func NewEmailValidator() EmailValidator {
	return EmailValidator{}
}

// AddValid check email
func (ev EmailValidator) AddValid(ival interface{}, basekey string, mm types.MessageMap) error {
	value, ok := ival.(string)
	if !ok {
		return fmt.Errorf("EmailValidator support only string as input")
	}
	if !regexpValidator.MatchString(value) {
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
