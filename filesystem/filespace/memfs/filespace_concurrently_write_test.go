package memfs

import (
	"testing"

	"github.com/goatcms/goatcore/workers"
)

func TestConcurrentlyWrite(t *testing.T) {
	t.Parallel()
	// init
	fs, err := NewFilespace()
	if err != nil {
		t.Error(err)
	}
	// create directories
	for ij := workers.AsyncTestReapeat; ij > 0; ij-- {
		path := randomPath(5)
		for i := workers.MaxJob; i > 0; i-- {
			go writeFileTestHelper(fs, path, t)
		}
	}
}
