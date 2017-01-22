package varutil

import "testing"

type MyStruct struct {
	Name string
	Age  int64 `source:"myage"`
}

func TestFillStruct(t *testing.T) {
	mapData := map[string]interface{}{}
	mapData["Name"] = "Sebastian"
	mapData["Age"] = int64(25)
	structData := &MyStruct{}
	err := FillStruct(structData, mapData)
	if err != nil {
		t.Error(err)
		return
	}
	if mapData["Name"] != structData.Name {
		t.Errorf("mapData and structData Name must be equal (%v %v)", mapData["Name"], structData.Name)
		return
	}
	if mapData["Age"] != structData.Age {
		t.Errorf("mapData and structData Age must be equal (%v %v)", mapData["Age"], structData.Age)
		return
	}
}

func TestLoadStruct(t *testing.T) {
	mapData := map[string]interface{}{}
	mapData["Name"] = "Sebastian"
	mapData["myage"] = int64(25)
	structData := &MyStruct{}
	err := LoadStruct(structData, mapData, "source", true)
	if err != nil {
		t.Error(err)
		return
	}
	if mapData["Name"] != structData.Name {
		t.Errorf("mapData and structData Name must be equal (%v %v)", mapData["Name"], structData.Name)
		return
	}
	if mapData["Age"] != nil {
		t.Errorf("mapData and structData Age must be equal (%v %v)", mapData["Age"], structData.Age)
		return
	}
}

func TestSetField(t *testing.T) {
	myStruct := &MyStruct{
		Name: "",
		Age:  0,
	}
	if err := SetField(myStruct, "Name", "Sebastian"); err != nil {
		t.Error(err)
		return
	}
	if err := SetField(myStruct, "Age", int64(25)); err != nil {
		t.Error(err)
		return
	}
	if myStruct.Name != "Sebastian" {
		t.Errorf("set name fail")
		return
	}
	if myStruct.Age != 25 {
		t.Errorf("set age fail")
		return
	}
}
