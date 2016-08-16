package simpletype

import (
	"github.com/goatcms/goat-core/types"
	"github.com/goatcms/goat-core/types/validator"
)

// ImageType is global image type instance
var ImageType = NewImageType([]string{})

// NewImageType create new instance of a image type with custom attributes
func NewImageType(attrs []string) types.CustomType {
	return &SimpleCustomType{
		SQLTypeName:  "varchar(500)",
		HTMLTypeName: "file",
		Validators:   []types.Validator{validator.IsValidEmail},
		Attributes:   attrs,
	}
}
