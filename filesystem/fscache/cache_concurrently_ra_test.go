package fscache

import (
	"testing"

	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/filesystem/filespace/memfs"
	"github.com/goatcms/goatcore/workers"
)

// It work if don't throw a panic
func TestConcurrentlyRandomAccess(t *testing.T) {
	var (
		fs, baseFS  filesystem.Filespace
		err         error
		expetedText = "some data"
	)
	t.Parallel()
	// init
	if baseFS, err = memfs.NewFilespace(); err != nil {
		t.Error(err)
		return
	}
	if fs, err = NewMemCache(baseFS); err != nil {
		t.Error(err)
		return
	}
	// write and remove the same file
	for ij := workers.AsyncTestReapeat; ij > 0; ij-- {
		rPath := randomPath(5)
		for i := workers.MaxJob; i > 0; i-- {
			go (func() {
				fs.WriteFile(rPath, []byte(expetedText), filesystem.DefaultUnixDirMode)
			})()
		}
		for i := workers.MaxJob; i > 0; i-- {
			go (func() {
				fs.Remove(rPath)
			})()
		}
	}
	// create and remove direcotry
	for ij := workers.AsyncTestReapeat; ij > 0; ij-- {
		dirPath := randomPath(5)
		filePath := dirPath + "/file.txt"
		for i := workers.MaxJob; i > 0; i-- {
			go (func() {
				fs.WriteFile(filePath, []byte(expetedText), filesystem.DefaultUnixDirMode)
			})()
		}
		for i := workers.MaxJob; i > 0; i-- {
			go (func() {
				fs.RemoveAll(dirPath)
			})()
		}
	}
}
