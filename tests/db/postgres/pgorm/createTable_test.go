package pgorm

import (
	"reflect"
	"testing"

	"github.com/goatcms/goatcore/db/orm"
)

func TestCreateTable(t *testing.T) {
	var ptr *TestEntity
	table := orm.NewTable("CreateTableTest", reflect.TypeOf(ptr).Elem())
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
}
