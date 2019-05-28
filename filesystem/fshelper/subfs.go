package fshelper

import (
	"os"
	"path"

	"github.com/goatcms/goatcore/filesystem"
)

// SubFS is a filespace related to a base path in other filespace
type SubFS struct {
	basePath string
	fs       filesystem.Filespace
}

// NewSubFS create new sub filesystem. Related to parent filesystem and path.
func NewSubFS(fs filesystem.Filespace, basePath string) filesystem.Filespace {
	basePath = path.Clean(basePath) + "/"
	return SubFS{
		basePath: basePath,
		fs:       fs,
	}
}

// Copy method run Copy method of parent filesystem but in relative base path
func (sub SubFS) Copy(src, dest string) error {
	return sub.fs.Copy(sub.basePath+src, sub.basePath+dest)
}

// CopyDirectory method run CopyDirectory method of parent filesystem but in relative base path
func (sub SubFS) CopyDirectory(src, dest string) error {
	return sub.fs.CopyDirectory(sub.basePath+src, sub.basePath+dest)
}

// CopyFile method run CopyFile method of parent filesystem but in relative base path
func (sub SubFS) CopyFile(src, dest string) error {
	return sub.fs.CopyFile(sub.basePath+src, sub.basePath+dest)
}

// ReadDir method run ReadDir method of parent filesystem but in relative base path
func (sub SubFS) ReadDir(src string) ([]os.FileInfo, error) {
	return sub.fs.ReadDir(sub.basePath + src)
}

// IsExist method run IsExist method of parent filesystem but in relative base path
func (sub SubFS) IsExist(src string) bool {
	return sub.fs.IsExist(sub.basePath + src)
}

// IsFile method run IsFile method of parent filesystem but in relative base path
func (sub SubFS) IsFile(src string) bool {
	return sub.fs.IsFile(sub.basePath + src)
}

// IsDir method run IsDir method of parent filesystem but in relative base path
func (sub SubFS) IsDir(src string) bool {
	return sub.fs.IsDir(sub.basePath + src)
}

// MkdirAll method run MkdirAll method of parent filesystem but in relative base path
func (sub SubFS) MkdirAll(dest string, filemode os.FileMode) error {
	return sub.fs.MkdirAll(sub.basePath+dest, filemode)
}

// ReadFile method run ReadFile method of parent filesystem but in relative base path
func (sub SubFS) ReadFile(src string) ([]byte, error) {
	return sub.fs.ReadFile(sub.basePath + src)
}

// WriteFile method run WriteFile method of parent filesystem but in relative base path
func (sub SubFS) WriteFile(dest string, data []byte, perm os.FileMode) error {
	return sub.fs.WriteFile(sub.basePath+dest, data, perm)
}

// Filespace create new filespace
func (sub SubFS) Filespace(src string) (filesystem.Filespace, error) {
	return SubFS{
		basePath: sub.basePath + path.Clean(src) + "/",
		fs:       sub.fs,
	}, nil
}

// Reader method run Reader method of parent filesystem but in relative base path
func (sub SubFS) Reader(src string) (filesystem.Reader, error) {
	return sub.fs.Reader(sub.basePath + src)
}

// Writer method run Writer method of parent filesystem but in relative base path
func (sub SubFS) Writer(dest string) (filesystem.Writer, error) {
	return sub.fs.Writer(sub.basePath + dest)
}

// Remove method run Remove method of parent filesystem but in relative base path
func (sub SubFS) Remove(dest string) error {
	return sub.fs.Remove(sub.basePath + dest)
}

// RemoveAll method run RemoveAll method of parent filesystem but in relative base path
func (sub SubFS) RemoveAll(dest string) error {
	return sub.fs.RemoveAll(sub.basePath + dest)
}

// Lstat method run Lstat method of parent filesystem but in relative base path
func (sub SubFS) Lstat(src string) (os.FileInfo, error) {
	return sub.fs.Lstat(sub.basePath + src)
}
