package encryptfs

import (
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/workers"
)

var randomPathNodes = []string{"dir1", "dir2", "dir3", "dir4", "dir5"}

func writeFileTestHelper(fs filesystem.Filespace, path string, t *testing.T) {
	var err error
	testData := []byte("There is test data")
	if err = fs.WriteFile(path, testData, filesystem.DefaultUnixFileMode); err != nil {
		t.Error(err)
		return
	}
}

func randomPath(n int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := []string{}
	for i := workers.MaxJob; i > 0; i-- {
		result = append(result, randomPathNodes[r.Intn(len(randomPathNodes))])
	}
	return strings.Join(result, "/")
}
