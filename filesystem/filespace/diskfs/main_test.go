package diskfs

import (
	"os"
	"testing"

	"github.com/goatcms/goatcore/filesystem"
)

func NewBaseFS(t *testing.T, testName string) (fs filesystem.Filespace, err error) {
	var (
		path = "./teststmp/filesystem/filespace/diskfs/" + testName
	)
	if err = os.MkdirAll(path, filesystem.DefaultUnixDirMode); err != nil {
		return
	}
	t.Cleanup(func() {
		os.RemoveAll(path)
	})
	return NewFilespace(path)
}
