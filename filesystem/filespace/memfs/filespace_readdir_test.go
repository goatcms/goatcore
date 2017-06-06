package memfs_test

import (
	"testing"

	"github.com/goatcms/goatcore/filesystem/filespace/memfs"
)

func TestReadDir(t *testing.T) {
	t.Parallel()
	var dir1 bool
	var dir2 bool
	// init
	fs, err := memfs.NewFilespace()
	if err != nil {
		t.Error(err)
	}
	// prepare data
	if err := fs.MkdirAll("dir1", 0777); err != nil {
		t.Error(err)
		return
	}
	if err := fs.MkdirAll("dir2", 0777); err != nil {
		t.Error(err)
		return
	}
	//testing
	list, err := fs.ReadDir("./")
	if err != nil {
		t.Error(err)
		return
	}
	for _, file := range list {
		switch file.Name() {
		case "dir1":
			dir1 = true
		case "dir2":
			dir2 = true
		default:
			t.Errorf("unknown file %s", file.Name())
			return
		}
	}
	if !dir1 {
		t.Errorf("don't read dir1")
	}
	if !dir2 {
		t.Errorf("don't read dir2")
	}
}
