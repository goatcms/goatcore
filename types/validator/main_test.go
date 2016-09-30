package validator

import (
	"reflect"

	"github.com/goatcms/goat-core/types"
	"github.com/goatcms/goat-core/types/abstracttype"
)

const (
	testFieldOne      = "fieldOne"
	testFieldTwo      = "fieldTwo"
	testFieldEmail    = "fieldEmail"
	testSQLType       = "char(300)"
	testSQLTypeKey    = "key"
	testHTMLType      = "text"
	testMaxlen        = "maxlength"
	testUndefinedAttr = "undefinedAttr"
)

type TestObject struct {
	fieldOne   string
	fieldTwo   string
	fieldEmail string
}

type TestCustomType struct {
	abstracttype.MetaType
	abstracttype.StringConverter
	EmptyValidator
}

type TestEmailType struct {
	abstracttype.MetaType
	abstracttype.StringConverter
	EmailValidator
}

func NewTestSingleCustomType() types.SingleCustomType {
	var ptr *string
	return &TestCustomType{
		MetaType: abstracttype.MetaType{
			SQLTypeName:  "varchar(100)",
			HTMLTypeName: "text",
			GoTypeRef:    reflect.TypeOf(ptr).Elem(),
			Attributes:   make(map[string]string),
		},
	}
}

func NewTestSingleEmailType() types.SingleCustomType {
	var ptr *string
	return &TestCustomType{
		MetaType: abstracttype.MetaType{
			SQLTypeName:  "varchar(100)",
			HTMLTypeName: "email",
			GoTypeRef:    reflect.TypeOf(ptr).Elem(),
			Attributes:   make(map[string]string),
		},
	}
}

func NewTestCustomType() types.CustomType {
	return &abstracttype.CustomType{
		SingleCustomType: NewTestSingleCustomType(),
	}
}

func NewTestEmailType() types.CustomType {
	return &abstracttype.CustomType{
		SingleCustomType: NewTestSingleEmailType(),
	}
}

func NewTestObjectCustomType() types.CustomType {
	return &abstracttype.ObjectCustomType{
		SingleCustomType: NewTestSingleCustomType(),
		Types: map[string]types.CustomType{
			testFieldOne:   NewTestCustomType(),
			testFieldTwo:   NewTestCustomType(),
			testFieldEmail: NewTestEmailType(),
		},
	}
}
