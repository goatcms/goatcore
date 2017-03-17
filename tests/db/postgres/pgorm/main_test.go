package pgorm

import (
	"reflect"

	"github.com/goatcms/goatcore/db"
	"github.com/goatcms/goatcore/db/adapter"
	"github.com/goatcms/goatcore/db/dsql/pgDSQL"
	"github.com/goatcms/goatcore/db/orm"
	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/filesystem/filespace/memfs"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	TestTableName = "TestTable"
)

type testScope struct {
	tx    db.TX
	table db.Table
	dsql  db.DSQL
	fs    filesystem.Filespace
}

type TestEntity struct {
	ID      int64  `db:"id" sqltype:"!id"`
	Title   string `db:"title" sqltype:"!char(400)"`
	Content string `db:"content" sqltype:"!char(400)"`
	Image   string `db:"image" sqltype:"!char(400)"`
}

func newTestScope() (*testScope, error) {
	var (
		err    error
		config *TestConfig
		db     *sqlx.DB
	)
	config, err = LoadTestConfig()
	if err != nil {
		return nil, err
	}
	db, err = sqlx.Open("postgres", config.URL)
	if err != nil {
		return nil, err
	}
	fs, err := memfs.NewFilespace()
	if err != nil {
		return nil, err
	}
	var ptr *TestEntity
	table := orm.NewTable(TestTableName, reflect.TypeOf(ptr).Elem())
	return &testScope{
		table: table,
		dsql:  pgDSQL.NewDSQL(),
		tx:    adapter.NewTXFromDB(db),
		fs:    fs,
	}, nil
}

func countResult(rows db.Rows) (int64, error) {
	counter := int64(0)
	for rows.Next() {
		row := TestEntity{}
		err := rows.StructScan(&row)
		if err != nil {
			return -1, err
		}
		counter++
	}
	return counter, nil
}
