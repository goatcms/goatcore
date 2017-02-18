package validator

import (
	"regexp"

	"github.com/goatcms/goatcore/messages"
)

const (
	// InvalidEmail is key of invalid email message
	InvalidEmail = "validator.email"
)

var (
	// emailRegexp is regullar expression for email
	emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&’*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\\.[a-zA-Z0-9-]+)*$")
)

// EmailValid add error message if email is invalid
func EmailValid(value string, basekey string, mm messages.MessageMap) error {
	if !emailRegexp.MatchString(value) {
		mm.Add(basekey, InvalidEmail)
		return nil
	}
	return nil
}
