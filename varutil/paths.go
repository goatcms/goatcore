package varutil

import (
	"path"
	"strings"

	"github.com/goatcms/goatcore/varutil/goaterr"
)

// GOPath return golang path like host.com/user/repo
// for example: github.com/goatcms/goatcore
// If it is not external path return empty string.
func GOPath(repourl string) (w string) {
	var e []string
	repourl = FullGOPath(repourl)
	if e = strings.Split(repourl, "/"); len(e) < 3 {
		return ""
	}
	return e[0] + "/" + e[1] + "/" + e[2]
}

// FullGOPath return full golang path like host.com/user/repo/dir1/dir2/endpackage
// for example: github.com/goatcms/goatcore/filesystem/disk
// If it is not external path return empty string.
func FullGOPath(repourl string) (w string) {
	var index int
	if index = strings.Index(repourl, "://"); index != -1 {
		repourl = repourl[index+3:]
	}
	if strings.HasSuffix(repourl, ".git") {
		repourl = repourl[0 : len(repourl)-4]
	}
	return repourl
}

// CleanPath clean path
func CleanPath(p string) (w string) {
	p = path.Clean(p)
	if strings.HasPrefix(p, "/") {
		p = p[1:]
	}
	return p
}

// ReduceAbsPath return shorter path version.
// Return error if the path refers to a parent directory.
func ReduceAbsPath(path string) (result string, err error) {
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
