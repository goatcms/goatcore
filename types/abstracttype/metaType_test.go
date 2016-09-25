package abstracttype

import (
	"testing"

	"github.com/goatcms/goat-core/types"
)

const (
	sqlType       = "char(300)"
	sqlTypeKey    = "key"
	htmlType      = "text"
	maxlen        = "maxlength"
	undefinedAttr = "undefinedAttr"
)

func TestBaseMetaType(t *testing.T) {
	bt := MetaType{
		SQLTypeName:  sqlType,
		HTMLTypeName: htmlType,
		Validators:   []types.Validator{},
		Attributes: map[string]string{
			maxlen: "10",
		},
	}

	m := map[string]string{}
	bt.AddSQLType(sqlTypeKey, m)
	if m[sqlTypeKey] != sqlType {
		t.Errorf("sql type is wrong %v != %v", m[sqlTypeKey], sqlType)
	}

	m = bt.GetSQLType()
	if m[MainElement] != sqlType {
		t.Errorf("bt.GetSQLType().MainElement is wrong %v != %v", m[sqlTypeKey], sqlType)
	}

	if bt.HTMLType() != htmlType {
		t.Errorf("html type is wrong %v", bt.HTMLType())
	}

	if !bt.HasAttribute(maxlen) {
		t.Errorf("maxlength should be defined %v", bt.Attributes)
	}
	if bt.GetAttribute(maxlen) != "10" {
		t.Errorf("maxlength should be '10' (%v)", bt.Attributes)
	}

	if bt.HasAttribute(undefinedAttr) == true {
		t.Errorf("undefinedAttr should be undefined")
	}
	if bt.GetAttribute(undefinedAttr) != "" {
		t.Errorf("undefinedAttr should be empty string")
	}
}
