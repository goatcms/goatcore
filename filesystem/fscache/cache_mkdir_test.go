package fscache

import (
	"testing"

	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/filesystem/filespace/memfs"
)

func TestMkdirStory(t *testing.T) {
	var (
		remoteFS filesystem.Filespace
		cache    *Cache
		err      error
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
	if err = cache.MkdirAll("some/new/directory", filesystem.DefaultUnixDirMode); err != nil {
		t.Error(err)
		return
	}
	// test: make sure it don't write to remote filespace before commit
	if remoteFS.IsExist("some/new/directory") {
		t.Errorf("remote filesystem should be unmodified before commit")
		return
	}
	if err = cache.Commit(); err != nil {
		t.Error(err)
		return
	}
	// test
	if !remoteFS.IsDir("some/new/directory") {
		t.Error("some/new/directory should be created")
		return
	}
}
