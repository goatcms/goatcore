package memfs

import (
	"os"
	"strings"

	"github.com/goatcms/goatcore/varutil/goaterr"
)

// getNodeByPath return node by path
func getNodeByPath(d *Dir, path string) (node os.FileInfo, err error) {
	if path == "." {
		return d, nil
	}
	return getNodeByPathNodes(d, strings.Split(path, "/"))
}

// getDirByPath return directory node by path
func getDirByPath(d *Dir, path string) (node *Dir, err error) {
	if path == "." {
		return d, nil
	}
	return getDirByPathNodes(d, strings.Split(path, "/"))
}

// getFileByPath return file node by path
func getFileByPath(d *Dir, path string) (node *File, err error) {
	if path == "." {
		return nil, goaterr.Errorf("Current node is directory (path: %s)", path)
	}
	return getFileByPathNodes(d, strings.Split(path, "/"))
}

// getDirByPathNodes return directory node by path nodes
func getDirByPathNodes(dir *Dir, pathNodes []string) (out *Dir, err error) {
	var (
		node os.FileInfo
		ok   bool
	)
	if node, err = getNodeByPathNodes(dir, pathNodes); err != nil {
		return nil, err
	}
	if out, ok = node.(*Dir); !ok {
		return nil, goaterr.Errorf("Node %s is not a directory in %v", node.Name(), pathNodes)
	}
	return out, nil
}

// getFileByPathNodes return file node by path nodes
func getFileByPathNodes(dir *Dir, pathNodes []string) (out *File, err error) {
	var (
		node os.FileInfo
		ok   bool
	)
	if node, err = getNodeByPathNodes(dir, pathNodes); err != nil {
		return nil, err
	}
	if out, ok = node.(*File); !ok {
		return nil, goaterr.Errorf("Node %s is not a file in %v", node.Name(), pathNodes)
	}
	return out, nil
}

// getNodeByPathNodes return node by path nodes
func getNodeByPathNodes(dir *Dir, pathNodes []string) (node os.FileInfo, err error) {
	if len(pathNodes) == 0 {
		return dir, nil
	}
	if node, err = dir.getNode(pathNodes[0]); err != nil {
		return nil, err
	}
	for _, nodeName := range pathNodes[1:] {
		if node.IsDir() != true {
			return nil, goaterr.Errorf("Node by name %s must be dir to get sub node (path %v )", node.Name(), pathNodes)
		}
		dir = node.(*Dir)
		if node, err = dir.getNode(nodeName); err != nil {
			return nil, err
		}
	}
	return node, nil
}
