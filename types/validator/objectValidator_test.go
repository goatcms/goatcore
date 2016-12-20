package validator

import "testing"

func TestObjectValidator_Pass(t *testing.T) {
	customType := NewTestObjectCustomType()
	o := &TestObject{
		FieldOne:   "one",
		FieldTwo:   "two",
		FieldEmail: "email@internet.pl",
	}
	result, err := customType.Valid(o)
	if err != nil {
		t.Error(err)
		return
	}
	if len(result.GetAll()) != 0 {
		t.Errorf("Validation return a error (expocted no error): %v", result.GetAll())
		return
	}
}

func TestObjectValidator_Fail(t *testing.T) {
	customType := NewTestObjectCustomType()
	o := &TestObject{
		FieldOne:   "one",
		FieldTwo:   "two",
		FieldEmail: "blablablabla",
	}
	result, err := customType.Valid(o)
	if err != nil {
		t.Error(err)
		return
	}
	if len(result.GetAll()) == 0 {
		t.Errorf("Validation should return a error: %v", result.GetAll())
		return
	}
}
