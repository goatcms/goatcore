package simpletype

import (
	"reflect"

	"github.com/goatcms/goat-core/types"
	"github.com/goatcms/goat-core/types/abstracttype"
	"github.com/goatcms/goat-core/types/validator"
)

// ContentType represent content field type
type ContentType struct {
	abstracttype.MetaType
	abstracttype.StringConverter
	validator.EmptyValidator
}

// NewContentType create new instance of content type with custom attributes
func NewContentType(attrs map[string]string) types.CustomType {
	var ptr *string
	return &abstracttype.CustomType{
		SingleCustomType: &ContentType{
			MetaType: abstracttype.MetaType{
				SQLTypeName:  "text",
				HTMLTypeName: "text",
				GoTypeRef:    reflect.TypeOf(ptr).Elem(),
				Attributes:   attrs,
			},
		},
	}
}
