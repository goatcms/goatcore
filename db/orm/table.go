package orm

import (
	"reflect"

	"github.com/goatcms/goatcore/db"
)

// Table represent a database table
type Table struct {
	name   string
	fields []string
	types  map[string]string
}

// Name return table name
func (t Table) Name() string {
	return t.name
}

// Fields return table fields array
func (t Table) Fields() []string {
	return t.fields
}

// Types return table fields types map
func (t Table) Types() map[string]string {
	return t.types
}

// NewTable create new base database table accessor
func NewTable(name string, entityType reflect.Type) db.Table {
	numFields := entityType.NumField()
	fields := make([]string, numFields)
	types := make(map[string]string)
	for i := 0; i < numFields; i++ {
		structField := entityType.Field(i)
		sqlTypeString := structField.Tag.Get(db.SQLTypeTagName)
		if sqlTypeString == "" {
			continue
		}
		fieldName := structField.Tag.Get("db")
		if fieldName == "" {
			fieldName = structField.Name
		}
		fields[i] = fieldName
		types[fieldName] = sqlTypeString
	}
	return &Table{
		name:   name,
		types:  types,
		fields: fields,
	}
}
