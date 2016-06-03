package definition

import (
	"testing"
	"reflect"
)

func TestEncodeNewNodeDef(t *testing.T) {
	str := "nodename;typename;generatorname;param1;param2;name1:value1;name2:value2"
	nodes := reflect.ValueOf(map[string]string{})
	nd, err := EncodeNewNodeDef(str, nodes)
	if err != nil {
		t.Error(err)
		return;
	}
	if nd.Name() != "nodename" {
		t.Error("incorrect node name conversion")
	}
	//type
	td := nd.Type()
	if td.TypeName() != "typename" {
		t.Error("incorrect type conversion")
	}
	if td.GeneratorName() != "generatorname" {
		t.Error("incorrect generatorname conversion")
	}
	params := td.Params()
	//params - args
	args := params.Args()
	if  args[0] != "param1" {
		t.Error("incorrect argument[0] conversion")
	}
	if  args[1] != "param2" {
		t.Error("incorrect argument[1] conversion")
	}
	if  args[2] != "name1:value1" {
		t.Error("incorrect argument[2] conversion")
	}
	if  args[3] != "name2:value2" {
		t.Error("incorrect argument[3] conversion")
	}
	//params - keys
	if  params.Key("name1") != "value1" {
		t.Error("incorrect key value for key name1 (expected value1): ", params.Key("name1"))
	}
	if  params.Key("name2") != "value2" {
		t.Error("incorrect key value for key name2 (expected value2): ", params.Key("name2"))
	}
	if  params.Key("name3") != "" {
		t.Error("incorrect key value for key name3 (expected empty): ", params.Key("name1"))
	}
}
