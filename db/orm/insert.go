package orm

import (
	"fmt"
	"math/rand"

	"github.com/goatcms/goatcore/db"
	"github.com/goatcms/goatcore/varutil"
)

// InsertContext is context for findByID function
type InsertContext struct {
	query string
}

// Insert create new record
func (q InsertContext) Insert(tx db.TX, entity interface{}) (int64, error) {
	var (
		//res sql.Result
		err error
		id  int64
	)
	id = rand.Int63()
	if err = varutil.SetField(entity, "ID", id); err != nil {
		return -1, fmt.Errorf("%s: %s", err.Error(), q.query)
	}
	if _, err = tx.NamedExec(q.query, entity); err != nil {
		return -1, fmt.Errorf("%s: %s", err.Error(), q.query)
	}
	return id, nil
}

// InsertContext create new dao function instance
func NewInsert(table db.Table, dsql db.DSQL) (db.Insert, error) {
	query, err := dsql.NewInsertSQL(table.Name(), table.Fields())
	if err != nil {
		return nil, err
	}
	context := &InsertContext{
		query: query,
	}
	return context.Insert, nil
}
