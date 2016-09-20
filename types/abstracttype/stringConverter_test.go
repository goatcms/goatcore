package abstracttype

import (
	"testing"

	"github.com/goatcms/goat-core/testbase/mocks/mockmultipart"
)

func TestStringConverter_FromString(t *testing.T) {
	converter := StringConverter{}
	iInterface, err := converter.FromString("11")
	if err != nil {
		t.Error(err)
		return
	}
	iv := iInterface.(string)
	if iv != "11" {
		t.Errorf("From string can not convert 11, result is %v", iv)
		return
	}
}

func TestStringConverter_FromMultipart(t *testing.T) {
	fh := mockmultipart.NewFileHeader([]byte("11"))
	converter := StringConverter{}
	iInterface, err := converter.FromMultipart(fh)
	if err != nil {
		t.Error(err)
		return
	}
	iv := iInterface.(string)
	if iv != "11" {
		t.Errorf("From string can not convert 11, result is %v", iv)
		return
	}
}

func TestStringConverter_ToString(t *testing.T) {
	converter := StringConverter{}
	istr, err := converter.ToString("11")
	if err != nil {
		t.Error(err)
		return
	}
	if istr != "11" {
		t.Errorf("To string can not convert 11, result is %v", istr)
		return
	}
}
