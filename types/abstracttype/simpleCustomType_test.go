package abstracttype

import (
	"testing"

	"github.com/goatcms/goat-core/types"
)

func TestSimpleCustomType_AddSQLType(t *testing.T) {
	objectCustomType := NewTestObjectCustomType()
	m := make(map[string]types.CustomType)
	objectCustomType.AddSubTypes("t", m)
	if _, ok := m["t_"+testFieldOne]; !ok {
		t.Errorf("t should be defined %v", m)
	}
	if _, ok := m["undefinedProp"]; ok {
		t.Errorf("undefinedProp should be undefined")
	}
	if _, ok := m["t_undefinedProp"]; ok {
		t.Errorf("t_undefinedProp should be undefined")
	}
}

func TestSimpleCustomType_GetSQLType(t *testing.T) {
	objectCustomType := NewTestObjectCustomType()
	m := objectCustomType.GetSubTypes()
	if _, ok := m[testFieldOne]; !ok {
		t.Errorf("fieldOne should be defined %v", m)
	}
	if _, ok := m[testFieldTwo]; !ok {
		t.Errorf("fieldTwo should be defined %v", m)
	}
	if _, ok := m[testUndefinedAttr]; ok {
		t.Errorf("undefinedProp should be undefined")
	}
}
