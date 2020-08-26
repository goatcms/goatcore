package diskfs

import (
	"path/filepath"
	"strings"

	"github.com/goatcms/goatcore/varutil/goaterr"
)

func safePath(base, subpath string) (path string, err error) {
	if path, err = filepath.Abs(base + subpath); err != nil {
		return "", err
	}
	if !strings.HasPrefix(path, base) {
		return "", goaterr.Errorf("break isolation space")
	}
	return path, nil
}
