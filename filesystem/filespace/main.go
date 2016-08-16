package filesystem

import (
	"os"
)

const (
	FilespaceKey    = "Filespace"
	DefaultUnixMode = 0644
)

type Filespace interface {
	Copy(src, dest string) error
	CopyDirectory(src, dest string, filter LoopFilter) error
	CopyFile(src, dest string) error
	ReadDir(path string) ([]os.FileInfo, error)
	IsExist(subPath string) bool
	IsFile(subPath string) bool
	IsDir(subPath string) bool
	MkdirAll(subPath string, filemode os.FileMode) error
	//Path(subPath string) string
	ReadFile(subPath string) ([]byte, error)
	WriteFile(subPath string, data []byte, perm os.FileMode) error
	Filespace(subPath string) (Filespace, error)
}

type LoopOn func(fs Filespace, subPath string, info os.FileInfo) error
type LoopFilter func(fs Filespace, subPath string, info os.FileInfo) bool

type Loop interface {
	OnFile(LoopOn)
	OnDir(LoopOn)
	Filter(LoopFilter)
	Run(fs Filespace) error
}
