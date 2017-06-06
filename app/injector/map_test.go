package injector

import (
	"testing"
)

type TestInjectableObject struct {
	SomeString string `tagname:"SomeStringKey"`
	SomeInt    int    `tagname:"SomeIntKey"`
}

func TestSimpleInject(t *testing.T) {
	t.Parallel()
	mapData := map[string]interface{}{}
	mapData["SomeStringKey"] = "SomeStringValue"
	mapData["SomeIntKey"] = int(11)

	object := &TestInjectableObject{}

	injector := NewMapInjector("tagname", mapData)
	injector.InjectTo(object)

	if object.SomeInt != 11 {
		t.Error("MapInjector didn't inject a int(11) to SomeInt field")
	}
	if object.SomeString != "SomeStringValue" {
		t.Error("MapInjector didn't inject 'SomeStringValue' to SomeString field")
	}
}
