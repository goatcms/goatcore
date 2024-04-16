package diskfs

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/goatcms/goatcore/varutil"

	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/filesystem/disk"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

// Filespace is a files set on local filesystem
type Filespace struct {
	path string
}

// NewFilespace create new Filespace instance
func NewFilespace(path string) (fs filesystem.Filespace, err error) {
	if path, err = filepath.Abs(path); err != nil {
		return nil, err
	}
	path += "/"
	return filesystem.Filespace(&Filespace{
		path: path,
	}), nil
}

// Copy a file or a directory inside the filespace
func (fs *Filespace) Copy(src, dest string) (err error) {
	if src, err = varutil.ReduceAbsPath(src); err != nil {
		return err
	}
	if dest, err = varutil.ReduceAbsPath(dest); err != nil {
		return err
	}
	return disk.Copy(fs.path+src, fs.path+dest)
}

// CopyDirectory duplicate a directory
func (fs *Filespace) CopyDirectory(src, dest string) (err error) {
	if src, err = varutil.ReduceAbsPath(src); err != nil {
		return err
	}
	if dest, err = varutil.ReduceAbsPath(dest); err != nil {
		return err
	}
	return disk.CopyDirectory(fs.path+src, fs.path+dest)
}

// CopyFile duplicate a file
func (fs *Filespace) CopyFile(src, dest string) (err error) {
	if src, err = varutil.ReduceAbsPath(src); err != nil {
		return err
	}
	if dest, err = varutil.ReduceAbsPath(dest); err != nil {
		return err
	}
	return disk.CopyFile(fs.path+src, fs.path+dest)
}

// ReadDir return directory nodes
func (fs *Filespace) ReadDir(subPath string) (infos []os.FileInfo, err error) {
	if subPath, err = varutil.ReduceAbsPath(subPath); err != nil {
		return nil, err
	}
	return ioutil.ReadDir(fs.path + subPath)
}

// IsExist return true if node exist
func (fs *Filespace) IsExist(subPath string) bool {
	var err error
	if subPath, err = varutil.ReduceAbsPath(subPath); err != nil {
		return false
	}
	return disk.IsExist(fs.path + subPath)
}

// IsFile return true if node exist and is a file
func (fs *Filespace) IsFile(src string) bool {
	var err error
	if src, err = varutil.ReduceAbsPath(src); err != nil {
		return false
	}
	return disk.IsFile(fs.path + src)
}

// IsDir return true if node exist and is a directory
func (fs *Filespace) IsDir(src string) bool {
	var err error
	if src, err = varutil.ReduceAbsPath(src); err != nil {
		return false
	}
	return disk.IsDir(fs.path + src)
}

// MkdirAll create directory recursively
func (fs *Filespace) MkdirAll(path string, filemode os.FileMode) (err error) {
	if path, err = varutil.ReduceAbsPath(path); err != nil {
		return err
	}
	return disk.MkdirAll(fs.path+path, filemode)
}

// Writer return a file node writer
func (fs *Filespace) Writer(path string) (w filesystem.Writer, err error) {
	var file *os.File
	if path, err = varutil.ReduceAbsPath(path); err != nil {
		return nil, err
	}
	if file, err = os.OpenFile(fs.path+path, os.O_WRONLY|os.O_CREATE, filesystem.DefaultUnixFileMode); err != nil {
		return nil, err
	}
	return NewFileHandler(file), nil
}

// Reader return a file node reader
func (fs *Filespace) Reader(path string) (r filesystem.Reader, err error) {
	if path, err = varutil.ReduceAbsPath(path); err != nil {
		return nil, err
	}
	return os.OpenFile(fs.path+path, os.O_RDONLY, filesystem.DefaultUnixFileMode)
}

// ReadFile return file data
func (fs *Filespace) ReadFile(path string) (data []byte, err error) {
	if path, err = varutil.ReduceAbsPath(path); err != nil {
		return nil, err
	}
	return ioutil.ReadFile(fs.path + path)
}

// WriteFile write file data
func (fs *Filespace) WriteFile(path string, data []byte, perm os.FileMode) (err error) {
	if path, err = varutil.ReduceAbsPath(path); err != nil {
		return err
	}
	fullPath := fs.path + path
	if err = disk.MkdirAll(filepath.Dir(fullPath), filesystem.DefaultUnixDirMode); err != nil {
		return err
	}
	return ioutil.WriteFile(fullPath, data, perm)
}

// Remove delete node by path
func (fs *Filespace) Remove(path string) (err error) {
	if path, err = varutil.ReduceAbsPath(path); err != nil {
		return err
	}
	return os.Remove(fs.path + path)
}

// RemoveAll delete node by path recursively
func (fs *Filespace) RemoveAll(path string) (err error) {
	if path, err = varutil.ReduceAbsPath(path); err != nil {
		return err
	}
	return os.RemoveAll(fs.path + path)
}

// Filespace get directory node and return it as filespace
func (fs *Filespace) Filespace(path string) (result filesystem.Filespace, err error) {
	if path, err = varutil.ReduceAbsPath(path); err != nil {
		return nil, err
	}
	fullPath := fs.path + path
	if !disk.IsDir(fullPath) {
		return nil, goaterr.Errorf("Path is not a directory %s", path)
	}
	return NewFilespace(fullPath)
}

// Lstat returns a FileInfo describing the named file.
func (fs *Filespace) Lstat(path string) (info os.FileInfo, err error) {
	if path, err = varutil.ReduceAbsPath(path); err != nil {
		return nil, err
	}
	return os.Lstat(fs.path + path)
}

// LocalPath return path in local filesystem
func (fs *Filespace) LocalPath() string {
	return fs.path
}
