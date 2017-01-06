package orm

import "testing"

func TestFindAll(t *testing.T) {
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
	if counter != 0 {
		t.Errorf("(during start) we have 0 record")
		return
	}
	insert, err := NewInsert(scope.table, scope.dsql)
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
