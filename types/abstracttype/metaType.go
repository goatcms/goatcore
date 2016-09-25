package abstracttype

import (
	"reflect"

	"github.com/goatcms/goat-core/types"
)

const (
	// MainElement is default name for eleemnt in tables with single field
	MainElement = "id"
)

// MetaType is abstract for image types
type MetaType struct {
	SQLTypeName  string
	HTMLTypeName string
	GoTypeRef    reflect.Type
	Validators   []types.Validator
	Attributes   map[string]string
}

// AddSQLType add types to sql types map
func (bt *MetaType) AddSQLType(base string, m map[string]string) {
	m[base] = bt.SQLTypeName
}

// GetSQLType return map of sql types for type
func (bt *MetaType) GetSQLType() map[string]string {
	m := map[string]string{}
	m[MainElement] = bt.SQLTypeName
	return m
}

// HTMLType return a HTML type name
func (bt *MetaType) HTMLType() string {
	return bt.HTMLTypeName
}

// GoType return a GoLang type
func (bt *MetaType) GoType() reflect.Type {
	return bt.GoTypeRef
}

// HasAttribute return true if type has attribute
func (bt *MetaType) HasAttribute(name string) bool {
	_, ok := bt.Attributes[name]
	return ok
}

// GetAttribute return value of attribute
func (bt *MetaType) GetAttribute(name string) string {
	return bt.Attributes[name]
}
