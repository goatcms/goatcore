package memfs_test

import (
	"testing"

	"github.com/goatcms/goatcore/filesystem/filespace/memfs"
	"github.com/goatcms/goatcore/testbase"
)

func TestWriteAndRead(t *testing.T) {
	// init
	fs, err := memfs.NewFilespace()
	if err != nil {
		t.Error(err)
	}
	//Create data
	testData := []byte("There is test data")

	// create directories
	path := "/mydir1/mydir2/mydir3/myfile.ex"
	fs.WriteFile(path, testData, 0777)
	readData, err := fs.ReadFile(path)
	if err != nil {
		t.Error("can not read file after write data ", err)
	}
	if !testbase.ByteArrayEq(readData, testData) {
		t.Error("read data are diffrent ", readData, testData)
	}
}
