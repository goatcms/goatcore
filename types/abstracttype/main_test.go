package abstracttype

import (
	"reflect"

	"github.com/goatcms/goat-core/types"
	"github.com/goatcms/goat-core/types/validator"
)

const (
	testFieldOne      = "fieldOne"
	testFieldTwo      = "fieldTwo"
	testSQLType       = "char(300)"
	testSQLTypeKey    = "key"
	testHTMLType      = "text"
	testMaxlen        = "maxlength"
	testUndefinedAttr = "undefinedAttr"
)

type TestSimpleCustomType struct {
	MetaType
	StringConverter
	validator.EmptyValidator
}

func NewTestSingleCustomType() types.SingleCustomType {
	var ptr *string
	return &TestSimpleCustomType{
		MetaType: MetaType{
			SQLTypeName:  "varchar(100)",
			HTMLTypeName: "text",
			GoTypeRef:    reflect.TypeOf(ptr).Elem(),
			Attributes:   make(map[string]string),
		},
	}
}

func NewTestSimpleCustomType() types.CustomType {
	return &SimpleCustomType{
		SingleCustomType: NewTestSingleCustomType(),
	}
}

func NewTestObjectCustomType() types.CustomType {
	return &ObjectCustomType{
		SingleCustomType: NewTestSingleCustomType(),
		Types: map[string]types.CustomType{
			testFieldOne: NewTestSimpleCustomType(),
			testFieldTwo: NewTestSimpleCustomType(),
		},
	}
}
