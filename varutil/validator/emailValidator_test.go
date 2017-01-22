package validator

import (
	"testing"

	"github.com/goatcms/goat-core/messages/msgcollection"
)

func TestEmailValidator_Pass(t *testing.T) {
	messagesMap := msgcollection.NewMessageMap()
	if err := EmailValid("myown@email.address", "", messagesMap); err != nil {
		t.Error(err)
		return
	}
	if len(messagesMap.GetAll()) != 0 {
		t.Errorf("Validation return a error (expocted no error): %v", messagesMap.GetAll())
		return
	}
}

func TestEmailValidator_Fail(t *testing.T) {
	messagesMap := msgcollection.NewMessageMap()
	if err := EmailValid("sdasdsadasd", "", messagesMap); err != nil {
		t.Error(err)
		return
	}
	if len(messagesMap.GetAll()) == 0 {
		t.Errorf("Validation should return a error: %v", messagesMap.GetAll())
		return
	}
}
