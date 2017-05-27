package itedriver

import (
	"fmt"

	"github.com/goatcms/goatcore/db"
	"github.com/jmoiron/sqlx"
)

type Conn struct {
	*sqlx.DB
}

func (c Conn) Queryx(query string, args ...interface{}) (db.Rows, error) {
	rows, err := c.DB.Queryx(query, args...)
	return db.Rows(rows), err
}

func (c Conn) QueryRowx(query string, args ...interface{}) (db.Row, error) {
	row := c.DB.QueryRowx(query, args...)
	return db.Row(row), row.Err()
}

func (c Conn) Commit() error {
	return nil
}

func (c Conn) Rollback() error {
	return fmt.Errorf("Database connection as transaction don't support rollback (all queries are autorun))")
}

func (c Conn) Begin() (db.TX, error) {
	x, err := c.DB.Beginx()
	if err != nil {
		return nil, err
	}
	return Tx{
		Tx: x,
	}, nil
}
