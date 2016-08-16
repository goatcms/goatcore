package simpletype

import "github.com/goatcms/goat-core/types"

// IDType is global id type instance
var IDType = &SimpleCustomType{
	SQLTypeName:  "int",
	HTMLTypeName: "number",
	Validators:   []types.Validator{},
	Attributes:   []string{types.Primary},
}
