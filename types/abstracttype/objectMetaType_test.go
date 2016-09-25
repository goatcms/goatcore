package abstracttype

import (
	"reflect"
	"testing"

	"github.com/goatcms/goat-core/types"
	"github.com/goatcms/goat-core/types/validator"
)

const (
	TestFieldOne = "fieldOne"
	TestFieldTwo = "fieldTwo"
)

// ContentType represent content field type
type TestSimpleType struct {
	MetaType
	StringConverter
	validator.EmptyValidator
}

// NewTestSimpleType create new field type
func NewTestSimpleType(attrs map[string]string) types.CustomType {
	var ptr *string
	return &TestSimpleType{
		MetaType: MetaType{
			SQLTypeName:  "text",
			HTMLTypeName: "text",
			GoTypeRef:    reflect.TypeOf(ptr).Elem(),
			Attributes:   attrs,
		},
	}
}

type TestObjectMetaScope struct {
	t *ObjectMetaType
}

func NewTestObjectMetaScope(attrs map[string]string) *TestObjectMetaScope {
	var ptr *ObjectMetaType
	return &TestObjectMetaScope{
		t: &ObjectMetaType{
			Types: map[string]types.CustomType{
				TestFieldOne: NewTestSimpleType(map[string]string{}),
				TestFieldTwo: NewTestSimpleType(map[string]string{}),
			},
			HTMLTypeName: "text",
			GoTypeRef:    reflect.TypeOf(ptr).Elem(),
			Attributes:   attrs,
		},
	}
}

func TestObjectMeta_AddSQLType(t *testing.T) {
	scope := NewTestObjectMetaScope(map[string]string{})
	m := make(map[string]string)
	scope.t.AddSQLType("t", m)
	if _, ok := m["t_"+TestFieldOne]; !ok {
		t.Errorf("t_fieldOne should be defined %v", m)
	}
	if _, ok := m["t_"+TestFieldTwo]; !ok {
		t.Errorf("t_fieldTwo should be defined %v", m)
	}
	if _, ok := m["undefinedProp"]; ok {
		t.Errorf("undefinedProp should be undefined")
	}
}

func TestObjectMeta_GetSQLType(t *testing.T) {
	scope := NewTestObjectMetaScope(map[string]string{})
	m := scope.t.GetSQLType()
	if _, ok := m[TestFieldOne]; !ok {
		t.Errorf("fieldOne should be defined %v", m)
	}
	if _, ok := m[TestFieldTwo]; !ok {
		t.Errorf("fieldTwo should be defined %v", m)
	}
	if _, ok := m[undefinedAttr]; ok {
		t.Errorf("undefinedProp should be undefined")
	}
}

func TestObjectBaseMetaType_data(t *testing.T) {
	scope := NewTestObjectMetaScope(map[string]string{
		maxlen: "10",
	})

	if scope.t.HTMLType() != htmlType {
		t.Errorf("html type is wrong %v", scope.t.HTMLType())
	}

	if !scope.t.HasAttribute(maxlen) {
		t.Errorf("maxlength should be defined %v", scope.t.Attributes)
	}
	if scope.t.GetAttribute(maxlen) != "10" {
		t.Errorf("maxlength should be '10' (%v)", scope.t.Attributes)
	}

	if scope.t.HasAttribute(undefinedAttr) == true {
		t.Errorf("undefinedAttr should be undefined")
	}
	if scope.t.GetAttribute(undefinedAttr) != "" {
		t.Errorf("undefinedAttr should be empty string")
	}
}
