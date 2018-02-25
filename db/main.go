// Package db provide database access interface
// Deprecated: Use generate orm with goatcli
package db

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

const (
	// SQLTypeTagName is a tag name for database injection
	SQLTypeTagName = "sqltype"
)

type FindAll func(TX) (Rows, error)
type FindByID func(TX, int64) (Row, error)
type Insert func(TX, interface{}) (int64, error)
type InsertWithID func(TX, interface{}) error
type Update func(TX, interface{}) error

// Delete remove a row
type Delete func(TX, int64) error

// CreateTable is a create table function
type CreateTable func(TX) error

// DropTable is a drop table function
type DropTable func(TX) error

// Database represent database
type Database struct {
	DSQL     DSQL
	instance *sqlx.DB
}

// Table sd
type Table interface {
	Name() string
	Fields() []string
	Types() map[string]string
}

// DSQL is interface for SQL generator
type DSQL interface {
	NewSelectSQL(table string, fields []string) (string, error)
	NewSelectWhereSQL(table string, fields []string, where string) (string, error)
	NewInsertSQL(table string, fields []string) (string, error)
	NewUpdateSQL(table string, fields []string) (string, error)
	NewUpdateWhereSQL(table string, fields []string, where string) (string, error)
	NewDeleteSQL(table string) (string, error)
	NewDeleteWhereSQL(table string, where string) (string, error)
	NewCreateSQL(table string, types map[string]string) (string, error)
	NewDropTableSQL(table string) (string, error)
}

// ORM is interface for simple orm
type ORM interface {
	NewSelectSQL(table string, fields []string) (string, error)
	NewSelectWhereSQL(table string, fields []string, where string) (string, error)
	NewInsertSQL(table string, fields []string) (string, error)
	NewUpdateSQL(table string, fields []string) (string, error)
	NewUpdateWhereSQL(table string, fields []string, where string) (string, error)
	NewDeleteSQL(table string) (string, error)
	NewDeleteWhereSQL(table string, where string) (string, error)
	NewCreateSQL(table string, types map[string]string) (string, error)
	NewDropTableSQL(table string) (string, error)
}

// TX represent a database transaction accessor
type TX interface {
	Queryx(query string, args ...interface{}) (Rows, error)
	QueryRowx(query string, args ...interface{}) (Row, error)
	NamedExec(query string, arg interface{}) (sql.Result, error)
	MustExec(query string, args ...interface{}) sql.Result
	Commit() error
	Rollback() error
}

// Rows represent a query response
type Rows interface {
	Close() error
	Next() bool
	StructScan(dest interface{}) error
}

// Rows represent a single row query response
type Row interface {
	Scan(...interface{}) error
	StructScan(interface{}) error
	Columns() ([]string, error)
	Err() error
}
