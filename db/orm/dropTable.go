package orm

import "github.com/goatcms/goatcore/db"

// DropTableContext is context for DropTable function
type DropTableContext struct {
	query string
}

// DropTable remove table with data
func (q DropTableContext) DropTable(tx db.TX) error {
	tx.MustExec(q.query)
	return nil
}

// NewDropTable create new dao function instance
func NewDropTable(table db.Table, dsql db.DSQL) (db.DropTable, error) {
	query, err := dsql.NewDropTableSQL(table.Name())
	if err != nil {
		return nil, err
	}
	context := &DropTableContext{
		query: query,
	}
	return context.DropTable, nil
}
