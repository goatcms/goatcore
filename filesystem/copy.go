package filesystem

import (
	"io"
	"os"
	"strings"
)

func Copy(src, dest string) error {
	if IsDir(src){
		return CopyDirectory(src, dest)
	}
	return cp(src, dest)
}

func CopyDirectory(src, dest string) error {
	if !strings.HasSuffix(src, "/") {
		src = src + "/"
	}
	if !strings.HasSuffix(dest, "/") {
		dest = dest + "/"
	}
	loop := DirLoop{
		OnFile: copyFileFactory(src, dest),
		OnDir:  copyDirFactory(dest),
		Filter: nil,
	}
	if err := loop.Run(src); err != nil {
		return err
	}
	return nil
}

func copyFileFactory(src, dest string) func(os.FileInfo, string) error {
	return func(file os.FileInfo, path string) error {
		return cp(src+path, dest+path)
	}
}

func copyDirFactory(dest string) func(os.FileInfo, string) error {
	return func(file os.FileInfo, path string) error {
		return os.MkdirAll(dest+path, FileMode)
	}
}

func cp(src, dst string) error {
	s, err := os.Open(src)
	if err != nil {
		return err
	}
	// no need to check errors on read only file, we already got everything
	// we need from the filesystem, so nothing can go wrong now.
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