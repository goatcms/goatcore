package fscache

import (
	"testing"

	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/filesystem/filespace/memfs"
	"github.com/goatcms/goatcore/testbase"
)

func TestWriteSingleFile(t *testing.T) {
	var (
		remoteFS filesystem.Filespace
		cache    Cache
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
	if err = cache.WriteFile("dir/testfile.txt", testData, 0777); err != nil {
		t.Error(err)
		return
	}
	// test: make sure it dont write to remote filespace before commit
	if remoteFS.IsExist("dir/testfile.txt") {
		t.Errorf("remote filesystem should be unmodified before commit")
		return
	}
	if err = cache.Commit(); err != nil {
		t.Error(err)
		return
	}
	// test
	if result, err = cache.ReadFile("dir/testfile.txt"); err != nil {
		t.Error(err)
	}
	if !testbase.ByteArrayEq(testData, result) {
		t.Error("test data and result are different ", testData, result)
	}
}
