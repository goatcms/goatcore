package simpletype

import (
	"reflect"

	"github.com/goatcms/goat-core/types"
	"github.com/goatcms/goat-core/types/abstracttype"
	"github.com/goatcms/goat-core/types/validator"
)

// IDType represent id type
type IDType struct {
	abstracttype.MetaType
	abstracttype.StringConverter
	validator.EmptyValidator
}

// NewIDType create new instance of id type
func NewIDType(attrs map[string]string) types.CustomType {
	var ptr *int64
	return &abstracttype.SimpleCustomType{
		SingleCustomType: &IDType{
			MetaType: abstracttype.MetaType{
				SQLTypeName:  "int primary key",
				HTMLTypeName: "number",
				GoTypeRef:    reflect.TypeOf(ptr).Elem(),
				Attributes:   attrs,
			},
		},
	}
}
