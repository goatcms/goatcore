package validator

import (
	"testing"

	"github.com/goatcms/goatcore/messages/msgcollection"
)

func TestLengthValidator_Min_Pass(t *testing.T) {
	messagesMap := msgcollection.NewMessageMap()
	str := "aa"
	if err := MinStringValid(str, "", messagesMap, 2); err != nil {
		t.Error(err)
		return
	}
	if len(messagesMap.GetAll()) != 0 {
		t.Errorf("Validation return a error (expocted no error): %v", messagesMap.GetAll())
		return
	}
}

func TestLengthValidator_Min_Fail(t *testing.T) {
	messagesMap := msgcollection.NewMessageMap()
	str := "a"
	if err := MinStringValid(str, "", messagesMap, 2); err != nil {
		t.Error(err)
		return
	}
	if len(messagesMap.GetAll()) == 0 {
		t.Errorf("Validation don't return a error (expocted error): %v", messagesMap.GetAll())
		return
	}
}
