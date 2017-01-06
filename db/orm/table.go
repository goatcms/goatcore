package orm

import (
	"github.com/goatcms/goat-core/db"
	"github.com/goatcms/goat-core/types"
)

// Table represent a database table
type Table struct {
	name   string
	fields []string
	types  map[string]types.CustomType
}

func (t Table) Name() string {
	return t.name
}

func (t Table) Fields() []string {
	return t.fields
}

func (t Table) Types() map[string]types.CustomType {
	return t.types
}

// NewTable create new base database table accessor
func NewTable(name string, types map[string]types.CustomType) db.Table {
	fields := make([]string, len(types))
	i := 0
	for name, _ := range types {
		fields[i] = name
		i++
	}
	return &Table{
		name:   name,
		types:  types,
		fields: fields,
	}
}
