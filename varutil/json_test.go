package varutil

import (
	"testing"
)

const (
	exampleJSON = `{
		"value1": "v1",
		"value2": "v2"
	}`
)

type TestObject struct {
	Value1 string `json:"value1"`
	Value2 string `json:"value2"`
}

func TestObjectToJSON(t *testing.T) {
	t.Parallel()
	var obj1 = TestObject{
		Value1: "v1",
		Value2: "v2",
	}
	json, err := ObjectToJSON(obj1)
	if err != nil {
		t.Error(err)
		return
	}
	var obj2 TestObject
	if err := ObjectFromJSON(&obj2, json); err != nil {
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

func TestObjectFromJSON(t *testing.T) {
	t.Parallel()
	var obj1 = TestObject{}
	err := ObjectFromJSON(&obj1, exampleJSON)
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
