package memfs

import (
	"testing"

	"github.com/goatcms/goatcore/filesystem"
)

func TestCopyIsolationWrapper(t *testing.T) {
	var (
		fs, childFS filesystem.Filespace
		err         error
	)
	if fs, err = NewFilespace(); err != nil {
		t.Error(err)
		return
	}
	if fs, err = NewFilespaceWrapper(fs, ""); err != nil {
		t.Error(err)
		return
	}
	if err = fs.MkdirAll("src/d1/d2/d3", filesystem.DefaultUnixDirMode); err != nil {
		t.Error(err)
		return
	}
	if err = fs.MkdirAll("child", filesystem.DefaultUnixDirMode); err != nil {
		t.Error(err)
		return
	}
	if childFS, err = fs.Filespace("child"); err != nil {
		t.Error(err)
		return
	}
	if err = childFS.Copy("../src", "dst"); err == nil {
		t.Errorf("Parent directory must be isolated")
		return
	}
}

func TestCopyDirectoryIsolationWrapper(t *testing.T) {
	var (
		fs, childFS filesystem.Filespace
		err         error
	)
	if fs, err = NewFilespace(); err != nil {
		t.Error(err)
		return
	}
	if fs, err = NewFilespaceWrapper(fs, ""); err != nil {
		t.Error(err)
		return
	}
	if err = fs.MkdirAll("src/d1/d2/d3", filesystem.DefaultUnixDirMode); err != nil {
		t.Error(err)
		return
	}
	if err = fs.MkdirAll("child", filesystem.DefaultUnixDirMode); err != nil {
		t.Error(err)
		return
	}
	if childFS, err = fs.Filespace("child"); err != nil {
		t.Error(err)
		return
	}
	if err = childFS.CopyDirectory("../src", "dst"); err == nil {
		t.Errorf("Parent directory must be isolated")
		return
	}
}

func TestCopyFileIsolationWrapper(t *testing.T) {
	var (
		fs, childFS filesystem.Filespace
		err         error
	)
	if fs, err = NewFilespace(); err != nil {
		t.Error(err)
		return
	}
	if fs, err = NewFilespaceWrapper(fs, ""); err != nil {
		t.Error(err)
		return
	}
	if err = fs.WriteFile("src.txt", []byte{1, 2, 3}, filesystem.DefaultUnixFileMode); err != nil {
		t.Error(err)
		return
	}
	if err = fs.MkdirAll("child", filesystem.DefaultUnixDirMode); err != nil {
		t.Error(err)
		return
	}
	if childFS, err = fs.Filespace("child"); err != nil {
		t.Error(err)
		return
	}
	if err = childFS.CopyFile("../src.txt", "dst.txt"); err == nil {
		t.Errorf("Parent directory must be isolated")
		return
	}
}

func TestReadDirIsolationWrapper(t *testing.T) {
	var (
		fs  filesystem.Filespace
		err error
	)
	if fs, err = NewFilespace(); err != nil {
		t.Error(err)
		return
	}
	if fs, err = NewFilespaceWrapper(fs, ""); err != nil {
		t.Error(err)
		return
	}
	if _, err = fs.ReadDir(".."); err == nil {
		t.Errorf("Parent directory must be isolated")
		return
	}
}

func TestIsExistIsolationWrapper(t *testing.T) {
	var (
		fs     filesystem.Filespace
		err    error
		result bool
	)
	if fs, err = NewFilespace(); err != nil {
		t.Error(err)
		return
	}
	if fs, err = NewFilespaceWrapper(fs, ""); err != nil {
		t.Error(err)
		return
	}
	if result = fs.IsExist(".."); result != false {
		t.Errorf("Only child directory or files should be checkable")
		return
	}
}

