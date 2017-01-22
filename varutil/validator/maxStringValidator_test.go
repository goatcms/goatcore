package validator

import (
	"testing"

	"github.com/goatcms/goat-core/messages/msgcollection"
)

func TestLengthValidator_Max_Pass(t *testing.T) {
	messagesMap := msgcollection.NewMessageMap()
	str := "a"
	if err := MaxStringValid(str, "", messagesMap, 1); err != nil {
		t.Error(err)
		return
	}
	if len(messagesMap.GetAll()) != 0 {
		t.Errorf("Validation return a error (expocted no error): %v", messagesMap.GetAll())
		return
	}
}

func TestLengthValidator_Max_Fail(t *testing.T) {
	messagesMap := msgcollection.NewMessageMap()
	str := "aa"
	if err := MaxStringValid(str, "", messagesMap, 1); err != nil {
		t.Error(err)
		return
	}
	if len(messagesMap.GetAll()) == 0 {
		t.Errorf("Validation don't return a error (expocted error): %v", messagesMap.GetAll())
		return
	}
}
