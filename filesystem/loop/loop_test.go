package loop

import (
	"strings"
	"sync"
	"testing"

	"github.com/goatcms/goat-core/filesystem"
	"github.com/goatcms/goat-core/filesystem/filespace/memfs"
)

func TestCountFilesAndDirs(t *testing.T) {
	var (
		dirMU     sync.Mutex
		dirCount  int
		fileMU    sync.Mutex
		fileCount int
	)
	// init
	fs, err := memfs.NewFilespace()
	if err != nil {
		t.Error(err)
	}
	// create directories
	if err := fs.MkdirAll("mydir1/mydir1t1/mydir1t1t1", 0777); err != nil {
		t.Error(err)
		return
	}
	if err := fs.MkdirAll("mydir1/mydir1t2/mydirt1t2t1", 0777); err != nil {
		t.Error(err)
		return
	}
	// loop
	loop := Loop{
		FS: fs,
		OnDir: func(fs filesystem.Filespace, subPath string) error {
			dirMU.Lock()
			dirCount++
			dirMU.Unlock()
			return nil
		},
		OnFile: func(fs filesystem.Filespace, subPath string) error {
			fileMU.Lock()
			fileCount++
			fileMU.Unlock()
			return nil
		},
	}
	if err := loop.Run("./"); err != nil {
		t.Error(err)
		return
	}
	if err := loop.Wait(); err != nil {
		t.Error(err)
		return
	}
	// test result
	if dirCount != 5 {
		t.Errorf("dir counter should be equals to 5 (it is %v)", dirCount)
	}
	if fileCount != 0 {
		t.Errorf("file counter should be equals to 0 (it is %v)", fileCount)
	}
}

func TestFilter(t *testing.T) {
	var (
		dirMU    sync.Mutex
		dirCount int
	)
	// init
	fs, err := memfs.NewFilespace()
	if err != nil {
		t.Error(err)
	}
	// create directories
	if err := fs.MkdirAll("mydir1/mydir1t1.ex/mydir1t1t1.ex", 0777); err != nil {
		t.Error(err)
		return
	}
	if err := fs.MkdirAll("mydir1/mydir1t2.ex/mydirt1t2t1", 0777); err != nil {
		t.Error(err)
		return
	}
	// loop
	loop := Loop{
		FS: fs,
		Filter: func(fs filesystem.Filespace, subPath string) bool {
			return strings.HasSuffix(subPath, ".ex")
		},
		OnDir: func(fs filesystem.Filespace, subPath string) error {
			dirMU.Lock()
			dirCount++
			dirMU.Unlock()
			return nil
		},
	}
	if err := loop.Run("./"); err != nil {
		t.Error(err)
		return
	}
	if err := loop.Wait(); err != nil {
		t.Error(err)
		return
	}
	// test result
	if dirCount != 3 {
		t.Errorf("dir counter should be equals to 3 (it is %v)", dirCount)
	}
}