func TestIsFileIsolationWrapper(t *testing.T) {
	var (
		fs, childFS filesystem.Filespace
		err         error
		result      bool
	)
	if fs, err = NewFilespace(); err != nil {
		t.Error(err)
		return
	}
	if fs, err = NewFilespaceWrapper(fs, ""); err != nil {
		t.Error(err)
		return
	}
	if err = fs.WriteFile("src.txt", []byte{1, 2, 3}, filesystem.DefaultUnixFileMode); err != nil {
		t.Error(err)
		return
	}
	if err = fs.MkdirAll("child", filesystem.DefaultUnixDirMode); err != nil {
		t.Error(err)
		return
	}
	if childFS, err = fs.Filespace("child"); err != nil {
		t.Error(err)
		return
	}
	if result = childFS.IsDir("../src.txt"); result != false {
		t.Errorf("Only child files should be checkable")
		return
	}
}

func TestIsDirIsolationWrapper(t *testing.T) {
	var (
		fs     filesystem.Filespace
		err    error
		result bool
	)
	if fs, err = NewFilespace(); err != nil {
		t.Error(err)
		return
	}
	if fs, err = NewFilespaceWrapper(fs, ""); err != nil {
		t.Error(err)
		return
	}
	if result = fs.IsDir(".."); result != false {
		t.Errorf("Only child directory should be checkable")
		return
	}
}

func TestMkdirAllIsolationWrapper(t *testing.T) {
	var (
		fs, childFS filesystem.Filespace
		err         error
	)
	if fs, err = NewFilespace(); err != nil {
		t.Error(err)
		return
	}
	if fs, err = NewFilespaceWrapper(fs, ""); err != nil {
		t.Error(err)
		return
	}
	if err = fs.MkdirAll("child", filesystem.DefaultUnixDirMode); err != nil {
		t.Error(err)
		return
	}
	if childFS, err = fs.Filespace("child"); err != nil {
		t.Error(err)
		return
	}
	if err = childFS.MkdirAll("../a/b/c", filesystem.DefaultUnixDirMode); err == nil {
		t.Errorf("only child nodes access is allowed")
		return
	}
	if fs.IsExist("a") {
		t.Errorf("a directory shouldn't exist")
		return
	}
}

func TestWriterIsolationWrapper(t *testing.T) {
	var (
		fs  filesystem.Filespace
		err error
	)
	if fs, err = NewFilespace(); err != nil {
		t.Error(err)
		return
	}
	if fs, err = NewFilespaceWrapper(fs, ""); err != nil {
		t.Error(err)
		return
	}
	if _, err = fs.Writer("../result.txt"); err == nil {
		t.Errorf("only child nodes access is allowed")
		return
	}
}

func TestReaderIsolationWrapper(t *testing.T) {
	var (
		fs, childFS filesystem.Filespace
		err         error
	)
	if fs, err = NewFilespace(); err != nil {
		t.Error(err)
		return
	}
	if fs, err = NewFilespaceWrapper(fs, ""); err != nil {
		t.Error(err)
		return
	}
	if err = fs.WriteFile("src.txt", []byte{1, 2, 3}, filesystem.DefaultUnixFileMode); err != nil {
		t.Error(err)
		return
	}
	if err = fs.MkdirAll("child", filesystem.DefaultUnixDirMode); err != nil {
		t.Error(err)
		return
	}
	if childFS, err = fs.Filespace("child"); err != nil {
		t.Error(err)
		return
	}
	if _, err = childFS.Reader("../src.txt"); err == nil {
		t.Errorf("only child nodes access is allowed")
		return
	}
}

func TestReadFileIsolationWrapper(t *testing.T) {
	var (
		fs, childFS filesystem.Filespace
		err         error
	)
	if fs, err = NewFilespace(); err != nil {
		t.Error(err)
		return
	}
	if fs, err = NewFilespaceWrapper(fs, ""); err != nil {
		t.Error(err)
		return
	}
	if err = fs.WriteFile("src.txt", []byte{1, 2, 3}, filesystem.DefaultUnixFileMode); err != nil {
		t.Error(err)
		return
	}
	if err = fs.MkdirAll("child", filesystem.DefaultUnixDirMode); err != nil {
		t.Error(err)
		return
	}
	if childFS, err = fs.Filespace("child"); err != nil {
		t.Error(err)
		return
	}
	if _, err = childFS.ReadFile("../src.txt"); err == nil {
		t.Errorf("only child nodes access is allowed")
		return
	}
}

