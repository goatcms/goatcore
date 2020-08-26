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

func reducePath(path string) (result string, err error) {
	var (
		baseNodes   = strings.Split(path, "/")
		resultNodes = make([]string, len(baseNodes))
		resultLen   = 0
	)
	for _, v := range baseNodes {
		if v == "" || v == "." {
			continue
		}
		if v == ".." {
			if resultLen == 0 {
				return "", goaterr.Errorf("%s: break isolation space", path)
			}
			resultLen--
			continue
		}
		resultNodes[resultLen] = v
		resultLen++
	}
	resultNodes = resultNodes[:resultLen]
	return strings.Join(resultNodes, "/"), nil
}
