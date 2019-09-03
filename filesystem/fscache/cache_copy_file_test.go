package fscache

import (
	"testing"

	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/filesystem/filespace/memfs"
	"github.com/goatcms/goatcore/testbase"
)

func TestCopyFile(t *testing.T) {
	var (
		remoteFS filesystem.Filespace
		cache    *Cache
		err      error
		result   []byte
		testData = []byte("Content of test file")
	)
	t.Parallel()
	if remoteFS, err = memfs.NewFilespace(); err != nil {
		t.Error(err)
		return
	}
	if cache, err = NewMemCache(remoteFS); err != nil {
		t.Error(err)
		return
	}
	// write
	if err = remoteFS.WriteFile("src/testfile.txt", testData, 0777); err != nil {
		t.Error(err)
		return
	}
	// test: make sure you have access to remote
	if !cache.IsExist("src/testfile.txt") {
		t.Errorf("remote filesystem should be readable ")
		return
	}
	// copy file
	if err = cache.Copy("src/testfile.txt", "dest/testfile.txt"); err != nil {
		t.Error(err)
		return
	}
	// test: before commit we shouldn't read new file from remote filespace
	if _, err = remoteFS.ReadFile("dest/testfile.txt"); err == nil {
		t.Errorf("before commit we shouldn't read new file from remote filespace")
	}
	// test: before commit we can read new file from cache
	if result, err = cache.ReadFile("dest/testfile.txt"); err != nil {
		t.Error(err)
	}
	if !testbase.ByteArrayEq(testData, result) {
		t.Error("test data and result are different ", testData, result)
	}
	// commit
	if err = cache.Commit(); err != nil {
		t.Error(err)
		return
	}
	// test: after commit we should read new file from remote filespace
	if result, err = remoteFS.ReadFile("dest/testfile.txt"); err != nil {
		t.Error(err)
	}
	if !testbase.ByteArrayEq(testData, result) {
		t.Error("test data and result are different ", testData, result)
	}
}
