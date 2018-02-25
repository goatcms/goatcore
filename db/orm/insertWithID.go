package orm

import (
	"fmt"

	"github.com/goatcms/goatcore/db"
)

// InsertWithIDContext is context for findByID function
type InsertWithIDContext struct {
	query string
}

// InsertWithID insert new row with defined id
func (q InsertWithIDContext) InsertWithID(tx db.TX, entity interface{}) error {
	if _, err := tx.NamedExec(q.query, entity); err != nil {
		return fmt.Errorf("%s: %s", err.Error(), q.query)
	}
	return nil
}

// NewInsertWithID create new InsertWithID instance
func NewInsertWithID(table db.Table, dsql db.DSQL) (db.InsertWithID, error) {
	query, err := dsql.NewInsertSQL(table.Name(), table.Fields())
	if err != nil {
		return nil, err
	}
	context := &InsertWithIDContext{
		query: query,
	}
	return context.InsertWithID, nil
}
