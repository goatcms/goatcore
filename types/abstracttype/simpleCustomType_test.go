package abstracttype

import (
	"testing"

	"github.com/goatcms/goat-core/types"
)

func TestCustomType_AddSQLType(t *testing.T) {
	objectCustomType := NewTestCustomType()
	m := make(map[string]types.CustomType)
	objectCustomType.AddSubTypes("t", m)
	if _, ok := m["t"]; !ok {
		t.Errorf("t should be defined %v", m)
	}
	if _, ok := m["undefinedProp"]; ok {
		t.Errorf("undefinedProp should be undefined")
	}
	if _, ok := m["t_undefinedProp"]; ok {
		t.Errorf("t_undefinedProp should be undefined")
	}
}

func TestCustomType_GetSQLType(t *testing.T) {
	objectCustomType := NewTestCustomType()
	m := objectCustomType.GetSubTypes()
	if _, ok := m[""]; !ok {
		t.Errorf("empty string should contains current type")
	}
	if _, ok := m[testFieldOne]; ok {
		t.Errorf("fieldOne should be undefined %v", m)
	}
	if _, ok := m[testFieldTwo]; ok {
		t.Errorf("fieldTwo should be undefined %v", m)
	}
	if _, ok := m[testUndefinedAttr]; ok {
		t.Errorf("undefinedProp should be undefined")
	}
}
