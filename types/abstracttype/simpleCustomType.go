package abstracttype

import "github.com/goatcms/goat-core/types"

// SimpleCustomType is abstract to storage single types
type SimpleCustomType struct {
	types.MetaType
	types.TypeConverter
	types.TypeValidator
}

// AddSubTypes add sub types to a map
func (bt *SimpleCustomType) AddSubTypes(base string, m map[string]types.CustomType) {
	m[base] = bt
}

// GetSubTypes return map of sub types
func (bt *SimpleCustomType) GetSubTypes() map[string]types.CustomType {
	m := make(map[string]types.CustomType)
	bt.AddSubTypes("", m)
	return m
}
