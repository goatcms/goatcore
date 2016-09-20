package simpletype

import (
	"reflect"

	"github.com/goatcms/goat-core/types"
	"github.com/goatcms/goat-core/types/abstracttype"
	"github.com/goatcms/goat-core/types/validator"
)

// FileType represent email field type
type FileType struct {
	abstracttype.MetaType
	abstracttype.FilespaceConverter
	validator.EmptyValidator
}

// NewFileType create new instance of a file type
func NewFileType(attrs map[string]string) types.CustomType {
	var ptr *types.File
	return &FileType{
		MetaType: abstracttype.MetaType{
			SQLTypeName:  "varchar(500)",
			HTMLTypeName: "file",
			GoTypeRef:    reflect.TypeOf(ptr).Elem(),
			Attributes:   attrs,
		},
	}
}
