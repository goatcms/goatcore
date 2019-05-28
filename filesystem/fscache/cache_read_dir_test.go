package fscache

import (
	"os"
	"testing"

	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/filesystem/filespace/memfs"
)

func TestReadDirStory(t *testing.T) {
	var (
		remoteFS filesystem.Filespace
		cache    Cache
		list     []os.FileInfo
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
	if err = remoteFS.MkdirAll("some/remote/directory", filesystem.DefaultUnixDirMode); err != nil {
		t.Error(err)
		return
	}
	if err = cache.MkdirAll("some/cached/directory", filesystem.DefaultUnixDirMode); err != nil {
		t.Error(err)
		return
	}
	//read directories
	if list, err = cache.ReadDir("some"); err != nil {
		t.Error(err)
		return
	}
	// test
	if len(list) != 2 {
		t.Error("expected list contains two results")
		return
	}
	if list[0].Name() != "cached" && list[1].Name() != "cached" {
		t.Error("some/cached should exists")
		return
	}
	if list[0].Name() != "remote" && list[1].Name() != "remote" {
		t.Error("some/remote should exists")
		return
	}
}

func TestReadDirWithCachedDuplicateStory(t *testing.T) {
	var (
		remoteFS filesystem.Filespace
		cache    Cache
		list     []os.FileInfo
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
	if err = remoteFS.MkdirAll("some/remote/directory", filesystem.DefaultUnixDirMode); err != nil {
		t.Error(err)
		return
	}
	if err = cache.MkdirAll("some/remote/duplicated/directory", filesystem.DefaultUnixDirMode); err != nil {
		t.Error(err)
		return
	}
	if err = cache.MkdirAll("some/cached/directory", filesystem.DefaultUnixDirMode); err != nil {
		t.Error(err)
		return
	}
	//read directories
	if list, err = cache.ReadDir("some"); err != nil {
		t.Error(err)
		return
	}
	// test
	if len(list) != 2 {
		t.Error("expected list contains two results")
		return
	}
	if list[0].Name() != "cached" && list[1].Name() != "cached" {
		t.Error("some/cached should exists")
		return
	}
	if list[0].Name() != "remote" && list[1].Name() != "remote" {
		t.Error("some/remote should exists")
		return
	}
}
