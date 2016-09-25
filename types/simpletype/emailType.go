package simpletype

import (
	"reflect"

	"github.com/goatcms/goat-core/types"
	"github.com/goatcms/goat-core/types/abstracttype"
	"github.com/goatcms/goat-core/types/validator"
)

// EmailType represent email field type
type EmailType struct {
	abstracttype.MetaType
	abstracttype.StringConverter
	validator.EmailValidator
}

// NewEmailType create new instance of a email type
func NewEmailType(attrs map[string]string) types.CustomType {
	var ptr *string
	return &abstracttype.CustomType{
		SingleCustomType: &EmailType{
			MetaType: abstracttype.MetaType{
				SQLTypeName:  "varchar(100)",
				HTMLTypeName: "text",
				GoTypeRef:    reflect.TypeOf(ptr).Elem(),
				Attributes:   attrs,
			},
		},
	}
}
