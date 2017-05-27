package orm

import (
	"testing"

	"github.com/goatcms/goatcore/db/dbdriver"
	_ "github.com/goatcms/goatcore/db/dbdriver/itedriver"
)

func TestCreateTable(t *testing.T) {
	scope, err := newTestScope()
	if err != nil {
		t.Error(err)
		return
	}
	driver, err := dbdriver.Driver("sqlite3")
	if err != nil {
		t.Error(err)
		return
	}
	ct, err := NewCreateTable(scope.table, driver)
	if err != nil {
		t.Error(err)
		return
	}
	if err = ct(scope.tx); err != nil {
		t.Error(err)
		return
	}
}
