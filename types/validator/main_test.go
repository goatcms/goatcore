package validator

import (
	"reflect"

	"github.com/goatcms/goat-core/types"
	"github.com/goatcms/goat-core/types/abstracttype"
)

const (
	testFieldOne      = "FieldOne"
	testFieldTwo      = "FieldTwo"
	testFieldEmail    = "FieldEmail"
	testSQLType       = "char(300)"
	testSQLTypeKey    = "key"
	testHTMLType      = "text"
	testMaxlen        = "maxlength"
	testUndefinedAttr = "undefinedAttr"
)

type TestObject struct {
	FieldOne   string
	FieldTwo   string
	FieldEmail string
}

func NewTestCustomType() types.CustomType {
	var ptr *string
	return &abstracttype.SimpleCustomType{
		MetaType: &abstracttype.MetaType{
			SQLTypeName:  "varchar(100)",
			HTMLTypeName: "text",
			GoTypeRef:    reflect.TypeOf(ptr).Elem(),
			Attributes:   make(map[string]string),
		},
		TypeConverter: abstracttype.NewStringConverter(),
		TypeValidator: NewNoValidator(),
	}
}

func NewTestEmailType() types.CustomType {
	var ptr *string
	return &abstracttype.SimpleCustomType{
		MetaType: &abstracttype.MetaType{
			SQLTypeName:  "varchar(100)",
			HTMLTypeName: "text",
			GoTypeRef:    reflect.TypeOf(ptr).Elem(),
			Attributes:   map[string]string{},
		},
		TypeConverter: abstracttype.NewStringConverter(),
		TypeValidator: NewEmailValidator(),
	}
}

func NewTestLengthType() types.CustomType {
	var ptr *string
	return &abstracttype.SimpleCustomType{
		MetaType: &abstracttype.MetaType{
			SQLTypeName:  "varchar(100)",
			HTMLTypeName: "text",
			GoTypeRef:    reflect.TypeOf(ptr).Elem(),
			Attributes:   map[string]string{},
		},
		TypeConverter: abstracttype.NewStringConverter(),
		TypeValidator: NewLengthValidator(3, 7),
	}
}

func NewTestObjectCustomType() types.CustomType {
	var ptr *TestObject
	goTypeRef := reflect.TypeOf(ptr).Elem()
	types := map[string]types.CustomType{
		testFieldOne:   NewTestCustomType(),
		testFieldTwo:   NewTestCustomType(),
		testFieldEmail: NewTestEmailType(),
	}
	return &abstracttype.ObjectCustomType{
		MetaType: &abstracttype.MetaType{
			SQLTypeName:  "text",
			HTMLTypeName: "text",
			GoTypeRef:    goTypeRef,
			Attributes:   map[string]string{},
		},
		TypeConverter: abstracttype.NewObjectConverterFromType(goTypeRef),
		TypeValidator: NewObjectValidator(types),
		Types:         types,
	}
}
