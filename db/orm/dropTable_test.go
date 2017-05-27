package orm

import "testing"

func TestDropTable(t *testing.T) {
	scope, err := newTestScope()
	if err != nil {
		t.Error(err)
		return
	}
	ct, err := NewCreateTable(scope.table, scope.driver)
	if err != nil {
		t.Error(err)
		return
	}
	if err = ct(scope.tx); err != nil {
		t.Error(err)
		return
	}
	ctd, err := NewDropTable(scope.table, scope.driver)
	if err != nil {
		t.Error(err)
		return
	}
	if err = ctd(scope.tx); err != nil {
		t.Error(err)
		return
	}
}
