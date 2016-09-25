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

// SQLType return a SQL type name
func (bt *MetaType) SQLType() string {
	return bt.SQLTypeName
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
