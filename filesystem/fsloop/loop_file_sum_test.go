package fsloop

import (
	"strconv"
	"sync"
	"testing"

	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/filesystem/filespace/memfs"
	"github.com/goatcms/goatcore/workers"
)

const (
	f1Content = `1`
	f2Content = `2`
	f3Content = `3`
	f4Content = `4`

	expectedSum = 10
)

type TestSum struct {
	mu    sync.Mutex
	value int
}

func (t *TestSum) Sum(value int) {
	t.mu.Lock()
	t.value += value
	t.mu.Unlock()
}

func TestSimpleLoad(t *testing.T) {
	fs, err := memfs.NewFilespace()
	if err != nil {
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

	for i := 0; i < workers.AsyncTestReapeat; i++ {
		testSum := &TestSum{}

		loop := NewLoop(&LoopData{
			Filespace: fs,
			OnFile: func(fs filesystem.Filespace, subPath string) error {
				data, err := fs.ReadFile(subPath)
				if err != nil {
					return err
				}
				i, err := strconv.Atoi(string(data))
				if err != nil {
					return err
				}
				testSum.Sum(i)
				return nil
			},
		}, nil)
		loop.Run("./")
		loop.Wait()
		if errs := loop.Errors(); len(errs) != 0 {
			t.Errorf("errors: %v", loop.Errors())
			return
		}
		if testSum.value != expectedSum {
			t.Errorf("expect %v, result is %v", expectedSum, testSum.value)
			return
		}
	}
}
