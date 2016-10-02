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

type TestLengthType struct {
	abstracttype.MetaType
	abstracttype.StringConverter
	LengthValidator
}

type TestObjectType struct {
	abstracttype.ObjectCustomType
	ObjectValidator
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
	return &TestEmailType{
		MetaType: abstracttype.MetaType{
			SQLTypeName:  "varchar(100)",
			HTMLTypeName: "email",
			GoTypeRef:    reflect.TypeOf(ptr).Elem(),
			Attributes:   make(map[string]string),
		},
	}
}

func NewTestSingleLengthType() types.SingleCustomType {
	var ptr *string
	return &TestLengthType{
		MetaType: abstracttype.MetaType{
			SQLTypeName:  "varchar(100)",
			HTMLTypeName: "email",
			GoTypeRef:    reflect.TypeOf(ptr).Elem(),
			Attributes:   make(map[string]string),
		},
		LengthValidator: NewLengthValidator(3, 7),
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

func NewTestLengthType() types.CustomType {
	return &abstracttype.CustomType{
		SingleCustomType: NewTestSingleLengthType(),
	}
}

func NewTestObjectCustomType() types.CustomType {
	types := map[string]types.CustomType{
		testFieldOne:   NewTestCustomType(),
		testFieldTwo:   NewTestCustomType(),
		testFieldEmail: NewTestEmailType(),
	}
	return &TestObjectType{
		ObjectCustomType: abstracttype.ObjectCustomType{
			SingleCustomType: NewTestSingleCustomType(),
			Types:            types,
		},
		ObjectValidator: ObjectValidator{
			Types: types,
		},
	}
}
