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
