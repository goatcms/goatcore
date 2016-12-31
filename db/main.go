package db

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// DAO is simple orm data access object
type DAO interface {
	FindAll(TX) (*sqlx.Rows, error)
	FindByID(TX, int64) *sqlx.Row
	Insert(TX, interface{}) (int64, error)
	Update(TX, interface{}) error
	Delete(TX, int64) error
	CreateTable(TX) error
}

// TX represent a database transaction accessor
type TX interface {
	Queryx(query string, args ...interface{}) (*sqlx.Rows, error)
	QueryRowx(query string, args ...interface{}) *sqlx.Row
	NamedExec(query string, arg interface{}) (sql.Result, error)
	MustExec(query string, args ...interface{}) sql.Result
}

// Rows represent a query response
type Rows interface {
	Close() error
	Next() bool
	StructScan(dest interface{}) error
}

/*package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/goatcms/cms-core/app/config"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"strings"
)

const (
	DbName = "db"
)

func DbFactory(dp dependency.Provider) (interface{}, error) {
	ins := dp.Get(config.ConfigName)
	config := ins.(config.Config)
	dbConfig := config.Database
	switch strings.ToLower(dbConfig.Adapter) {
	case "mysql":
		return CreateMyslq(dbConfig)
	case "pgsql":
		return CreatePgsql(dbConfig)
	case "sqlite":
		return CreateSqlite(dbConfig)
	default:
		return fmt.Errorf("adapter no supported")
	}
}

func CreateMyslq(dbConfig *config.Database) (interface{}, error) {
	return sql.Open("mysql", config.Source)
}

func CreatePgsql(dbConfig *config.Database) (interface{}, error) {
	return sql.Open("postgres", config.Source)
}

func CreateSqlite(dbConfig *config.Database) (interface{}, error) {
	return sql.Open("sqlite3", config.Source)
}*/
