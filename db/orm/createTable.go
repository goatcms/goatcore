package orm

import (
	"github.com/goatcms/goat-core/db"
	"github.com/goatcms/goat-core/dependency"
)

// CreateTableContext is context for findByID function
type CreateTableContext struct {
	query string
}

// Insert create new record
func (q CreateTableContext) CreateTable(tx db.TX) error {
	tx.MustExec(q.query)
	return nil
}

// CreateTableContext create new CreateTable function
func NewCreateTable(table db.Table, dsql db.DSQL) (db.CreateTable, error) {
	query, err := dsql.NewCreateSQL(table.Name(), table.Types())
	if err != nil {
		return nil, err
	}
	context := &CreateTableContext{
		query: query,
	}
	return context.CreateTable, nil
}

// CreateTableContext create new CreateTable function
func NewCreateTableFactory(table db.Table) dependency.Factory {
	return func(dp dependency.Provider) (interface{}, error) {
		dsqlIns, err := dp.Get(db.DSQLService)
		if err != nil {
			return nil, err
		}
		dsql := dsqlIns.(db.DSQL)
		return NewCreateTable(table, dsql)
	}
}
