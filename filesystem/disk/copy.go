package disk

import (
	"io"
	"os"
	"path/filepath"

	"github.com/goatcms/goatcore/filesystem"
)

// Copy duplicate a file or a directory
func Copy(src, dest string) error {
	if IsDir(src) {
		return CopyDirectory(src, dest)
	}
	return CopyFile(src, dest)
}

// CopyDirectory copy a directory and sub-direcotories and files on local files system.
func CopyDirectory(src, dest string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		subPath := path + "/" + info.Name()
		if info.IsDir() {
			return MkdirAll(subPath, filesystem.DefaultUnixDirMode)
		}
		return CopyFile(src+subPath, dest+subPath)
	})
}

// CopyFile copy a single file on local files system.
func CopyFile(src, dst string) error {
	s, err := os.Open(src)
	if err != nil {
		return err
	}
	defer s.Close()
	d, err := os.Create(dst)
	if err != nil {
		return err
	}
	if _, err := io.Copy(d, s); err != nil {
		d.Close()
		return err
	}
	return d.Close()
}
