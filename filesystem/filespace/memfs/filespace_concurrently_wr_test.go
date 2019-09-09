package memfs

import (
	"testing"

	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/workers"
)

func TestConcurrentlyRead(t *testing.T) {
	var (
		fs          filesystem.Filespace
		err         error
		expetedText = "some data"
	)
	t.Parallel()
	// init
	if fs, err = NewFilespace(); err != nil {
		t.Error(err)
		return
	}
	// create directories
	for ij := workers.AsyncTestReapeat; ij > 0; ij-- {
		rPath := randomPath(5)
		if err = fs.WriteFile(rPath, []byte(expetedText), filesystem.DefaultUnixDirMode); err != nil {
			t.Error(err)
			return
		}
		for i := workers.MaxJob; i > 0; i-- {
			go (func() {
				var data []byte
				if data, err = fs.ReadFile(rPath); err != nil {
					t.Error(err)
					return
				}
				if string(data) != expetedText {
					t.Errorf("Expected '%s' text and take %s", expetedText, data)
					return
				}
			})()
		}
	}
}
