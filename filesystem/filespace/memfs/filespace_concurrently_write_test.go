package memfs

import (
	"math/rand"
	"strings"
	"testing"
	"time"

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

var randomPathNodes = []string{"dir1", "dir2", "dir3", "dir4", "dir5"}

func randomPath(n int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := []string{}
	for i := workers.MaxJob; i > 0; i-- {
		result = append(result, randomPathNodes[r.Intn(len(randomPathNodes))])
	}
	return strings.Join(result, "/")
}
