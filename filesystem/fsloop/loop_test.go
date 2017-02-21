package fsloop

import (
	"fmt"
	"strings"
	"sync"
	"testing"

	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/filesystem/filespace/memfs"
	"github.com/goatcms/goatcore/workers"
)

type TestCounter struct {
	mu         sync.Mutex
	DirCounter int
}

func (t *TestCounter) CountDir(fs filesystem.Filespace, subPath string) error {
	t.mu.Lock()
	t.DirCounter++
	t.mu.Unlock()
	return nil
}

func TestFilter(t *testing.T) {
	fileIteratorCounter := 1000
	expectedCount := 4 * fileIteratorCounter
	for ti := 0; ti < workers.AsyncTestReapeat; ti++ {
		counter := &TestCounter{
			DirCounter: 0,
		}
		// init
		fs, err := memfs.NewFilespace()
		if err != nil {
			t.Error(err)
		}
		// create directories
		for i := 0; i < fileIteratorCounter; i++ {
			path := fmt.Sprintf("mydir%v.ex/mydir%vt1.ex/mydir%vt1t1.ex", i, i, i)
			if err := fs.MkdirAll(path, 0777); err != nil {
				t.Error(err)
				return
			}
			path = fmt.Sprintf("mydir%v.ex/mydir%vt2.ex/mydir%vt1t1", i, i, i)
			if err := fs.MkdirAll(path, 0777); err != nil {
				t.Error(err)
				return
			}
		}
		// loop
		loop := NewLoop(&LoopData{
			Filespace: fs,
			DirFilter: func(fs filesystem.Filespace, subPath string) bool {
				return strings.HasSuffix(subPath, ".ex")
			},
			OnDir: counter.CountDir,
			//Consumers:  1,
			//Producents: 1,
		}, nil)
		loop.Run("./")
		loop.Wait()
		// test result
		if len(loop.Errors()) != 0 {
			t.Errorf("Errors: %v", loop.Errors())
			return
		}
		if loop.lifecycle.IsKilled() {
			t.Errorf("loop was killed")
			return
		}
		if counter.DirCounter != expectedCount {
			t.Errorf("dir counter is equals to %v (expected value is: %v)", counter.DirCounter, expectedCount)
			return
		}
	}
}
