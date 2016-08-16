package simpletype

import "github.com/goatcms/goat-core/types"

// ContentType is a global content type instance
var ContentType = NewContentType([]string{})

// NewContentType create new instance of content type with custom attributes
func NewContentType(attrs []string) types.CustomType {
	return &SimpleCustomType{
		SQLTypeName:  "text",
		HTMLTypeName: "text",
		Validators:   []types.Validator{},
		Attributes:   attrs,
	}
}
