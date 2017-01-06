package orm

import (
	"github.com/goatcms/goat-core/db"
	"github.com/goatcms/goat-core/db/adapter"
	"github.com/goatcms/goat-core/db/dsql"
	"github.com/goatcms/goat-core/filesystem"
	"github.com/goatcms/goat-core/filesystem/filespace/memfs"
	"github.com/goatcms/goat-core/types"
	"github.com/goatcms/goat-core/types/simpletype"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

const (
	TestTableName = "TestTable"
)

func NewTestTypes(fs filesystem.Filespace) map[string]types.CustomType {
	return map[string]types.CustomType{
		"title":   simpletype.NewTitleType(map[string]string{types.Required: "true"}),
		"content": simpletype.NewContentType(map[string]string{}),
		"image":   simpletype.NewImageType(map[string]string{types.Required: "true"}, fs),
	}
}

type testScope struct {
	tx    db.TX
	table db.Table
	dsql  db.DSQL
	fs    filesystem.Filespace
}

type TestEntity struct {
	ID      int64  `db:"id"`
	Title   string `db:"title"`
	Content string `db:"content"`
	Image   string `db:"image"`
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
	table := NewTable(TestTableName, NewTestTypes(fs))
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
