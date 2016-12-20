package orm

import (
	"testing"

	"github.com/goatcms/goat-core/filesystem"
	"github.com/goatcms/goat-core/filesystem/filespace/memfs"
	"github.com/goatcms/goat-core/types"
	"github.com/goatcms/goat-core/types/simpletype"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

const (
	testTable = "testtable"
)

func NewTestTypes(fs filesystem.Filespace) map[string]types.CustomType {
	return map[string]types.CustomType{
		"title":   simpletype.NewTitleType(map[string]string{types.Required: "true"}),
		"content": simpletype.NewContentType(map[string]string{}),
		"image":   simpletype.NewImageType(map[string]string{types.Required: "true"}, fs),
	}
}

type testScope struct {
	bdtx  *sqlx.DB
	dao   *BaseDAO
	table *BaseTable
}

type TestEntity struct {
	ID      int64  `db:"id"`
	Title   string `db:"title"`
	Content string `db:"content"`
	Image   string `db:"image"`
}

func newTestScope() (*testScope, error) {
	bdtx, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}
	fs, err := memfs.NewFilespace()
	if err != nil {
		return nil, err
	}
	types := NewTestTypes(fs)
	table := NewBaseTable(testTable, types)
	dao := NewBaseDAO(table)
	dao.CreateTable(bdtx)
	if err != nil {
		return nil, err
	}
	return &testScope{
		table: table,
		dao:   dao,
		bdtx:  bdtx,
	}, nil
}

func lodaFixtures(s *testScope) error {
	if _, err := s.bdtx.NamedExec(s.table.insertWithIDSQL, &TestEntity{10, "title1", "content1", "path1"}); err != nil {
		return err
	}
	if _, err := s.bdtx.NamedExec(s.table.insertWithIDSQL, TestEntity{20, "title2", "content2", "path2"}); err != nil {
		return err
	}
	return nil
}

func loadResult(rows *sqlx.Rows) ([]TestEntity, error) {
	result := []TestEntity{}
	for rows.Next() {
		row := TestEntity{}
		err := rows.StructScan(&row)
		if err != nil {
			return nil, err
		}
		result = append(result, row)
	}
	return result, nil
}

/*func compareResult(t *testing.T, context *testScope, expectetResult []TestEntity) {

	compareSlice(t, result, expectetResult)
}*/

func compareSlice(t *testing.T, result, expectetResult []TestEntity) {
	if len(result) != len(expectetResult) {
		t.Errorf("table contains %v records (findAll return %v records)", len(expectetResult), len(result))
		return
	}
	for _, row := range expectetResult {
		finded := false
		for i, resRow := range result {
			if compareEntity(row, resRow) {
				result = append(result[:i], result[i+1:]...)
				finded = true
				break
			}
		}
		if finded == false {
			t.Errorf("A result doesn't contain a expected record %v (in result: %v)", row, result)
			return
		}
	}
	if len(result) != 0 {
		t.Errorf("A result contains more records than expected %v", result)
		return
	}
}

func compareEntity(row TestEntity, expectet TestEntity) bool {
	return row.ID == expectet.ID &&
		row.Title == expectet.Title &&
		row.Content == expectet.Content &&
		row.Image == expectet.Image
}
