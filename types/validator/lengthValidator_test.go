package validator

import "testing"

func TestLengthValidator_Min_Pass(t *testing.T) {
	customType := NewTestLengthType()
	email := "12345"
	result, err := customType.Valid(email)
	if err != nil {
		t.Error(err)
		return
	}
	if len(result.GetAll()) != 0 {
		t.Errorf("Validation return a error (expocted no error): %v", result.GetAll())
		return
	}
}

func TestLengthValidator_Min_Fail(t *testing.T) {
	customType := NewTestLengthType()
	email := "1"
	result, err := customType.Valid(email)
	if err != nil {
		t.Error(err)
		return
	}
	if len(result.GetAll()) == 0 {
		t.Errorf("Validation should return a error: %v", result.GetAll())
		return
	}
}
