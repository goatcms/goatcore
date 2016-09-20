package abstracttype

import (
	"testing"

	"github.com/goatcms/goat-core/types"
)

const (
	sqlType       = "char(300)"
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

	if bt.SQLType() != sqlType {
		t.Errorf("sql type is wrong %v", bt.SQLType())
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
