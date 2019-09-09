package memfs

import (
	"os"
	"testing"
	"time"

	"github.com/goatcms/goatcore/filesystem"
)

// testPathsRootDir contains test file tree
var testPathsRootDir = NewDir("root", filesystem.DefaultUnixDirMode, time.Now(), []os.FileInfo{
	NewDir("dir1", filesystem.DefaultUnixDirMode, time.Now(), []os.FileInfo{
		NewDir("dir2", filesystem.DefaultUnixDirMode, time.Now(), []os.FileInfo{
			NewFile("file", filesystem.DefaultUnixFileMode, time.Now(), []byte("abc")),
		}),
	}),
})

func writeFileTestHelper(fs filesystem.Filespace, path string, t *testing.T) {
	var err error
	testData := []byte("There is test data")
	if err = fs.WriteFile(path, testData, filesystem.DefaultUnixFileMode); err != nil {
		t.Error(err)
		return
	}
}
