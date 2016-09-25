package simpletype

import (
	"reflect"

	"github.com/goatcms/goat-core/types"
	"github.com/goatcms/goat-core/types/abstracttype"
	"github.com/goatcms/goat-core/types/validator"
)

// Int64Type represent int64 type
type Int64Type struct {
	abstracttype.MetaType
	abstracttype.Int64Converter
	validator.EmptyValidator
}

// NewInt64Type create new instance of nt64 type with custom attributes
func NewInt64Type(attrs map[string]string) types.CustomType {
	var ptr *int64
	return &abstracttype.SimpleCustomType{
		SingleCustomType: &Int64Type{
			MetaType: abstracttype.MetaType{
				SQLTypeName:  "int",
				HTMLTypeName: "number",
				GoTypeRef:    reflect.TypeOf(ptr).Elem(),
				Attributes:   attrs,
			},
		},
	}
}
