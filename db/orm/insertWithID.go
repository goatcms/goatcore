package orm

import (
	"fmt"

	"github.com/goatcms/goatcore/db"
)

// InsertWithIDContext is context for findByID function
type InsertWithIDContext struct {
	query string
}

// Insert create new record
func (q InsertWithIDContext) InsertWithID(tx db.TX, entity interface{}) error {
	if _, err := tx.NamedExec(q.query, entity); err != nil {
		return fmt.Errorf("%s: %s", err.Error(), q.query)
	}
	return nil
}

// NewInsertWithID create new dao function instance
func NewInsertWithID(table db.Table, driver db.Driver) (db.InsertWithID, error) {
	dsql := driver.DSQL()
	query, err := dsql.NewInsertSQL(table.Name(), table.Fields())
	if err != nil {
		return nil, err
	}
	context := &InsertWithIDContext{
		query: query,
	}
	return context.InsertWithID, nil
}
