package itedriver

import (
	"database/sql"
	"fmt"

	"github.com/goatcms/goatcore/db"
	"github.com/goatcms/goatcore/db/dbdriver"
	"github.com/goatcms/goatcore/db/dsql/sqliteDSQL"
	"github.com/goatcms/goatcore/varutil"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type EmptyData struct{}

func init() {
	dbdriver.Register("sqlite3", NewDriver())
}

type Driver struct {
	dsql db.DSQL
}

func NewDriver() db.Driver {
	return &Driver{
		dsql: sqliteDSQL.NewDSQL(),
	}
}

func (d *Driver) DSQL() db.DSQL {
	return d.dsql
}

func (dw *Driver) Open(dataSourceName string) (db.Connection, error) {
	c, err := sqlx.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, err
	}
	return &Conn{
		DB: c,
	}, nil
}

func (dw *Driver) RunSelect(tx db.TX, query string) (db.Rows, error) {
	rows, err := tx.Queryx(query)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", err.Error(), query)
	}
	return rows.(db.Rows), nil
}

func (dw *Driver) RunInsert(tx db.TX, query string, data interface{}) (int64, error) {
	var (
		res sql.Result
		err error
		id  int64
	)
	if res, err = tx.NamedExec(query, data); err != nil {
		return -1, fmt.Errorf("%s: %s", err.Error(), query)
	}
	if id, err = res.LastInsertId(); err != nil {
		return -1, fmt.Errorf("%s: %s", err.Error(), query)
	}
	if err = varutil.SetField(data, "ID", id); err != nil {
		return -1, fmt.Errorf("%s: %s", err.Error(), query)
	}
	return id, nil
}

func (dw *Driver) RunUpdate(tx db.TX, query string, data interface{}) error {
	var (
		res   sql.Result
		err   error
		count int64
	)
	if res, err = tx.NamedExec(query, data); err != nil {
		return fmt.Errorf("%s: %s", err.Error(), query)
	}
	if count, err = res.RowsAffected(); err != nil {
		return fmt.Errorf("%s: %s", err.Error(), query)
	}
	if count != 1 {
		return fmt.Errorf("Update modified more then one record (%v records modyfieds): %s", count, query)
	}
	return nil
}

func (dw *Driver) RunDelete(tx db.TX, query string, data interface{}) error {
	var (
		res   sql.Result
		err   error
		count int64
	)
	if res, err = tx.NamedExec(query, data); err != nil {
		return fmt.Errorf("%s: %s", err.Error(), query)
	}
	if count, err = res.RowsAffected(); err != nil {
		return fmt.Errorf("%s: %s", err.Error(), query)
	}
	if count != 1 {
		return fmt.Errorf("Delete more than one record (%v records deleted)", count)
	}
	return nil
}

func (dw *Driver) RunCreateTable(tx db.TX, query string) error {
	if _, err := tx.NamedExec(query, EmptyData{}); err != nil {
		return err
	}
	return nil
}

func (dw *Driver) RunDropTable(tx db.TX, query string) error {
	if _, err := tx.NamedExec(query, EmptyData{}); err != nil {
		return err
	}
	return nil
}
