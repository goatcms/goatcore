package fshelper

import (
	"fmt"
	"os"

	"github.com/goatcms/goatcore/filesystem"
)

// ROFilespace is a readonly filespace
type ROFilespace struct {
	fs filesystem.Filespace
}

// NewReadonlyFS create readonly mask for filespace
func NewReadonlyFS(fs filesystem.Filespace) filesystem.Filespace {
	return ROFilespace{
		fs: fs,
	}
}

// Copy method is unavailable
func (ro ROFilespace) Copy(src, dest string) error {
	return fmt.Errorf("Copy is unavailable on readonly filespace")
}

// CopyDirectory is unavailable
func (ro ROFilespace) CopyDirectory(src, dest string) error {
	return fmt.Errorf("CopyDirectory is unavailable on readonly filespace")
}

// CopyFile method is unavailable
func (ro ROFilespace) CopyFile(src, dest string) error {
	return fmt.Errorf("CopyFile is unavailable on readonly filespace")
}

// ReadDir method run ReadDir method of parent filespace
func (ro ROFilespace) ReadDir(src string) ([]os.FileInfo, error) {
	return ro.fs.ReadDir(src)
}

// IsExist method run IsExist method of parent filespace
func (ro ROFilespace) IsExist(src string) bool {
	return ro.fs.IsExist(src)
}

// IsFile method run IsFile method of parent filespace
func (ro ROFilespace) IsFile(src string) bool {
	return ro.fs.IsFile(src)
}

// IsDir method run IsDir method of parent filespace
func (ro ROFilespace) IsDir(src string) bool {
	return ro.fs.IsDir(src)
}

// MkdirAll method is unavailable
func (ro ROFilespace) MkdirAll(dest string, filemode os.FileMode) error {
	return fmt.Errorf("MkdirAll is unavailable on readonly filespace")
}

// ReadFile method run ReadFile method of parent filespace
func (ro ROFilespace) ReadFile(src string) ([]byte, error) {
	return ro.fs.ReadFile(src)
}

// WriteFile method is unavailable
func (ro ROFilespace) WriteFile(dest string, data []byte, perm os.FileMode) error {
	return fmt.Errorf("WriteFile is unavailable on readonly filespace")
}

// Filespace create new filespace
func (ro ROFilespace) Filespace(src string) (filesystem.Filespace, error) {
	return NewReadonlyFS(NewSubFS(ro.fs, src)), nil
}

// Reader method run Reader method of parent filespace
func (ro ROFilespace) Reader(src string) (filesystem.Reader, error) {
	return ro.fs.Reader(src)
}

// Writer method is unavailable
func (ro ROFilespace) Writer(dest string) (filesystem.Writer, error) {
	return nil, fmt.Errorf("Writer is unavailable on readonly filespace")
}

// Remove method is unavailable
func (ro ROFilespace) Remove(dest string) error {
	return fmt.Errorf("Remove is unavailable on readonly filespace")
}

// RemoveAll method is unavailable
func (ro ROFilespace) RemoveAll(dest string) error {
	return fmt.Errorf("RemoveAll is unavailable on readonly filespace")
}

// Lstat method run Lstat method of parent filespace
func (ro ROFilespace) Lstat(src string) (os.FileInfo, error) {
	return ro.fs.Lstat(src)
}
