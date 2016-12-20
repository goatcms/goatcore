package abstracttype

import (
	"reflect"

	"github.com/goatcms/goat-core/types"
	"github.com/goatcms/goat-core/varutil"
)

// ObjectFactory is a new object factory
type ObjectFactory func() interface{}

// ObjectConverter represent a object abstract type
type ObjectConverter struct {
	ObjectFactory ObjectFactory
}

// NewObjectConverter create new instance of object converter
func NewObjectConverter(f ObjectFactory) *ObjectConverter {
	return &ObjectConverter{
		ObjectFactory: f,
	}
}

// NewObjectConverterFromType create new instance of object converter from reflect type
func NewObjectConverterFromType(t reflect.Type) *ObjectConverter {
	return &ObjectConverter{
		ObjectFactory: func() interface{} {
			return reflect.New(t).Interface()
		},
	}
}

// FromString decode string value
func (obt *ObjectConverter) FromString(s string) (interface{}, error) {
	o := obt.ObjectFactory()
	if err := varutil.ObjectFromJSON(o, s); err != nil {
		return nil, err
	}
	return o, nil
}

// FromMultipart convert multipartdata to int64
func (obt *ObjectConverter) FromMultipart(fh types.FileHeader) (interface{}, error) {
	s, err := StringFromMultipart(fh)
	if err != nil {
		return nil, err
	}
	return obt.FromString(s)
}

// ToString change object to string
func (obt *ObjectConverter) ToString(ival interface{}) (string, error) {
	return varutil.ObjectToJSON(ival)
}
