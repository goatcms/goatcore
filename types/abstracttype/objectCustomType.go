package abstracttype

import "github.com/goatcms/goat-core/types"

// ObjectCustomType represent abstract to sub types iterate
type ObjectCustomType struct {
	types.MetaType
	types.TypeConverter
	types.TypeValidator
	Types map[string]types.CustomType
}

// AddSubTypes add sub types to sql types map
func (bt *ObjectCustomType) AddSubTypes(base string, m map[string]types.CustomType) {
	if base != "" {
		base = base + types.SQLSeparator
	}
	for key, typeIns := range bt.Types {
		typeIns.AddSubTypes(base+key, m)
	}
}

// GetSubTypes return map of sub custom types for current type
func (bt *ObjectCustomType) GetSubTypes() map[string]types.CustomType {
	m := make(map[string]types.CustomType)
	bt.AddSubTypes("", m)
	return m
}
