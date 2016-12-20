package abstracttype

import (
	"testing"

	"github.com/goatcms/goat-core/testbase/mocks/mockmultipart"
)

func TestInt64Converter_FromString(t *testing.T) {
	ic := Int64Converter{}
	iInterface, err := ic.FromString("11")
	if err != nil {
		t.Error(err)
		return
	}
	iv := iInterface.(int64)
	if iv != 11 {
		t.Errorf("From string can not convert 11, result is %v", iv)
		return
	}
}

func TestInt64Converter_FromMultipart(t *testing.T) {
	fh := mockmultipart.NewFileHeader([]byte("11"))
	ic := Int64Converter{}
	iInterface, err := ic.FromMultipart(fh)
	if err != nil {
		t.Error(err)
		return
	}
	iv, ok := iInterface.(int64)
	if !ok {
		t.Errorf("can not convert to int64: %v", iv)
		return
	}
	if iv != 11 {
		t.Errorf("From multipart can not convert 11, result is %v", iv)
		return
	}
}

func TestInt64Converter_ToString(t *testing.T) {
	//oString(interface{})(string, error)
	ic := Int64Converter{}
	istr, err := ic.ToString(int64(11))
	if err != nil {
		t.Error(err)
		return
	}
	if istr != "11" {
		t.Errorf("To string can not convert 11, result is %v", istr)
		return
	}
}
