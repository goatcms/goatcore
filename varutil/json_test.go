package varutil

import (
	"testing"
)

const (
	exampleJson = `{
		"value1": "v1",
		"value2": "v2"
	}`
)

type TestObject struct {
	Value1 string `json:"value1"`
	Value2 string `json:"value2"`
}

func TestObjectToJson(t *testing.T) {
	var obj1 TestObject = TestObject{
		Value1: "v1",
		Value2: "v2",
	}
	json, err := ObjectToJson(obj1)
	if err != nil {
		t.Error(err)
		return
	}
	var obj2 TestObject
	if err := ObjectFromJson(&obj2, json); err != nil {
		t.Error(err)
		return
	}
	if obj1.Value1 != obj2.Value1 {
		t.Error("Value1 is diffrent")
	}
	if obj1.Value2 != obj2.Value2 {
		t.Error("Value2 is diffrent")
	}
}

func TestObjectFromJson(t *testing.T) {
	var obj1 TestObject = TestObject{}
	err := ObjectFromJson(&obj1, exampleJson)
	if err != nil {
		t.Error(err)
		return
	}
	if obj1.Value1 != "v1" {
		t.Error("Value1 import is incorrect")
	}
	if obj1.Value2 != "v2" {
		t.Error("Value2 import is incorrect")
	}
}
