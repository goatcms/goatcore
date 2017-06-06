package r

import (
	"reflect"
	"testing"
)

const (
	someString = "someString"
)

type TestType struct {
	StringValue string
}

func TestSetValueFromString_String(t *testing.T) {
	t.Parallel()
	var value string
	if err := SetValueFromString(reflect.ValueOf(&value).Elem(), "testString"); err != nil {
		t.Error(err)
		return
	}
	if value != "testString" {
		t.Errorf("Value incorrect %v != testString", value)
		return
	}
}

func TestSetValueFromString_StringPtr(t *testing.T) {
	t.Parallel()
	var value *string
	if err := SetValueFromString(reflect.ValueOf(&value).Elem(), "testString"); err != nil {
		t.Error(err)
		return
	}
	if *value != "testString" {
		t.Errorf("Value incorrect %v != testString", value)
		return
	}
}

func TestSetValueFromString_Int(t *testing.T) {
	t.Parallel()
	var value int
	if err := SetValueFromString(reflect.ValueOf(&value).Elem(), "111"); err != nil {
		t.Error(err)
		return
	}
	if value != 111 {
		t.Errorf("Value incorrect %v != 111", value)
		return
	}
}

func TestSetValueFromString_IntPtr(t *testing.T) {
	t.Parallel()
	var value *int
	if err := SetValueFromString(reflect.ValueOf(&value).Elem(), "111"); err != nil {
		t.Error(err)
		return
	}
	if *value != 111 {
		t.Errorf("Value incorrect %v != 111", value)
		return
	}
}

func TestSetValueFromString_Int16(t *testing.T) {
	t.Parallel()
	var value int16
	if err := SetValueFromString(reflect.ValueOf(&value).Elem(), "111"); err != nil {
		t.Error(err)
		return
	}
	if value != 111 {
		t.Errorf("Value incorrect %v != 111", value)
		return
	}
}

func TestSetValueFromString_Int16Ptr(t *testing.T) {
	t.Parallel()
	var value *int16
	if err := SetValueFromString(reflect.ValueOf(&value).Elem(), "111"); err != nil {
		t.Error(err)
		return
	}
	if *value != 111 {
		t.Errorf("Value incorrect %v != 111", value)
		return
	}
}

func TestSetValueFromString_Int32(t *testing.T) {
	t.Parallel()
	var value int32
	if err := SetValueFromString(reflect.ValueOf(&value).Elem(), "111"); err != nil {
		t.Error(err)
		return
	}
	if value != 111 {
		t.Errorf("Value incorrect %v != 111", value)
		return
	}
}

func TestSetValueFromString_Int32Ptr(t *testing.T) {
	t.Parallel()
	var value *int32
	if err := SetValueFromString(reflect.ValueOf(&value).Elem(), "111"); err != nil {
		t.Error(err)
		return
	}
	if *value != 111 {
		t.Errorf("Value incorrect %v != 111", value)
		return
	}
}

func TestSetValueFromString_Int64(t *testing.T) {
	t.Parallel()
	var value int64
	if err := SetValueFromString(reflect.ValueOf(&value).Elem(), "111"); err != nil {
		t.Error(err)
		return
	}
	if value != 111 {
		t.Errorf("Value incorrect %v != 111", value)
		return
	}
}

func TestSetValueFromString_Int64Ptr(t *testing.T) {
	t.Parallel()
	var value *int64
	if err := SetValueFromString(reflect.ValueOf(&value).Elem(), "111"); err != nil {
		t.Error(err)
		return
	}
	if *value != 111 {
		t.Errorf("Value incorrect %v != 111", value)
		return
	}
}

func TestSetValueFromString_Bool(t *testing.T) {
	t.Parallel()
	var value bool
	if err := SetValueFromString(reflect.ValueOf(&value).Elem(), "true"); err != nil {
		t.Error(err)
		return
	}
	if value != true {
		t.Errorf("Value incorrect %v != 111", value)
		return
	}
}

func TestSetValueFromString_BoolPtr(t *testing.T) {
	t.Parallel()
	var value *bool
	if err := SetValueFromString(reflect.ValueOf(&value).Elem(), "true"); err != nil {
		t.Error(err)
		return
	}
	if *value != true {
		t.Errorf("Value incorrect %v != true", value)
		return
	}
}

func TestSetValueFromString_UnknowError(t *testing.T) {
	t.Parallel()
	var value struct{}
	if err := SetValueFromString(reflect.ValueOf(&value).Elem(), "true"); err == nil {
		t.Errorf("Should return error for unknow type")
		return
	}
}
