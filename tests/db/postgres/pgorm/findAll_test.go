package pgorm

import (
	"reflect"
	"testing"

	"github.com/goatcms/goatcore/db/orm"
)

func TestFindAll(t *testing.T) {
	var ptr *TestEntity
	table := orm.NewTable("FindAllTest", reflect.TypeOf(ptr).Elem())
	scope, err := newTestScope()
	if err != nil {
		t.Error(err)
		return
	}
	dropTable, _ := orm.NewDropTable(table, scope.dsql)
	dropTable(scope.tx)
	createTable, _ := orm.NewCreateTable(table, scope.dsql)
	if err = createTable(scope.tx); err != nil {
		t.Error(err)
		return
	}
	findAll, err := orm.NewFindAll(table, scope.dsql)
	if err != nil {
		t.Error(err)
		return
	}
	rows, err := findAll(scope.tx)
	if err != nil {
		t.Error(err)
		return
	}
	counter, err := countResult(rows)
	if err != nil {
		t.Error(err)
		return
	}
	if counter != 0 {
		t.Errorf("(during start) we have 0 record")
		return
	}
	insert, err := orm.NewInsert(table, scope.dsql)
	if err != nil {
		t.Error(err)
		return
	}
	_, err = insert(scope.tx, &TestEntity{10, "title1", "content1", "path1"})
	if err != nil {
		t.Error(err)
		return
	}
	rows, err = findAll(scope.tx)
	if err != nil {
		t.Error(err)
		return
	}
	counter, err = countResult(rows)
	if err != nil {
		t.Error(err)
		return
	}
	if counter != 1 {
		t.Errorf("after insert should return one record")
		return
	}
}
