package orm

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// DBTX represent single database table
type DBTX interface {
	Queryx(query string, args ...interface{}) (*sqlx.Rows, error)
	QueryRowx(query string, args ...interface{}) *sqlx.Row
	NamedExec(query string, arg interface{}) (sql.Result, error)
	MustExec(query string, args ...interface{}) sql.Result
}

// IDContainer is aheler to contains id
type IDContainer struct {
	ID int64 `db:"id"`
}
