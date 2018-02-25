package orm

import "github.com/goatcms/goatcore/db"

// IDContainer is simple struct to contains id
type IDContainer struct {
	ID int64 `db:"id"`
}

// QueryDependency is simple struct to contains dependencies
type QueryDependency struct {
	DSQL db.DSQL
}
