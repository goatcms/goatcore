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

type TestCustomType struct {
	MetaType
	StringConverter
}

/*
func NewTestSingleCustomType() types.CustomType {
	var ptr *string
	return &TestCustomType{
		MetaType: MetaType{
			SQLTypeName:  "varchar(100)",
			HTMLTypeName: "text",
			GoTypeRef:    reflect.TypeOf(ptr).Elem(),
			Attributes:   make(map[string]string),
		},
	}
}*/

func NewTestCustomType() types.CustomType {
	var ptr *string
	return &SimpleCustomType{
		MetaType: &MetaType{
			SQLTypeName:  "varchar(100)",
			HTMLTypeName: "text",
			GoTypeRef:    reflect.TypeOf(ptr).Elem(),
			Attributes:   make(map[string]string),
		},
		TypeConverter: NewStringConverter(),
		TypeValidator: validator.NewNoValidator(),
	}
}

func NewTestObjectCustomType() types.CustomType {
	var ptr *string
	types := map[string]types.CustomType{
		testFieldOne: NewTestCustomType(),
		testFieldTwo: NewTestCustomType(),
	}
	return &ObjectCustomType{
		MetaType: &MetaType{
			SQLTypeName:  "varchar(100)",
			HTMLTypeName: "text",
			GoTypeRef:    reflect.TypeOf(ptr).Elem(),
			Attributes:   make(map[string]string),
		},
		TypeConverter: NewStringConverter(),
		Types:         types,
		TypeValidator: validator.ObjectValidator{
			Types: types,
		},
	}
}
