package memfs

import (
	"os"
	"strings"

	"github.com/goatcms/goatcore/varutil/goaterr"
)

// removeByPath remove node by path
func removeNodeByPath(d *Dir, path string, emptyOnly bool) error {
	return removeNodeByNodePath(d, strings.Split(path, "/"), emptyOnly)
}

func removeNodeByNodePath(d *Dir, nodePath []string, emptyOnly bool) (err error) {
	var (
		dirNode *Dir
	)
	if len(nodePath) < 1 {
		return goaterr.Errorf("Expected node name")
	}
	dirNodePath := nodePath[:len(nodePath)-1]
	lastNodeName := nodePath[len(nodePath)-1]
	if dirNode, err = getDirByPathNodes(d, dirNodePath); err != nil {
		return err
	}
	if emptyOnly {
		var (
			ok       bool
			lastNode os.FileInfo
			lastDir  *Dir
		)
		if lastNode, err = dirNode.getNode(lastNodeName); err != nil {
			return err
		}
		if lastDir, ok = lastNode.(*Dir); !ok {
			return dirNode.removeNodeByName(lastNodeName)
		}
		if len(lastDir.nodes) != 0 {
			return goaterr.Errorf("Can not remove empty node")
		}
		return dirNode.removeNodeByName(lastNodeName)
	}
	return dirNode.removeNodeByName(lastNodeName)
}
