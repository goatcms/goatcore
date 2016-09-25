package abstracttype

import (
	"testing"

	"github.com/goatcms/goat-core/types"
)

func TestBaseMetaType(t *testing.T) {
	bt := MetaType{
		SQLTypeName:  testSQLType,
		HTMLTypeName: testHTMLType,
		Validators:   []types.Validator{},
		Attributes: map[string]string{
			testMaxlen: "10",
		},
	}

	if bt.SQLType() != testSQLType {
		t.Errorf("SQLType() is wrong %v != %v", bt.SQLType(), testSQLType)
	}

	if bt.HTMLType() != testHTMLType {
		t.Errorf("html type is wrong %v", bt.HTMLType())
	}

	if !bt.HasAttribute(testMaxlen) {
		t.Errorf("maxlength should be defined %v", bt.Attributes)
	}
	if bt.GetAttribute(testMaxlen) != "10" {
		t.Errorf("maxlength should be '10' (%v)", bt.Attributes)
	}

	if bt.HasAttribute(testUndefinedAttr) == true {
		t.Errorf("undefinedAttr should be undefined")
	}
	if bt.GetAttribute(testUndefinedAttr) != "" {
		t.Errorf("undefinedAttr should be empty string")
	}
}
