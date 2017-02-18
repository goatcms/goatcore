package adapter

import (
	"fmt"

	"github.com/goatcms/goatcore/db"
	"github.com/jmoiron/sqlx"
)

type SQLXDB struct {
	*sqlx.DB
}

func (x SQLXDB) Queryx(query string, args ...interface{}) (db.Rows, error) {
	rows, err := x.DB.Queryx(query, args...)
	return db.Rows(rows), err
}

func (x SQLXDB) QueryRowx(query string, args ...interface{}) (db.Row, error) {
	row := x.DB.QueryRowx(query, args...)
	return db.Row(row), row.Err()
}

func (x SQLXDB) Commit() error {
	return nil
}

func (x SQLXDB) Rollback() error {
	return fmt.Errorf("Database connection as transaction don't support rollback (all queries are autorun))")
}

func NewTXFromDB(db *sqlx.DB) db.TX {
	return SQLXDB{db}
}
