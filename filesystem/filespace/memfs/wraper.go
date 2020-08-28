package memfs

import (
	"os"

	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/varutil"
)

// FilespaceWrapper is memory filespace wraper
type FilespaceWrapper struct {
	basePath string
	fs       filesystem.Filespace
}

// NewFilespaceWrapper create new memory filespace  wrapper instance
func NewFilespaceWrapper(fs filesystem.Filespace, basePath string) (wrapFS filesystem.Filespace, err error) {
	if basePath, err = varutil.ReduceAbsPath(basePath); err != nil {
		return nil, err
	}
	return &FilespaceWrapper{
		fs:       fs,
		basePath: basePath + "/",
	}, nil
}

// Copy duplicate a file or directory
func (w *FilespaceWrapper) Copy(src, dest string) (err error) {
	if src, err = varutil.ReduceAbsPath(src); err != nil {
		return err
	}
	if dest, err = varutil.ReduceAbsPath(dest); err != nil {
		return err
	}
	return w.fs.Copy(w.basePath+src, w.basePath+dest)
}

// CopyDirectory duplicate a directory
func (w *FilespaceWrapper) CopyDirectory(src, dest string) (err error) {
	if src, err = varutil.ReduceAbsPath(src); err != nil {
		return err
	}
	if dest, err = varutil.ReduceAbsPath(dest); err != nil {
		return err
	}
	return w.fs.CopyDirectory(w.basePath+src, w.basePath+dest)
}

// CopyFile duplicate a file
func (w *FilespaceWrapper) CopyFile(src, dest string) (err error) {
	if src, err = varutil.ReduceAbsPath(src); err != nil {
		return err
	}
	if dest, err = varutil.ReduceAbsPath(dest); err != nil {
		return err
	}
	return w.fs.CopyFile(w.basePath+src, w.basePath+dest)
}

// ReadDir return directory nodes
func (w *FilespaceWrapper) ReadDir(src string) (nodes []os.FileInfo, err error) {
	if src, err = varutil.ReduceAbsPath(src); err != nil {
		return nil, err
	}
	return w.fs.ReadDir(w.basePath + src)
}

// IsExist return true if node exist
func (w *FilespaceWrapper) IsExist(src string) bool {
	var err error
	if src, err = varutil.ReduceAbsPath(src); err != nil {
		return false
	}
	return w.fs.IsExist(w.basePath + src)
}

// IsFile return true if node exist and is a file
func (w *FilespaceWrapper) IsFile(src string) bool {
	var err error
	if src, err = varutil.ReduceAbsPath(src); err != nil {
		return false
	}
	return w.fs.IsFile(w.basePath + src)
}

// IsDir return true if node exist and is a directory
func (w *FilespaceWrapper) IsDir(src string) bool {
	var err error
	if src, err = varutil.ReduceAbsPath(src); err != nil {
		return false
	}
	return w.fs.IsDir(w.basePath + src)
}

// MkdirAll create directory recursively
func (w *FilespaceWrapper) MkdirAll(path string, filemode os.FileMode) (err error) {
	if path, err = varutil.ReduceAbsPath(path); err != nil {
		return err
	}
	return w.fs.MkdirAll(w.basePath+path, filemode)
}

// Writer return a file node writer
func (w *FilespaceWrapper) Writer(path string) (writer filesystem.Writer, err error) {
	if path, err = varutil.ReduceAbsPath(path); err != nil {
		return nil, err
	}
	return w.fs.Writer(w.basePath + path)
}

// Reader return a file node reader
func (w *FilespaceWrapper) Reader(path string) (reader filesystem.Reader, err error) {
	if path, err = varutil.ReduceAbsPath(path); err != nil {
		return nil, err
	}
	return w.fs.Reader(w.basePath + path)
}

// ReadFile return file data
func (w *FilespaceWrapper) ReadFile(src string) (data []byte, err error) {
	if src, err = varutil.ReduceAbsPath(src); err != nil {
		return nil, err
	}
	return w.fs.ReadFile(w.basePath + src)
}

// WriteFile write file data
func (w *FilespaceWrapper) WriteFile(path string, data []byte, filemode os.FileMode) (err error) {
	if path, err = varutil.ReduceAbsPath(path); err != nil {
		return err
	}
	return w.fs.WriteFile(w.basePath+path, data, filemode)
}

// Filespace get directory node and return it as filespace
func (w *FilespaceWrapper) Filespace(path string) (childFS filesystem.Filespace, err error) {
	if path, err = varutil.ReduceAbsPath(path); err != nil {
		return nil, err
	}
	return NewFilespaceWrapper(w.fs, w.basePath+path)
}

// Remove delete node by path
func (w *FilespaceWrapper) Remove(path string) (err error) {
	if path, err = varutil.ReduceAbsPath(path); err != nil {
		return err
	}
	return w.fs.Remove(w.basePath + path)
}

// RemoveAll delete node by path recursively
func (w *FilespaceWrapper) RemoveAll(path string) (err error) {
	if path, err = varutil.ReduceAbsPath(path); err != nil {
		return err
	}
	return w.fs.RemoveAll(w.basePath + path)
}

// Lstat returns a FileInfo describing the named file.
func (w *FilespaceWrapper) Lstat(path string) (info os.FileInfo, err error) {
	if path, err = varutil.ReduceAbsPath(path); err != nil {
		return nil, err
	}
	return w.fs.Lstat(w.basePath + path)
}
