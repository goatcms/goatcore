package orm

import "github.com/goatcms/goatcore/db"

// InsertWithIDContext is context for findByID function
type InsertWithIDContext struct {
	query string
}

// Insert create new record
func (q InsertWithIDContext) InsertWithID(tx db.TX, entity interface{}) error {
	if _, err := tx.NamedExec(q.query, entity); err != nil {
		return err
	}
	return nil
}

// NewInsertWithID create new dao function instance
func NewInsertWithID(table db.Table, dsql db.DSQL) (db.InsertWithID, error) {
	query, err := dsql.NewInsertSQL(table.Name(), append(table.Fields(), "id"))
	if err != nil {
		return nil, err
	}
	context := &InsertWithIDContext{
		query: query,
	}
	return context.InsertWithID, nil
}
