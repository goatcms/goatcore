package memfs

import (
	"os"
	"strings"

	"github.com/goatcms/goatcore/varutil"
)

// mkdirAll crete directories recursive
func mkdirAll(d *Dir, subPath string, filemode os.FileMode) (dir *Dir, err error) {
	nodesPath := strings.Split(varutil.CleanPath(subPath), "/")
	return mkdirAllNodes(d, nodesPath, filemode)
}

func mkdirAllNodes(d *Dir, nodesPath []string, filemode os.FileMode) (dir *Dir, err error) {
	for _, nodeName := range nodesPath {
		if d, err = d.mkdir(nodeName, filemode); err != nil {
			return nil, err
		}
	}
	return d, nil
}
