package abstracttype

import "github.com/goatcms/goat-core/types"

// CustomType is abstract to storage single types
type CustomType struct {
	types.SingleCustomType
}

// AddSubTypes add sub types to a map
func (bt *CustomType) AddSubTypes(base string, m map[string]types.CustomType) {
	m[base] = bt
}

// GetSubTypes return map of sub types
func (bt *CustomType) GetSubTypes() map[string]types.CustomType {
	m := make(map[string]types.CustomType)
	bt.AddSubTypes("", m)
	return m
}
