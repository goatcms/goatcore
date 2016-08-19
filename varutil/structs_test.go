package varutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
	assert.Equal(t, mapData["Name"], structData.Name, "mapData and structData Name must be equal (%v %v)", mapData, structData)
	assert.Equal(t, mapData["Age"], structData.Age, "mapData and structData Age must be equal (%v %v)", mapData, structData)
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
	assert.Equal(t, mapData["Name"], structData.Name, "mapData and structData Name must be equal (%v %v)", mapData, structData)
	assert.Equal(t, mapData["myage"], structData.Age, "mapData and structData Age must be equal (%v %v)", mapData, structData)
}
