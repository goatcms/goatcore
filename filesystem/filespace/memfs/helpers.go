package memfs

import (
	"strings"

	"github.com/goatcms/goatcore/varutil/goaterr"
)

func splitContainsPath(p string) (dirNodePath []string, nodeName string, err error) {
	nodePath := strings.Split(p, "/")
	if len(nodePath) < 1 {
		return nil, "", goaterr.Errorf("Path must contains nodename")
	}
	dirNodePath = nodePath[:len(nodePath)-1]
	nodeName = nodePath[len(nodePath)-1]
	return dirNodePath, nodeName, nil
}
