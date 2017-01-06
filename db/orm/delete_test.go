package orm

import "testing"

func TestDelete(t *testing.T) {
	scope, err := newTestScope()
	if err != nil {
		t.Error(err)
		return
	}
	createTable, err := NewCreateTable(scope.table, scope.dsql)
	if err != nil {
		t.Error(err)
		return
	}
	if err = createTable(scope.tx); err != nil {
		t.Error(err)
		return
	}
	insert, err := NewInsert(scope.table, scope.dsql)
	if err != nil {
		t.Error(err)
		return
	}
	rowID, err := insert(scope.tx, &TestEntity{10, "title1", "content1", "path1"})
	if err != nil {
		t.Error(err)
		return
	}
	_, err = insert(scope.tx, &TestEntity{0, "title2", "content2", "path2"})
	if err != nil {
		t.Error(err)
		return
	}
	findAll, err := NewFindAll(scope.table, scope.dsql)
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
	delete, err := NewDelete(scope.table, scope.dsql)
	if err != nil {
		t.Error(err)
		return
	}
	if err = delete(scope.tx, rowID); err != nil {
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
