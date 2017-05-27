package itedriver

import (
	"database/sql"

	"github.com/goatcms/goatcore/db"
	"github.com/jmoiron/sqlx"
)

type Tx struct {
	*sqlx.Tx
}

func (t Tx) Queryx(query string, args ...interface{}) (db.Rows, error) {
	rows, err := t.Tx.Queryx(query, args...)
	return db.Rows(rows), err
}

func (t Tx) QueryRowx(query string, args ...interface{}) (db.Row, error) {
	row := t.Tx.QueryRowx(query, args...)
	return db.Row(row), row.Err()
}

func (t Tx) MustExec(query string, data interface{}) sql.Result {
	result, err := t.Tx.Exec(query, data)
	if err != nil {
		panic(err)
	}
	return result
}
