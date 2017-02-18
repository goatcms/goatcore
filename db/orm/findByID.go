package orm

import (
	"strconv"
	"strings"

	"github.com/goatcms/goatcore/db"
)

// FindByIDContext is context for findByID function
type FindByIDContext struct {
	query string
}

// FindByIDContext return row by id
func (q FindByIDContext) FindByID(tx db.TX, id int64) (db.Row, error) {
	row, err := tx.QueryRowx(strings.Replace(q.query, ":id", strconv.FormatInt(id, 10), -1))
	return row.(db.Row), err
}

// NewFindByID create new dao function instance
func NewFindByID(table db.Table, dsql db.DSQL) (db.FindByID, error) {
	query, err := dsql.NewSelectWhereSQL(table.Name(), table.Fields(), "id=:id LIMIT 1")
	if err != nil {
		return nil, err
	}
	context := &FindByIDContext{
		query: query,
	}
	return context.FindByID, nil
}
