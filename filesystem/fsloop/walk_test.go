package fsloop

import (
	"os"
	"testing"

	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/filesystem/filespace/memfs"
)

func TestWalk(t *testing.T) {
	var (
		fs    filesystem.Filespace
		err   error
		files []string
		dirs  []string
	)
	if fs, err = memfs.NewFilespace(); err != nil {
		t.Error(err)
		return
	}
	if err = fs.WriteFile("f1.int", []byte(f1Content), 0777); err != nil {
		t.Error(err)
		return
	}
	if err = fs.WriteFile("f2.int", []byte(f2Content), 0777); err != nil {
		t.Error(err)
		return
	}
	if err = fs.WriteFile("dir/f3.int", []byte(f3Content), 0777); err != nil {
		t.Error(err)
		return
	}
	if err = fs.WriteFile("dir/f4.json", []byte(f4Content), 0777); err != nil {
		t.Error(err)
		return
	}
	if err = WalkFS(fs, ".", func(path string, info os.FileInfo) (err error) {
		files = append(files, path)
		return nil
	}, func(path string, info os.FileInfo) (err error) {
		dirs = append(dirs, path)
		return nil
	}); err != nil {
		t.Error(err)
		return
	}
	if len(files) != 4 {
		t.Errorf("Expected 4 file and take %v", files)
	}
	if len(dirs) != 1 {
		t.Errorf("Expected 1 directory and take %v", dirs)
	}
}
