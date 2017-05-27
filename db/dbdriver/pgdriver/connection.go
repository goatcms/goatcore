package pgdriver

import (
	"fmt"
	"time"

	"github.com/goatcms/goatcore/db"
	"github.com/goatcms/goatcore/db/dbconfig"
	"github.com/jackc/pgx"
)

// Conn is a database connection
type Conn struct {
	pool *pgx.ConnPool
}

// NewConnection create a new database conncetion
func NewConnection(source string) (db.Connection, error) {
	var pool *pgx.ConnPool
	configDecoder, err := dbconfig.NewDecoderFromKeyValueString(source)
	if err != nil {
		return nil, err
	}
	pool, err = pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:              configDecoder.Get("host", "localhost"),
			Port:              uint16(configDecoder.GetInt("port", 5432)),
			Database:          configDecoder.Get("dataabse", ""),
			User:              configDecoder.Get("user", ""),
			Password:          configDecoder.Get("password", ""),
			TLSConfig:         nil,
			UseFallbackTLS:    false,
			FallbackTLSConfig: nil,
			Logger:            nil,
			LogLevel:          configDecoder.GetInt("loglevel", 0),
			Dial:              nil,
			RuntimeParams:     make(map[string]string),
		},
		MaxConnections: configDecoder.GetInt("maxConnections", 20),
		AfterConnect:   nil,
		AcquireTimeout: time.Duration(configDecoder.GetInt64("acquireTimeout", 2)) * time.Second,
	})
	if err != nil {
		return nil, err
	}
	return &Conn{
		pool: pool,
	}, nil
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
