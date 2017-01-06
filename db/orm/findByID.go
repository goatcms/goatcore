package orm

import (
	"github.com/goatcms/goat-core/db"
	"github.com/goatcms/goat-core/dependency"
)

// FindByIDContext is context for findByID function
type FindByIDContext struct {
	query string
}

// FindByIDContext return row by id
func (q FindByIDContext) FindByID(tx db.TX, id int64) (db.Row, error) {
	row, err := tx.QueryRowx(q.query, id)
	return row.(db.Row), err
}

// NewFindByID create new dao function instance
func NewFindByID(table db.Table, dsql db.DSQL) (db.FindByID, error) {
	query, err := dsql.NewSelectWhereSQL(table.Name(), table.Fields(), "id=:$1")
	if err != nil {
		return nil, err
	}
	context := &FindByIDContext{
		query: query,
	}
	return context.FindByID, nil
}

// CreateTableContext create new CreateTable function
func NewFindByIDFactory(table db.Table) dependency.Factory {
	return func(dp dependency.Provider) (interface{}, error) {
		dsqlIns, err := dp.Get(db.DSQLService)
		if err != nil {
			return nil, err
		}
		dsql := dsqlIns.(db.DSQL)
		return NewFindByID(table, dsql)
	}
}
