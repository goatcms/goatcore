package validator

import (
	"github.com/asaskevich/govalidator"
	"github.com/goatcms/goat-core/types"
)

const (
	// InvalidEmail is key of invalid email message
	InvalidEmail = "validator.email"
)

// IsValidEmail return error validation message or NoErr
func IsValidEmail(value string) (string, error) {
	if govalidator.IsEmail(value) {
		return types.NoErr, nil
	}
	return InvalidEmail, nil
}
