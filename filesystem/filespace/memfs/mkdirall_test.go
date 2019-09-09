package memfs

import (
	"os"
	"testing"
	"time"

	"github.com/goatcms/goatcore/filesystem"
)

func TestMKDirAll(t *testing.T) {
	var (
		err     error
		dir     *Dir
		treeDir *Dir
	)
	t.Parallel()
	root := NewDir("root", filesystem.DefaultUnixDirMode, time.Now(), []os.FileInfo{
		NewDir("dir1", filesystem.DefaultUnixDirMode, time.Now(), []os.FileInfo{}),
	})
	if dir, err = mkdirAll(root, "dir1/dir2/dir3/dir4", filesystem.DefaultUnixDirMode); err != nil {
		t.Error(err)
		return
	}
	if dir == nil || dir.Name() != "dir4" {
		t.Errorf("Should return last node")
	}
	if treeDir, err = getDirByPath(root, "dir1/dir2/dir3/dir4"); err != nil {
		t.Errorf("Should contains new node")
		return
	}
	if treeDir != dir {
		t.Errorf("return directory should be added to directory tree (dir1/dir2/dir3/dir4)")
	}
}
