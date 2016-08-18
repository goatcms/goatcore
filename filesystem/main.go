package filesystem

import (
	"os"
)

const (
	// FilespaceID ia a default dependency name
	FilespaceID = "Filespace"
	// DefaultUnixFileMode is a default file mode for unix base filesystems
	DefaultUnixFileMode = 0644
	// DefaultUnixDirMode is a default dir mode for unix base filesystems
	DefaultUnixDirMode = 0644
)

// Filespace is a abstract filesystem interface
type Filespace interface {
	Copy(src, dest string) error
	CopyDirectory(src, dest string) error
	CopyFile(src, dest string) error
	ReadDir(path string) ([]os.FileInfo, error)
	IsExist(subPath string) bool
	IsFile(subPath string) bool
	IsDir(subPath string) bool
	MkdirAll(subPath string, filemode os.FileMode) error
	ReadFile(subPath string) ([]byte, error)
	WriteFile(subPath string, data []byte, perm os.FileMode) error
	Filespace(subPath string) (Filespace, error)
}

// LoopOn is a callback type trigged on a file or directory
type LoopOn func(fs Filespace, subPath string, info os.FileInfo) error

// LoopFilter is a callback type which is used to filter filespace tree
type LoopFilter func(fs Filespace, subPath string, info os.FileInfo) bool

// Loop is standard loop interface
type Loop interface {
	OnFile(LoopOn)
	OnDir(LoopOn)
	Filter(LoopFilter)
	Run(fs Filespace) error
}
