package abstracttype

import (
	"reflect"

	"github.com/goatcms/goat-core/types"
)

const (
	// SQLTypeDotChar represen sql field separator
	SQLTypeDotChar = "_"
)

// ObjectMetaType is abstract for object fields types
type ObjectMetaType struct {
	Types        map[string]types.CustomType
	HTMLTypeName string
	GoTypeRef    reflect.Type
	Validators   []types.Validator
	Attributes   map[string]string
}

// AddSQLType add types to sql types map
func (bt *ObjectMetaType) AddSQLType(base string, m map[string]string) {
	if base != "" {
		base = base + SQLTypeDotChar
	}
	for key, typeIns := range bt.Types {
		typeIns.AddSQLType(base+key, m)
	}
}

// GetSQLType return map of sql types for type
func (bt *ObjectMetaType) GetSQLType() map[string]string {
	m := map[string]string{}
	bt.AddSQLType("", m)
	return m
}

// HTMLType return a HTML type name
func (bt *ObjectMetaType) HTMLType() string {
	return bt.HTMLTypeName
}

// GoType return a GoLang type
func (bt *ObjectMetaType) GoType() reflect.Type {
	return bt.GoTypeRef
}

// HasAttribute return true if type has attribute
func (bt *ObjectMetaType) HasAttribute(name string) bool {
	_, ok := bt.Attributes[name]
	return ok
}

// GetAttribute return value of attribute
func (bt *ObjectMetaType) GetAttribute(name string) string {
	return bt.Attributes[name]
}
