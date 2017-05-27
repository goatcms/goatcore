package orm

import "github.com/goatcms/goatcore/db"

// DropTableContext is context for DropTable function
type DropTableContext struct {
	query  string
	driver db.Driver
}

// DropTable remove table with data
func (q DropTableContext) DropTable(tx db.TX) error {
	q.driver.RunCreateTable(tx, q.query)
	return nil
}

// NewDropTable create new dao function instance
func NewDropTable(table db.Table, driver db.Driver) (db.DropTable, error) {
	query, err := driver.DSQL().NewDropTableSQL(table.Name())
	if err != nil {
		return nil, err
	}
	context := &DropTableContext{
		query:  query,
		driver: driver,
	}
	return context.DropTable, nil
}
