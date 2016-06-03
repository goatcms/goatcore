package diskfs

import (
	"github.com/goatcms/goat-core/dependency"
	"github.com/goatcms/goat-core/filesystem"
	"github.com/goatcms/goat-core/filesystem/disk"
	"github.com/goatcms/goat-core/varutil"
	"io/ioutil"
)

type Filespace struct {
	path string
}

func NewFilespace(path string) (filesystem.Filespace, error) {
	varutil.FixDirPath(&path)
	return filesystem.Filespace(&DefaultFilespace{
		path: path,
	}), nil
}

func (fs *Filespace) Copy(src, dest string) error {
	return disk.Copy(fs.path+src, fs.path+dest)
}

func (fs *Filespace) CopyDirectory(src, dest string, filter filesystem.LoopFilter) error {
	return disk.CopyDirectory(fs.path+src, fs.path+dest, filter)
}

func (fs *Filespace) CopyFile(src, dest string) error {
	return disk.CopyFile(fs.path+src, fs.path+dest)
}

func (fs *Filespace) ReadDir(subPath string) ([]os.FileInfo, error) {
	return ioutil.ReadDir(fs.path + subPath)
}

func (fs *Filespace) IsExist(subPath string) bool {
	return filesystem.IsExist(fs.path + subPath)
}

func (fs *Filespace) IsFile(subPath string) bool {
	return filesystem.IsFile(fs.path + subPath)
}

func (fs *Filespace) IsDir(subPath string) bool {
	return filesystem.IsDir(fs.path + subPath)
}

func (fs *Filespace) MkdirAll(subPath string, filemode os.FileMode) error {
	return filesystem.MkdirAll(fs.path+subPath, filemode)
}

func (fs *Filespace) ReadFile(subPath string) ([]byte, error) {
	return ioutil.ReadFile(fs.path + subPath)
}

func (fs *Filespace) WriteFile(subPath string, data []byte, perm os.FileMode) error {
	return filesystem.WriteFile(fs.path+subPath, data, perm)
}

func (fs *Filespace) Filespace(subPath string) (filesystem.Filespace, error) {
	if !fs.IsDir(subPath) {
		return nil, fmt.Errorf("Path is not a directory " + fs.path + subPath)
	}
	return NewFilespace(fs.path + subPath), nil
}
