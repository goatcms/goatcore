package memfs

import (
	"math/rand"
	"testing"
	"time"

	"github.com/goatcms/goatcore/workers"
)

var testPaths = []string{
	"masternode/dir1/dir1.2/file1.txt",
	"masternode/dir2/dir2.2/file2.txt",
	"masternode/dir3/dir3.2/file3.txt",
}

func TestSimpleConcurrentlyWrite(t *testing.T) {
	t.Parallel()
	// init
	fs, err := NewFilespace()
	if err != nil {
		t.Error(err)
	}
	// create directories
	for ij := workers.AsyncTestReapeat; ij > 0; ij-- {
		path := randomValue(testPaths)
		for i := workers.MaxJob; i > 0; i-- {
			go writeFileTestHelper(fs, path, t)
		}
	}
}

func randomValue(strs []string) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return strs[r.Intn(len(strs))]
}
