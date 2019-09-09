package memfs

import (
	"os"

	"github.com/goatcms/goatcore/filesystem"
)

// FilespaceWrapper is memory filespace wraper
type FilespaceWrapper struct {
	basePath string
	fs       filesystem.Filespace
}

// NewFilespaceWrapper create new memory filespace  wrapper instance
func NewFilespaceWrapper(fs filesystem.Filespace, basePath string) filesystem.Filespace {
	return &FilespaceWrapper{
		fs:       fs,
		basePath: basePath + "/",
	}
}

// Copy duplicate a file or directory
func (w *FilespaceWrapper) Copy(src, dest string) (err error) {
	return w.fs.Copy(w.basePath+src, w.basePath+dest)
}

// CopyDirectory duplicate a directory
func (w *FilespaceWrapper) CopyDirectory(src, dest string) error {
	return w.fs.CopyDirectory(w.basePath+src, w.basePath+dest)
}

// CopyFile duplicate a file
func (w *FilespaceWrapper) CopyFile(src, dest string) error {
	return w.fs.CopyFile(w.basePath+src, w.basePath+dest)
}

// ReadDir return directory nodes
func (w *FilespaceWrapper) ReadDir(srcPath string) (nodes []os.FileInfo, err error) {
	return w.fs.ReadDir(w.basePath + srcPath)
}

// IsExist return true if node exist
func (w *FilespaceWrapper) IsExist(srcPath string) bool {
	return w.fs.IsExist(w.basePath + srcPath)
}

// IsFile return true if node exist and is a file
func (w *FilespaceWrapper) IsFile(srcPath string) bool {
	return w.fs.IsFile(w.basePath + srcPath)
}

// IsDir return true if node exist and is a directory
func (w *FilespaceWrapper) IsDir(srcPath string) bool {
	return w.fs.IsDir(w.basePath + srcPath)
}

// MkdirAll create directory recursively
func (w *FilespaceWrapper) MkdirAll(destPath string, filemode os.FileMode) (err error) {
	return w.fs.MkdirAll(w.basePath+destPath, filemode)
}

// Writer return a file node writer
func (w *FilespaceWrapper) Writer(destPath string) (writer filesystem.Writer, err error) {
	return w.fs.Writer(w.basePath + destPath)
}

// Reader return a file node reader
func (w *FilespaceWrapper) Reader(srcPath string) (reader filesystem.Reader, err error) {
	return w.fs.Reader(w.basePath + srcPath)
}

// ReadFile return file data
func (w *FilespaceWrapper) ReadFile(srcPath string) (data []byte, err error) {
	return w.fs.ReadFile(w.basePath + srcPath)
}

// WriteFile write file data
func (w *FilespaceWrapper) WriteFile(destPath string, data []byte, filemode os.FileMode) (err error) {
	return w.fs.WriteFile(w.basePath+destPath, data, filemode)
}

// Filespace get directory node and return it as filespace
func (w *FilespaceWrapper) Filespace(srcPath string) (filesystem.Filespace, error) {
	return NewFilespaceWrapper(w.fs, w.basePath+srcPath), nil
}

// Remove delete node by path
func (w *FilespaceWrapper) Remove(nodePath string) error {
	return w.fs.Remove(w.basePath + nodePath)
}

// RemoveAll delete node by path recursively
func (w *FilespaceWrapper) RemoveAll(nodePath string) error {
	return w.fs.RemoveAll(w.basePath + nodePath)
}

// Lstat returns a FileInfo describing the named file.
func (w *FilespaceWrapper) Lstat(nodePath string) (os.FileInfo, error) {
	return w.fs.Lstat(w.basePath + nodePath)
}
