package diskfs

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/filesystem/disk"
	"github.com/goatcms/goatcore/varutil"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

// Filespace is a files set on local filesystem
type Filespace struct {
	path string
}

// NewFilespace create new Filespace instance
func NewFilespace(path string) (filesystem.Filespace, error) {
	varutil.FixDirPath(&path)
	return filesystem.Filespace(&Filespace{
		path: path,
	}), nil
}

// Copy a file or a directory inside the filespace
func (fs *Filespace) Copy(src, dest string) error {
	return disk.Copy(fs.path+src, fs.path+dest)
}

// CopyDirectory duplicate a directory
func (fs *Filespace) CopyDirectory(src, dest string) error {
	return disk.CopyDirectory(fs.path+src, fs.path+dest)
}

// CopyFile duplicate a file
func (fs *Filespace) CopyFile(src, dest string) error {
	return disk.CopyFile(fs.path+src, fs.path+dest)
}

// ReadDir return directory nodes
func (fs *Filespace) ReadDir(subPath string) ([]os.FileInfo, error) {
	return ioutil.ReadDir(fs.path + subPath)
}

// IsExist return true if node exist
func (fs *Filespace) IsExist(subPath string) bool {
	return disk.IsExist(fs.path + subPath)
}

// IsFile return true if node exist and is a file
func (fs *Filespace) IsFile(subPath string) bool {
	return disk.IsFile(fs.path + subPath)
}

// IsDir return true if node exist and is a directory
func (fs *Filespace) IsDir(subPath string) bool {
	return disk.IsDir(fs.path + subPath)
}

// MkdirAll create directory recursively
func (fs *Filespace) MkdirAll(subPath string, filemode os.FileMode) error {
	return disk.MkdirAll(fs.path+subPath, filemode)
}

// Writer return a file node writer
func (fs *Filespace) Writer(subPath string) (filesystem.Writer, error) {
	return os.OpenFile(fs.path+subPath, os.O_WRONLY|os.O_CREATE, filesystem.DefaultUnixFileMode)
}

// Reader return a file node reader
func (fs *Filespace) Reader(subPath string) (filesystem.Reader, error) {
	return os.OpenFile(fs.path+subPath, os.O_RDONLY, filesystem.DefaultUnixFileMode)
}

// ReadFile return file data
func (fs *Filespace) ReadFile(subPath string) ([]byte, error) {
	return ioutil.ReadFile(fs.path + subPath)
}

// WriteFile write file data
func (fs *Filespace) WriteFile(subPath string, data []byte, perm os.FileMode) (err error) {
	path := fs.path + subPath
	if err = disk.MkdirAll(filepath.Dir(path), filesystem.DefaultUnixDirMode); err != nil {
		return err
	}
	return ioutil.WriteFile(path, data, perm)
}

// Remove delete node by path
func (fs *Filespace) Remove(subPath string) error {
	return os.Remove(fs.path + subPath)
}

// RemoveAll delete node by path recursively
func (fs *Filespace) RemoveAll(subPath string) error {
	return os.RemoveAll(fs.path + subPath)
}

// Filespace get directory node and return it as filespace
func (fs *Filespace) Filespace(subPath string) (filesystem.Filespace, error) {
	if !fs.IsDir(subPath) {
		return nil, goaterr.Errorf("Path is not a directory " + fs.path + subPath)
	}
	return NewFilespace(fs.path + subPath)
}

// Lstat returns a FileInfo describing the named file.
func (fs *Filespace) Lstat(subPath string) (os.FileInfo, error) {
	return os.Lstat(fs.path + subPath)
}

// LocalPath return path in local filesystem
func (fs *Filespace) LocalPath() string {
	return fs.path
}
