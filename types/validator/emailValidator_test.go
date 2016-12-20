package validator

import "testing"

func TestEmailValidator_Pass(t *testing.T) {
	customType := NewTestEmailType()
	result, err := customType.Valid("myown@email.address")
	if err != nil {
		t.Error(err)
		return
	}
	if len(result.GetAll()) != 0 {
		t.Errorf("Validation return a error (expocted no error): %v", result.GetAll())
		return
	}
}

func TestEmailValidator_Fail(t *testing.T) {
	customType := NewTestEmailType()
	result, err := customType.Valid("sdasdsadasd")
	if err != nil {
		t.Error(err)
		return
	}
	if len(result.GetAll()) == 0 {
		t.Errorf("Validation should return a error: %v", result.GetAll())
		return
	}
}
