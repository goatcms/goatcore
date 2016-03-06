package varutil

import (
	"strings"
)

func FixDirPath(path *string) {
	if *path == "" {
		return
	}
	if !strings.HasSuffix(*path, "/") {
		*path = *path + "/"
	}
}
