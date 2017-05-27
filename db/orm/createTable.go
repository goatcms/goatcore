package orm

import "github.com/goatcms/goatcore/db"

// CreateTableContext is context for findByID function
type CreateTableContext struct {
	query  string
	driver db.Driver
}

// Insert create new record
func (q CreateTableContext) CreateTable(tx db.TX) error {
	q.driver.RunCreateTable(tx, q.query)
	return nil
}

// CreateTableContext create new CreateTable function
func NewCreateTable(table db.Table, driver db.Driver) (db.CreateTable, error) {
	dsql := driver.DSQL()
	query, err := dsql.NewCreateSQL(table.Name(), table.Types())
	if err != nil {
		return nil, err
	}
	context := &CreateTableContext{
		query:  query,
		driver: driver,
	}
	return context.CreateTable, nil
}
