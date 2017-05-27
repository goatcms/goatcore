package orm

import (
	"database/sql"
	"fmt"

	"github.com/goatcms/goatcore/db"
)

// UpdateContext is context for findByID function
type UpdateContext struct {
	query string
}

// Insert create new record
func (q UpdateContext) Update(tx db.TX, entity interface{}) error {
	var (
		res   sql.Result
		err   error
		count int64
	)
	if res, err = tx.NamedExec(q.query, entity); err != nil {
		return fmt.Errorf("%s: %s", err.Error(), q.query)
	}
	if count, err = res.RowsAffected(); err != nil {
		return fmt.Errorf("%s: %s", err.Error(), q.query)
	}
	if count != 1 {
		return fmt.Errorf("Update modified more then one record (%v records modyfieds): %s", count, q.query)
	}
	return nil
}

// NewUpdate create new dao function instance
func NewUpdate(table db.Table, driver db.Driver) (db.Update, error) {
	dsql := driver.DSQL()
	query, err := dsql.NewUpdateSQL(table.Name(), table.Fields())
	if err != nil {
		return nil, err
	}
	context := &UpdateContext{
		query: query,
	}
	return context.Update, nil
}
