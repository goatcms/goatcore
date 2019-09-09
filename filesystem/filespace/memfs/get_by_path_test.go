package memfs

import (
	"os"
	"testing"

	"github.com/goatcms/goatcore/workers"
)

func TestGetNodeByPath(t *testing.T) {
	t.Parallel()
	for ij := workers.AsyncTestReapeat; ij > 0; ij-- {
		go (func() {
			var node os.FileInfo
			var err error
			if node, err = getNodeByPath(testPathsRootDir, "dir1/dir2/file"); err != nil {
				t.Error(err)
				return
			}
			if node == nil {
				t.Errorf("expected node and take nil")
				return
			}
			if node.Name() != "file" {
				t.Errorf("expected node named 'file' and take %s", node.Name())
			}
		})()
	}
}

func TestGetDirByPath(t *testing.T) {
	t.Parallel()
	for ij := workers.AsyncTestReapeat; ij > 0; ij-- {
		go (func() {
			var node *Dir
			var err error
			if node, err = getDirByPath(testPathsRootDir, "dir1/dir2"); err != nil {
				t.Error(err)
				return
			}
			if node == nil {
				t.Errorf("expected node and take nil")
				return
			}
			if node.Name() != "dir2" {
				t.Errorf("expected node named 'dir2' and take %s", node.Name())
			}
		})()
	}
}

func TestGetDirByPathFail(t *testing.T) {
	t.Parallel()
	for ij := workers.AsyncTestReapeat; ij > 0; ij-- {
		go (func() {
			var err error
			if _, err = getDirByPath(testPathsRootDir, "dir1/dir2/file"); err == nil {
				t.Errorf("Should fail when node is a file")
				return
			}
		})()
	}
}

func TestGetFileByPath(t *testing.T) {
	t.Parallel()
	for ij := workers.AsyncTestReapeat; ij > 0; ij-- {
		go (func() {
			var node *File
			var err error
			if node, err = getFileByPath(testPathsRootDir, "dir1/dir2/file"); err != nil {
				t.Error(err)
				return
			}
			if node == nil {
				t.Errorf("expected node and take nil")
				return
			}
			if node.Name() != "file" {
				t.Errorf("expected node named 'file' and take %s", node.Name())
			}
		})()
	}
}

func TestGetFileByPathFail(t *testing.T) {
	t.Parallel()
	for ij := workers.AsyncTestReapeat; ij > 0; ij-- {
		go (func() {
			var err error
			if _, err = getFileByPath(testPathsRootDir, "dir1/dir2"); err == nil {
				t.Errorf("Should fail when node is a directory")
				return
			}
		})()
	}
}
