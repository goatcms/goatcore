package orm

import (
	"fmt"

	"github.com/goatcms/goatcore/db"
)

// FindAllContext is context for findAll function
type FindAllContext struct {
	query string
}

// FindAll obtain all articles from database
func (q FindAllContext) FindAll(tx db.TX) (db.Rows, error) {
	rows, err := tx.Queryx(q.query)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", err.Error(), q.query)
	}
	return rows.(db.Rows), nil
}

// NewFindAll create new FindAll function
func NewFindAll(table db.Table, dsql db.DSQL) (db.FindAll, error) {
	query, err := dsql.NewSelectSQL(table.Name(), table.Fields())
	if err != nil {
		return nil, err
	}
	FindAllContext := &FindAllContext{
		query: query,
	}
	return FindAllContext.FindAll, nil
}
