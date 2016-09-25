package simpletype

import (
	"reflect"

	"github.com/goatcms/goat-core/types"
	"github.com/goatcms/goat-core/types/abstracttype"
	"github.com/goatcms/goat-core/types/validator"
)

// ImageType represent image field type
type ImageType struct {
	abstracttype.MetaType
	abstracttype.FilespaceConverter
	validator.EmptyValidator
	//TODO: Add image validation
}

// NewImageType create new instance of a image file type
func NewImageType(attrs map[string]string) types.CustomType {
	var ptr *types.File
	return &abstracttype.CustomType{
		SingleCustomType: &ImageType{
			MetaType: abstracttype.MetaType{
				SQLTypeName:  "varchar(500)",
				HTMLTypeName: "file",
				GoTypeRef:    reflect.TypeOf(ptr).Elem(),
				Attributes:   attrs,
			},
		},
	}
}
