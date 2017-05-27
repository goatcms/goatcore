package orm

import (
	"strings"

	"github.com/goatcms/goatcore/db"
)

// InsertContext is context for findByID function
type InsertContext struct {
	query  string
	driver db.Driver
}

// Insert create new record
func (q InsertContext) Insert(tx db.TX, entity interface{}) (int64, error) {
	return q.driver.RunInsert(tx, q.query, entity)
}

// InsertContext create new dao function instance
func NewInsert(table db.Table, driver db.Driver) (db.Insert, error) {
	dsql := driver.DSQL()
	fromFields := table.Fields()
	fields := make([]string, len(fromFields))
	i := 0
	for _, v := range fromFields {
		if strings.ToLower(v) != "id" {
			fields[i] = v
			i++
		}
	}
	fields = fields[:i]
	query, err := dsql.NewInsertSQL(table.Name(), fields)
	if err != nil {
		return nil, err
	}
	context := &InsertContext{
		query:  query,
		driver: driver,
	}
	return context.Insert, nil
}
