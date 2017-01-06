package db

import (
	"database/sql"

	"github.com/goatcms/goat-core/types"
)

const (
	// DSQLService is a key to access DSQL service
	DSQLService = "DSQL"
)

type FindAll func(TX) (Rows, error)
type FindByID func(TX, int64) (Row, error)
type Insert func(TX, interface{}) (int64, error)
type InsertWithID func(TX, interface{}) error
type Update func(TX, interface{}) error
type Delete func(TX, int64) error
type CreateTable func(TX) error

// DSQL is dynamic sql generator (build sql queries)
type DSQL interface {
	NewSelectSQL(table string, fields []string) (string, error)
	NewSelectWhereSQL(table string, fields []string, where string) (string, error)
	NewInsertSQL(table string, fields []string) (string, error)
	NewUpdateSQL(table string, fields []string) (string, error)
	NewUpdateWhereSQL(table string, fields []string, where string) (string, error)
	NewDeleteSQL(table string) (string, error)
	NewDeleteWhereSQL(table string, where string) (string, error)
	NewSQLType(t types.CustomType) (string, error)
	NewCreateSQL(table string, types map[string]types.CustomType) (string, error)
}

// Table sd
type Table interface {
	Name() string
	Fields() []string
	Types() map[string]types.CustomType
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

type Row interface {
	Scan(...interface{}) error
	StructScan(interface{}) error
	Columns() ([]string, error)
	Err() error
}
