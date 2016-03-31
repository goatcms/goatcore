package filespace

import (
	"github.com/goatcms/goat-core/dependency"
	"github.com/goatcms/goat-core/filesystem"
	"github.com/goatcms/goat-core/varutil"
)

const (
	FilespaceName = "Filespace"
)

type Filespace interface {
	Copy(src, dest string) error
	CopyDirectory(src, dest string, filter filesystem.LoopFilter) error
	CopyFile(src, dest string) error
	RunLoop(loop *filesystem.DirLoop, subPath string) error
	IsExist(subPath string) bool
	IsFile(subPath string) bool
	IsDir(subPath string) bool
	MkdirAll(subPath string) error
	Path(subPath string) string
}

func NewFilespace(path string) (Filespace, error){
	varutil.FixDirPath(&path)
	return Filespace(&DefaultFilespace{
		path: path,
	}), nil
}

type DefaultFilespace struct {
	path string
}

func (fs *DefaultFilespace) Copy(src, dest string) error {
	return filesystem.Copy(fs.path+src, fs.path+dest)
}

func (fs *DefaultFilespace) CopyDirectory(src, dest string, filter filesystem.LoopFilter) error {
	return filesystem.CopyDirectory(fs.path+src, fs.path+dest, filter)
}

func (fs *DefaultFilespace) CopyFile(src, dest string) error {
	return filesystem.CopyFile(fs.path+src, fs.path+dest)
}

func (fs *DefaultFilespace) RunLoop(loop *filesystem.DirLoop, subPath string) error {
	return loop.Run(fs.path + subPath)
}

func (fs *DefaultFilespace) IsExist(subPath string) bool {
	return filesystem.IsExist(fs.path + subPath)
}

func (fs *DefaultFilespace) IsFile(subPath string) bool {
	return filesystem.IsFile(fs.path + subPath)
}

func (fs *DefaultFilespace) IsDir(subPath string) bool {
	return filesystem.IsDir(fs.path + subPath)
}

func (fs *DefaultFilespace) MkdirAll(subPath string) error {
	return filesystem.MkdirAll(fs.path + subPath)
}

func (fs *DefaultFilespace) Path(subPath string) string {
	return  fs.path + subPath
}

func BuildFilespaceFactory(path string) dependency.Factory {
	return func(dp dependency.Provider) (dependency.Instance, error) {
		return NewFilespace(path)
	}
}
