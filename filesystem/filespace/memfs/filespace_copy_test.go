package memfs_test

import (
	"testing"

	"github.com/goatcms/goatcore/filesystem/filespace/memfs"
	"github.com/goatcms/goatcore/testbase"
)

func TestCopy(t *testing.T) {
	const (
		srcPath   = "src"
		destPath  = "dest"
		file1Path = "/d1/d2/f1.ex"
		file2Path = "/d1/z1/f2.exx"
	)

	// init
	fs, err := memfs.NewFilespace()
	if err != nil {
		t.Error(err)
	}
	//Create data
	testData1 := []byte("Content of file 1")
	testData2 := []byte("Content of file 2")

	// create test model
	fs.WriteFile(srcPath+file1Path, testData1, 0777)
	fs.WriteFile(srcPath+file2Path, testData2, 0777)

	// copy
	fs.Copy(srcPath, destPath)

	// test
	readData1, err := fs.ReadFile(destPath + file1Path)
	if err != nil {
		t.Error("can not read file1 after write data ", err)
	} else {
		if !testbase.ByteArrayEq(testData1, readData1) {
			t.Error("read1 and test1 data are diffrent ", testData1, readData1)
		}
	}
	readData2, err := fs.ReadFile(destPath + file2Path)
	if err != nil {
		t.Error("can not read file2 after write data ", err)
	} else {
		if !testbase.ByteArrayEq(testData2, readData2) {
			t.Error("read2 and test2 data are diffrent ", testData2, readData2)
		}
	}
}

func TestCopySingleFile(t *testing.T) {
	const (
		srcPath   = "src"
		destPath  = "dest"
		file1Path = "/d1/d2/f1.ex"
	)
	// init
	fs, err := memfs.NewFilespace()
	if err != nil {
		t.Error(err)
	}
	//Create data
	testData1 := []byte("Content of file 1")
	// create test model
	fs.WriteFile(srcPath+file1Path, testData1, 0777)
	// copy
	fs.Copy(srcPath+file1Path, destPath+file1Path)
	// test
	readData1, err := fs.ReadFile(destPath + file1Path)
	if err != nil {
		t.Error("can not read file1 after write data ", err)
	} else {
		if !testbase.ByteArrayEq(testData1, readData1) {
			t.Error("read1 and test1 data are diffrent ", testData1, readData1)
		}
	}
}
