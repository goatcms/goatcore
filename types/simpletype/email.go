package simpletype

import (
	"github.com/goatcms/goat-core/types"
	"github.com/goatcms/goat-core/types/validator"
)

// EmailType is global email type instance
var EmailType = NewEmailType([]string{})

// NewEmailType create new instance of email type with custom attributes
func NewEmailType(attrs []string) types.CustomType {
	return &SimpleCustomType{
		SQLTypeName:  "varchar(100)",
		HTMLTypeName: "text",
		Validators:   []types.Validator{validator.IsValidEmail},
		Attributes:   attrs,
	}
}
