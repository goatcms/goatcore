package abstracttype

import (
	"fmt"
	"strconv"

	"github.com/goatcms/goat-core/types"
)

// Int64Converter is abstract for types converi
type Int64Converter struct{}

// FromString decode string value
func (bt *Int64Converter) FromString(s string) (interface{}, error) {
	return strconv.ParseInt(s, 10, 64)
}

// FromMultipart convert multipartdata to int64
func (bt *Int64Converter) FromMultipart(fh types.FileHeader) (interface{}, error) {
	s, err := StringFromMultipart(fh)
	if err != nil {
		return nil, err
	}
	return bt.FromString(s)
}

// ToString change object to string
func (bt *Int64Converter) ToString(ival interface{}) (string, error) {
	value, ok := ival.(int64)
	if !ok {
		return "", fmt.Errorf("Int64Converter support only int64 as input")
	}
	return strconv.FormatInt(value, 10), nil
}
