package simpletype

import (
	"reflect"

	"github.com/goatcms/goat-core/filesystem"
	"github.com/goatcms/goat-core/types"
	"github.com/goatcms/goat-core/types/abstracttype"
	"github.com/goatcms/goat-core/types/validator"
)

// NewFileType create new instance of a file type
func NewFileType(attrs map[string]string, fs filesystem.Filespace) types.CustomType {
	var ptr *types.File
	return &abstracttype.SimpleCustomType{
		MetaType: &abstracttype.MetaType{
			SQLTypeName:  "varchar(500)",
			HTMLTypeName: "file",
			GoTypeRef:    reflect.TypeOf(ptr).Elem(),
			Attributes:   attrs,
		},
		TypeConverter: abstracttype.NewFilespaceConverter(fs),
		TypeValidator: validator.NewNoValidator(),
	}
}
