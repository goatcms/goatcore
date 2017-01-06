package adapter

import (
	"github.com/goatcms/goat-core/db"
	"github.com/jmoiron/sqlx"
)

type SQLXTX struct {
	*sqlx.Tx
}

func (x SQLXTX) Queryx(query string, args ...interface{}) (db.Rows, error) {
	rows, err := x.Tx.Queryx(query, args...)
	return db.Rows(rows), err
}

func (x SQLXTX) QueryRowx(query string, args ...interface{}) (db.Row, error) {
	row := x.Tx.QueryRowx(query, args...)
	return db.Row(row), row.Err()
}

func NewTX(tx *sqlx.Tx) db.TX {
	return SQLXTX{tx}
}
