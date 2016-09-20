package json_test

import (
	"testing"

	"github.com/goatcms/goat-core/filesystem/filespace/memfs"
	"github.com/goatcms/goat-core/filesystem/json"
)

type TestObject struct {
	Value1 string `json:"value1"`
	Value2 string `json:"value2"`
}

func TestWriteAndRead(t *testing.T) {
	var writeObject, readObject TestObject
	// init
	fs, err := memfs.NewFilespace()
	if err != nil {
		t.Error(err)
	}
	// create test data
	path := "fyfile.json"
	writeObject = TestObject{
		Value1: "str1",
		Value2: "str2",
	}
	// write & read
	json.WriteJSON(fs, path, writeObject)
	json.ReadJSON(fs, path, &readObject)
	// test node type
	if !fs.IsFile(path) {
		t.Error("filesystem not conatin file adter write")
	}
	if writeObject.Value1 != readObject.Value1 || writeObject.Value2 != readObject.Value2 {
		t.Error("read date is wrong", writeObject, readObject)
	}
}
