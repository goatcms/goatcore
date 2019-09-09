package memfs

import (
	"bytes"
	"testing"
	"time"
)

func TestCopyFile(t *testing.T) {
	var (
		f1, f2 *File
		err    error
	)
	t.Parallel()
	f1 = NewFile("fileName", 0676, time.Now(), []byte("some text data"))
	if f2, err = copyFile(f1, f1.Name()); err != nil {
		t.Error(err)
		return
	}
	if f1.Name() != f2.Name() {
		t.Errorf("Name are different")
	}
	if f1.Mode() != f2.Mode() {
		t.Errorf("Mode are different")
	}
	if f1.ModTime() != f2.ModTime() {
		t.Errorf("ModTime are different")
	}
	if bytes.Compare(f1.getData(), f2.getData()) != 0 {
		t.Errorf("Data are different. Take %s and expect %s", f2.getData(), f1.getData())
	}
}

func TestCopyDirectory(t *testing.T) {
	var (
		copiedRoot *Dir
		file       *File
		err        error
	)
	t.Parallel()
	if copiedRoot, err = copyDir(testPathsRootDir, testPathsRootDir.Name()); err != nil {
		t.Error(err)
		return
	}
	if file, err = getFileByPath(copiedRoot, "dir1/dir2/file"); err != nil {
		t.Error(err)
		return
	}
	copiedFileData := file.getData()
	if bytes.Compare(copiedFileData, []byte("abc")) != 0 {
		t.Errorf("Expected copied file data equals to 'abc' and take %s", copiedFileData)
	}
}
