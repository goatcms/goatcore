package orm

import (
	"github.com/goatcms/goat-core/db/dsql"
	"github.com/goatcms/goat-core/types"
)

// BaseTable represent a database table
type BaseTable struct {
	table           string
	fields          []string
	types           map[string]types.CustomType
	selectSQL       string
	selectByIDSQL   string
	insertSQL       string
	insertWithIDSQL string
	updateByIDSQL   string
	deleteByIDSQL   string
	createSQL       string
}

// NewBaseTable create new base database table accessor
func NewBaseTable(table string, types map[string]types.CustomType) *BaseTable {
	fieldsWithID := make([]string, len(types)+1)
	fieldsWithID[0] = "id"
	fields := make([]string, len(types))
	i := 0
	for name := range types {
		fields[i] = name
		fieldsWithID[i+1] = name
		i++
	}
	return &BaseTable{
		table:           table,
		types:           types,
		fields:          fields,
		selectSQL:       dsql.NewSelectSQL(table, fields),
		selectByIDSQL:   dsql.NewSelectWhereSQL(table, fields, "id=:$1"),
		insertSQL:       dsql.NewInsertSQL(table, fields),
		insertWithIDSQL: dsql.NewInsertSQL(table, fieldsWithID),
		updateByIDSQL:   dsql.NewUpdateWhereSQL(table, fields, "id=:id"),
		deleteByIDSQL:   dsql.NewDeleteWhereSQL(table, "id=:id"),
		createSQL:       dsql.NewCreateSQL(table, types),
	}
}
