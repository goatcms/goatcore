package orm

import "testing"

func TestCreateTable(t *testing.T) {
	scope, err := newTestScope()
	if err != nil {
		t.Error(err)
		return
	}
	ct, err := NewCreateTable(scope.table, scope.dsql)
	if err != nil {
		t.Error(err)
		return
	}
	if err = ct(scope.tx); err != nil {
		// first create table should be ok
		t.Error(err)
		return
	}
}
