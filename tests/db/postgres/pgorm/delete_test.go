package pgorm

import (
	"reflect"
	"testing"

	"github.com/goatcms/goatcore/db/orm"
)

func TestDelete(t *testing.T) {
	var ptr *TestEntity
	table := orm.NewTable("DeleteTest", reflect.TypeOf(ptr).Elem())
	scope, err := newTestScope()
	if err != nil {
		t.Error(err)
		return
	}
	dropTable, _ := orm.NewDropTable(table, scope.driver)
	dropTable(scope.tx)
	createTable, _ := orm.NewCreateTable(table, scope.driver)
	if err = createTable(scope.tx); err != nil {
		t.Error(err)
		return
	}
	insert, err := orm.NewInsertWithID(table, scope.driver)
	if err != nil {
		t.Error(err)
		return
	}
	if err = insert(scope.tx, &TestEntity{10, "title1", "content1", "path1"}); err != nil {
		t.Error(err)
		return
	}
	if err = insert(scope.tx, &TestEntity{20, "title2", "content2", "path2"}); err != nil {
		t.Error(err)
		return
	}
	findAll, err := orm.NewFindAll(table, scope.driver)
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
	if counter != 2 {
		t.Errorf("insert record error")
		return
	}
	delete, err := orm.NewDelete(table, scope.driver)
	if err != nil {
		t.Error(err)
		return
	}
	if err = delete(scope.tx, 10); err != nil {
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
		t.Errorf("delete record error")
		return
	}
}
