package abstracttype

import (
	"testing"

	"github.com/goatcms/goat-core/types"
)

func TestObjectCustomType_AddSQLType(t *testing.T) {
	objectCustomType := NewTestObjectCustomType()
	m := make(map[string]types.CustomType)
	objectCustomType.AddSubTypes("t", m)
	if _, ok := m["t_"+testFieldOne]; !ok {
		t.Errorf("t_fieldOne should be defined %v", m)
	}
	if _, ok := m["t_"+testFieldTwo]; !ok {
		t.Errorf("t_fieldTwo should be defined %v", m)
	}
	if _, ok := m["undefinedProp"]; ok {
		t.Errorf("undefinedProp should be undefined")
	}
}

func TestObjectCustomType_GetSQLType(t *testing.T) {
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
