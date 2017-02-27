package memfs_test

import (
	"testing"

	"github.com/goatcms/goatcore/filesystem/filespace/memfs"
)

func TestRemove(t *testing.T) {
	// init
	fs, err := memfs.NewFilespace()
	if err != nil {
		t.Error(err)
	}
	// create directories
	path := "/mydir1/mydir2/mydir3/mydir4"
	if err := fs.MkdirAll(path, 0777); err != nil {
		t.Error("Fail when create directories", err)
		return
	}
	// test node type
	if fs.Remove("/mydir1/mydir2") == nil {
		t.Error("Remove remove no empty directory is not allowed")
	}
	if err := fs.Remove("/mydir1/mydir2/mydir3/mydir4"); err != nil {
		t.Errorf("Remove should remove empty directory (Error: %v)", err)
	}
	if err := fs.RemoveAll("/mydir1"); err != nil {
		t.Errorf("RemoveAll should remove empty directory (Error: %v )", err)
	}
}
