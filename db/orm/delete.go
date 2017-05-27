package orm

import (
	"database/sql"
	"fmt"

	"github.com/goatcms/goatcore/db"
)

// DeleteContext is context for findByID function
type DeleteContext struct {
	query string
}

// Insert create new record
func (q DeleteContext) Delete(tx db.TX, id int64) error {
	var (
		res   sql.Result
		err   error
		count int64
	)
	if res, err = tx.NamedExec(q.query, &IDContainer{ID: id}); err != nil {
		return fmt.Errorf("%s: %s", err.Error(), q.query)
	}
	if count, err = res.RowsAffected(); err != nil {
		return fmt.Errorf("%s: %s", err.Error(), q.query)
	}
	if count != 1 {
		return fmt.Errorf("Delete more than one record (%v records deleted)", count)
	}
	return nil
}

// NewDelete create new delete function instance
func NewDelete(table db.Table, driver db.Driver) (db.Delete, error) {
	dsql := driver.DSQL()
	query, err := dsql.NewDeleteWhereSQL(table.Name(), "id=:id")
	if err != nil {
		return nil, err
	}
	context := &DeleteContext{
		query: query,
	}
	return context.Delete, nil
}
