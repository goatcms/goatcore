package datascope

import (
	"testing"
)

func TestSimpleInject(t *testing.T) {
	t.Parallel()
	var object struct {
		SomeString string `tagname:"SomeStringKey"`
		SomeInt    int    `tagname:"SomeIntKey"`
	}
	dataScope := New(map[interface{}]interface{}{
		"SomeStringKey": "SomeStringValue",
		"SomeIntKey":    int(11),
	})
	injector := NewInjector("tagname", dataScope)
	injector.InjectTo(&object)
	if object.SomeInt != 11 {
		t.Error("MapInjector didn't inject a int(11) to SomeInt field")
	}
	if object.SomeString != "SomeStringValue" {
		t.Error("MapInjector didn't inject 'SomeStringValue' to SomeString field")
	}
}
