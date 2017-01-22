package orm

import (
	"reflect"

	"github.com/goatcms/goat-core/db"
	"github.com/goatcms/goat-core/db/adapter"
	"github.com/goatcms/goat-core/db/dsql"
	"github.com/goatcms/goat-core/filesystem"
	"github.com/goatcms/goat-core/filesystem/filespace/memfs"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
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
	ID      int64  `db:"id" sql:"INTEGER PRIMARY KEY"`
	Title   string `db:"title" sql:"VARCHAR(400)"`
	Content string `db:"content" sql:"VARCHAR(400)"`
	Image   string `db:"image" sql:"VARCHAR(400)"`
}

func newTestScope() (*testScope, error) {
	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}
	fs, err := memfs.NewFilespace()
	if err != nil {
		return nil, err
	}
	var ptr *TestEntity
	table := NewTable(TestTableName, reflect.TypeOf(ptr).Elem())
	return &testScope{
		table: table,
		dsql:  dsql.NewDSQL(),
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
