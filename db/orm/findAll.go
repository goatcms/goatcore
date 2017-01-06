package orm

import (
	"github.com/goatcms/goat-core/db"
	"github.com/goatcms/goat-core/dependency"
)

// FindAllContext is context for findAll function
type FindAllContext struct {
	query string
}

// FindAll obtain all articles from database
func (q FindAllContext) FindAll(tx db.TX) (db.Rows, error) {
	rows, err := tx.Queryx(q.query)
	return rows.(db.Rows), err
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

// CreateTableContext create new CreateTable function
func NewFindAllFactory(table db.Table) dependency.Factory {
	return func(dp dependency.Provider) (interface{}, error) {
		dsqlIns, err := dp.Get(db.DSQLService)
		if err != nil {
			return nil, err
		}
		dsql := dsqlIns.(db.DSQL)
		return NewFindAll(table, dsql)
	}
}
