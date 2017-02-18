package orm

import "github.com/goatcms/goatcore/db"

// IDContainer is simple struct to contains id
type IDContainer struct {
	ID int64 `db:"id"`
}

// IDContainer is simple struct to contains id
type QueryDependency struct {
	DSQL db.DSQL
}
