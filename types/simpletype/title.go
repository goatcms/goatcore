package simpletype

import (
	"github.com/goatcms/goat-core/types"
	"github.com/goatcms/goat-core/types/validator"
)

// TitleType is global title type instance
var TitleType = NewTitleType([]string{})

// NewTitleType create new instance of title type with custom attributes
func NewTitleType(attrs []string) types.CustomType {
	return &SimpleCustomType{
		SQLTypeName:  "varchar(400)",
		HTMLTypeName: "text",
		Validators:   []types.Validator{validator.IsValidEmail},
		Attributes:   attrs,
	}
}
