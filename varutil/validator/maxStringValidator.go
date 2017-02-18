package validator

import "github.com/goatcms/goatcore/messages"

const (
	// InvalidMaxLength is key for too long strings
	InvalidMaxLength = "validator.length.max"
)

// MinStringValid add error message if string is shorten then some value
func MaxStringValid(value string, basekey string, mm messages.MessageMap, max int) error {
	if len(value) > max {
		mm.Add(basekey, InvalidMaxLength)
		return nil
	}
	return nil
}
