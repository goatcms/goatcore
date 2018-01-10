package validator

import "github.com/goatcms/goatcore/messages"

// MinStringValid add error message if string is shorten then some value
func MaxStringValid(value string, basekey string, mm messages.MessageMap, max int) error {
	if len(value) > max {
		mm.Add(basekey, InvalidMaxLength)
		return nil
	}
	return nil
}
