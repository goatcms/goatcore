package fscache

import (
	"testing"

	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/filesystem/filespace/memfs"
	"github.com/goatcms/goatcore/testbase"
)

func TestWriteFileInRemoteExistDirectory(t *testing.T) {
	var (
		remoteFS filesystem.Filespace
		cache    *Cache
		err      error
		result   []byte
		w        filesystem.Writer
		testData = []byte("Content of test file")
	)
	t.Parallel()
	if remoteFS, err = memfs.NewFilespace(); err != nil {
		t.Error(err)
		return
	}
	if w, err = remoteFS.Writer("dir1/dir2/testfile.txt"); err != nil {
		t.Error(err)
		return
	}
	w.Write([]byte(testData))
	w.Close()
	if cache, err = NewMemCache(remoteFS); err != nil {
		t.Error(err)
		return
	}
	// write
	if w, err = cache.Writer("dir1/testfile.txt"); err != nil {
		t.Error(err)
		return
	}
	if _, err = w.Write([]byte(testData)); err != nil {
		t.Error(err)
		return
	}
	w.Close()
	// test: make sure it dont write to remote filespace before commit
	if remoteFS.IsExist("dir1/testfile.txt") {
		t.Errorf("remote filesystem should be unmodified before commit")
		return
	}
	if err = cache.Commit(); err != nil {
		t.Error(err)
		return
	}
	// test
	if result, err = cache.ReadFile("dir1/testfile.txt"); err != nil {
		t.Error(err)
	}
	if !testbase.ByteArrayEq(testData, result) {
		t.Error("test data and result are different ", testData, result)
	}
}
