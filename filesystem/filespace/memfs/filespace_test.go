package memfs_test

import (
	"github.com/goatcms/goat-core/filesystem/filespace/memfs"
	"github.com/goatcms/goat-core/testbase"
	"testing"
)

func TestMkdir(t *testing.T) {
	// init
	fs, err := memfs.NewFilespace()
	if err != nil {
		t.Error(err)
	}
	// create directories
	path := "/mydir1/mydir2/mydir3"
	if err := fs.MkdirAll(path, 0777); err != nil {
		t.Error("Fail when create directories", err)
	}
	// test node type
	if !fs.IsDir("/mydir1/mydir2") {
		t.Error("node is not a directory or not exists")
	}
	if !fs.IsDir(path) {
		t.Error("node is not a directory or not exists")
	}
	if fs.IsDir("/noExistPath") {
		t.Error("node is not a directory or not exists")
	}
}

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
