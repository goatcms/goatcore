package orm

import "github.com/goatcms/goatcore/db"

// CreateTableContext is context for findByID function
type CreateTableContext struct {
	query string
}

// CreateTable create new table
func (q CreateTableContext) CreateTable(tx db.TX) error {
	tx.MustExec(q.query)
	return nil
}

// NewCreateTable create new CreateTable instance
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
