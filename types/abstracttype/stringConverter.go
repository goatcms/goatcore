package abstracttype

import (
	"fmt"

	"github.com/goatcms/goat-core/types"
)

var (
	instance *StringConverter
)

// StringConverter is converter for strings
type StringConverter struct{}

// NewStringConverter return instance of string converter
func NewStringConverter() *StringConverter {
	if instance == nil {
		instance = &StringConverter{}
	}
	return instance
}

// FromString decode string value
func (s *StringConverter) FromString(value string) (interface{}, error) {
	return value, nil
}

// FromMultipart convert multipartdata to string
func (s *StringConverter) FromMultipart(fh types.FileHeader) (interface{}, error) {
	return StringFromMultipart(fh)
}

// ToString change object to string
func (s *StringConverter) ToString(ival interface{}) (string, error) {
	value, ok := ival.(string)
	if !ok {
		return "", fmt.Errorf("It is string type. It aceppt only string type.")
	}
	return value, nil
}