func TestWriteFileIsolationWrapper(t *testing.T) {
	var (
		fs, childFS filesystem.Filespace
		err         error
	)
	if fs, err = NewFilespace(); err != nil {
		t.Error(err)
		return
	}
	if fs, err = NewFilespaceWrapper(fs, ""); err != nil {
		t.Error(err)
		return
	}
	if err = fs.MkdirAll("child", filesystem.DefaultUnixDirMode); err != nil {
		t.Error(err)
		return
	}
	if childFS, err = fs.Filespace("child"); err != nil {
		t.Error(err)
		return
	}
	if err = childFS.WriteFile("../dest.txt", []byte{1, 2}, filesystem.DefaultUnixFileMode); err == nil {
		t.Errorf("only child nodes access is allowed")
		return
	}
	if fs.IsExist("dest.txt") {
		t.Errorf("dest.txt shouldn't exist")
		return
	}
}
func TestRemoveIsolationWrapper(t *testing.T) {
	var (
		fs, childFS filesystem.Filespace
		err         error
	)
	if fs, err = NewFilespace(); err != nil {
		t.Error(err)
		return
	}
	if fs, err = NewFilespaceWrapper(fs, ""); err != nil {
		t.Error(err)
		return
	}
	if err = fs.WriteFile("some.bin", []byte{1, 2, 3}, filesystem.DefaultUnixFileMode); err != nil {
		t.Error(err)
		return
	}
	if err = fs.MkdirAll("child", filesystem.DefaultUnixDirMode); err != nil {
		t.Error(err)
		return
	}
	if childFS, err = fs.Filespace("child"); err != nil {
		t.Error(err)
		return
	}
	if err = childFS.Remove("../some.bin"); err == nil {
		t.Errorf("only child nodes access is allowed")
		return
	}
	if !fs.IsExist("some.bin") {
		t.Errorf("some.bin should exist")
		return
	}
}

func TestRemoveAllIsolationWrapper(t *testing.T) {
	var (
		fs, childFS filesystem.Filespace
		err         error
	)
	if fs, err = NewFilespace(); err != nil {
		t.Error(err)
		return
	}
	if fs, err = NewFilespaceWrapper(fs, ""); err != nil {
		t.Error(err)
		return
	}
	if err = fs.MkdirAll("a/b/c", filesystem.DefaultUnixDirMode); err != nil {
		t.Error(err)
		return
	}
	if err = fs.MkdirAll("child", filesystem.DefaultUnixDirMode); err != nil {
		t.Error(err)
		return
	}
	if childFS, err = fs.Filespace("child"); err != nil {
		t.Error(err)
		return
	}
	if err = childFS.RemoveAll("../a"); err == nil {
		t.Errorf("only child nodes access is allowed")
		return
	}
	if !fs.IsExist("a") {
		t.Errorf("'./a' directory should exist")
		return
	}
}
func TestFilespaceIsolationWrapper(t *testing.T) {
	var (
		fs  filesystem.Filespace
		err error
	)
	if fs, err = NewFilespace(); err != nil {
		t.Error(err)
		return
	}
	if fs, err = NewFilespaceWrapper(fs, ""); err != nil {
		t.Error(err)
		return
	}
	if _, err = fs.Filespace(".."); err == nil {
		t.Errorf("only child nodes access is allowed")
		return
	}
}

func TestLstatIsolationWrapper(t *testing.T) {
	var (
		fs  filesystem.Filespace
		err error
	)
	if fs, err = NewFilespace(); err != nil {
		t.Error(err)
		return
	}
	if fs, err = NewFilespaceWrapper(fs, ""); err != nil {
		t.Error(err)
		return
	}
	if _, err = fs.Lstat(".."); err == nil {
		t.Errorf("only child nodes access is allowed")
		return
	}
}
