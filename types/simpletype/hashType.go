package simpletype

import (
	"reflect"

	"github.com/goatcms/goat-core/types"
	"github.com/goatcms/goat-core/types/abstracttype"
	"github.com/goatcms/goat-core/types/validator"
)

// NewHashType create new instance of a hash type
func NewHashType(attrs map[string]string) types.CustomType {
	var ptr *string
	return &abstracttype.SimpleCustomType{
		MetaType: &abstracttype.MetaType{
			SQLTypeName:  "varchar(100)",
			HTMLTypeName: "text",
			GoTypeRef:    reflect.TypeOf(ptr).Elem(),
			Attributes:   attrs,
		},
		TypeConverter: abstracttype.NewStringConverter(),
		TypeValidator: validator.NewNoValidator(),
	}
}
