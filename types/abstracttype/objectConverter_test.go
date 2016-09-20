package abstracttype

import (
	"strings"
	"testing"

	"github.com/goatcms/goat-core/testbase/mocks/mockmultipart"
)

type TestObject struct {
	A string
}

// NewTestObject create new instance of test object
func NewTestObject() interface{} {
	return &TestObject{}
}

const (
	TestObjectJSON = "{\"A\":\"value\"}"
)

func TestObjectConverter_FromString(t *testing.T) {
	converter := NewObjectConverter(NewTestObject)
	iInterface, err := converter.FromString(TestObjectJSON)
	if err != nil {
		t.Error(err)
		return
	}
	o := iInterface.(*TestObject)
	if o.A != "value" {
		t.Errorf("A should be 'value', result is %v", o)
		return
	}
}

func TestObjectConverter_FromMultipart(t *testing.T) {
	fh := mockmultipart.NewFileHeader([]byte(TestObjectJSON))
	converter := NewObjectConverter(NewTestObject)
	iInterface, err := converter.FromMultipart(fh)
	if err != nil {
		t.Error(err)
		return
	}
	o := iInterface.(*TestObject)
	if o.A != "value" {
		t.Errorf("A should be 'value', result is %v", o)
		return
	}
}

func TestObjectConverter_ToString(t *testing.T) {
	converter := NewObjectConverter(NewTestObject)
	o := &TestObject{
		A: "value",
	}
	istr, err := converter.ToString(o)
	if err != nil {
		t.Error(err)
		return
	}
	istr = strings.Replace(istr, " ", "", -1)
	istr = strings.Replace(istr, "\t", "", -1)
	istr = strings.Replace(istr, "\n", "", -1)
	if istr != TestObjectJSON {
		t.Errorf("To string can not convert to correct json %v, result is %v", TestObjectJSON, istr)
		return
	}
}
