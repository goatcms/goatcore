package validator

import "github.com/goatcms/goat-core/types"

// EmptyValidator represent not valid field
type EmptyValidator struct{}

// Valid check email
func (v EmptyValidator) Valid(ival interface{}, basekey string, mm types.MessageMap) error {
	return nil
}
