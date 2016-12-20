package simpletype

import (
	"reflect"

	"github.com/goatcms/goat-core/types"
	"github.com/goatcms/goat-core/types/abstracttype"
	"github.com/goatcms/goat-core/types/validator"
)

// NewInt64Type create new instance of nt64 type with custom attributes
func NewInt64Type(attrs map[string]string) types.CustomType {
	var ptr *int64
	return &abstracttype.SimpleCustomType{
		MetaType: &abstracttype.MetaType{
			SQLTypeName:  "int",
			HTMLTypeName: "number",
			GoTypeRef:    reflect.TypeOf(ptr).Elem(),
			Attributes:   attrs,
		},
		TypeConverter: abstracttype.NewInt64Converter(),
		TypeValidator: validator.NewNoValidator(),
	}
}
