package pgdriver

import (
	"database/sql"
	"fmt"

	"github.com/goatcms/goatcore/db"
	"github.com/goatcms/goatcore/db/dbdriver"
	"github.com/goatcms/goatcore/db/dsql/pgDSQL"
	"github.com/goatcms/goatcore/varutil"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type EmptyData struct{}

func init() {
	dbdriver.Register("postgres", NewDriver())
}

type Driver struct {
	dsql db.DSQL
}

func NewDriver() db.Driver {
	return &Driver{
		dsql: pgDSQL.NewDSQL(),
	}
}

func (d *Driver) DSQL() db.DSQL {
	return d.dsql
}

func (dw *Driver) Open(dataSourceName string) (db.Connection, error) {
	c, err := sqlx.Open("postgres", dataSourceName)
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
		res db.Row
		err error
		id  int64
	)
	if res, err = tx.QueryRowx(query, data); err != nil {
		return -1, fmt.Errorf("%s: %s", err.Error(), query)
	}
	if err = res.Scan(&id); err != nil {
		return -1, fmt.Errorf("%s: %s", err.Error(), query)
	}
	if id == 0 {
		return -1, fmt.Errorf("id can not be equals to zero: %s", query)
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
