package simpletype

import (
	"reflect"

	"github.com/goatcms/goat-core/types"
	"github.com/goatcms/goat-core/types/abstracttype"
	"github.com/goatcms/goat-core/types/validator"
)

// TitleType represent email field type
type TitleType struct {
	abstracttype.MetaType
	abstracttype.StringConverter
	validator.EmailValidator
}

// NewTitleType create new instance of a email type
func NewTitleType(attrs map[string]string) types.CustomType {
	var ptr *string
	return &abstracttype.CustomType{
		SingleCustomType: &TitleType{
			MetaType: abstracttype.MetaType{
				SQLTypeName:  "varchar(100)",
				HTMLTypeName: "text",
				GoTypeRef:    reflect.TypeOf(ptr).Elem(),
				Attributes:   attrs,
			},
		},
	}
}
