package memfs

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/varutil"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

const (
	currentDir = "."
)

// Filespace is memory filespace
type Filespace struct {
	root *Dir
}

// NewFilespace create new memory filespace instance
func NewFilespace() (*Filespace, error) {
	return &Filespace{
		root: NewDir("", filesystem.DefaultUnixDirMode, time.Now(), []os.FileInfo{}),
	}, nil
}

// Copy duplicate a file or directory
func (fs *Filespace) Copy(src, dest string) error {
	src = varutil.CleanPath(src)
	dest = varutil.CleanPath(dest)
	srcNode, err := fs.root.GetByPath(src)
	if err != nil {
		return err
	}
	if srcNode.IsDir() {
		return fs.CopyDirectory(src, dest)
	}
	return fs.CopyFile(src, dest)
}

// CopyDirectory duplicate a directory
func (fs *Filespace) CopyDirectory(src, dest string) error {
	src = varutil.CleanPath(src)
	dest = varutil.CleanPath(dest)
	srcNode, err := fs.root.GetByPath(src)
	if err != nil {
		return err
	}
	if !srcNode.IsDir() {
		return goaterr.Errorf("Source node must be a directory")
	}
	dir := srcNode.(*Dir)
	copiedDir, err := dir.Copy()
	if err != nil {
		return err
	}
	destDirName := path.Dir(dest)
	err = fs.MkdirAll(destDirName, 0777)
	if err != nil {
		return err
	}
	destNode, err := fs.root.GetByPath(destDirName)
	if err != nil {
		return err
	}
	if !destNode.IsDir() {
		return goaterr.Errorf("Destination path exist and there is a file " + dest)
	}
	copiedDir.name = path.Base(dest)
	var destDir = destNode.(*Dir)
	destDir.AddNode(copiedDir)
	return nil
}

// CopyFile duplicate a file
func (fs *Filespace) CopyFile(src, dest string) error {
	src = varutil.CleanPath(src)
	dest = varutil.CleanPath(dest)
	srcNode, err := fs.root.GetByPath(src)
	if err != nil {
		return err
	}
	if srcNode.IsDir() {
		return goaterr.Errorf("Source node must be a file")
	}
	file := srcNode.(*File)
	copiedFile, err := file.Copy()
	if err != nil {
		return err
	}
	destDirName := path.Dir(dest)
	err = fs.MkdirAll(destDirName, 0777)
	if err != nil {
		return err
	}
	destNode, err := fs.root.GetByPath(destDirName)
	if err != nil {
		return err
	}
	if !destNode.IsDir() {
		return goaterr.Errorf("Destination path exist and there is a file " + dest)
	}
	copiedFile.name = path.Base(dest)
	var destDir = destNode.(*Dir)
	destDir.AddNode(copiedFile)
	return nil
}

// ReadDir return directory nodes
func (fs *Filespace) ReadDir(subPath string) ([]os.FileInfo, error) {
	subPath = varutil.CleanPath(subPath)
	srcNode, err := fs.root.GetByPath(subPath)
	if err != nil {
		return nil, err
	}
	if !srcNode.IsDir() {
		return nil, goaterr.Errorf(subPath + " is not a directory")
	}
	var dir = srcNode.(*Dir)
	return dir.GetNodes(), nil
}

// IsExist return true if node exist
func (fs *Filespace) IsExist(subPath string) bool {
	srcNode, err := fs.root.GetByPath(subPath)
	if err != nil || srcNode == nil {
		return false
	}
	return true
}

// IsFile return true if node exist and is a file
func (fs *Filespace) IsFile(subPath string) bool {
	subPath = varutil.CleanPath(subPath)
	srcNode, err := fs.root.GetByPath(subPath)
	if err != nil || srcNode == nil {
		return false
	}
	return !srcNode.IsDir()
}

// IsDir return true if node exist and is a directory
func (fs *Filespace) IsDir(subPath string) bool {
	subPath = varutil.CleanPath(subPath)
	srcNode, err := fs.root.GetByPath(subPath)
	if err != nil || srcNode == nil {
		return false
	}
	return srcNode.IsDir()
}

// MkdirAll create directory recursively
func (fs *Filespace) MkdirAll(subPath string, filemode os.FileMode) error {
	subPath = varutil.CleanPath(subPath)
	return fs.root.MkdirAll(subPath, filemode)
}

// Writer return a file node writer
func (fs *Filespace) Writer(subPath string) (filesystem.Writer, error) {
	subPath = varutil.CleanPath(subPath)
	if err := fs.root.WriteFile(subPath, []byte{}, filesystem.DefaultUnixFileMode); err != nil {
		return nil, err
	}
	node, err := fs.root.GetByPath(subPath)
	if err != nil {
		return nil, err
	}
	if node.IsDir() {
		return nil, goaterr.Errorf("node is a dir %v", node)
	}
	return node.(filesystem.Writer), nil
}

// Reader return a file node reader
func (fs *Filespace) Reader(subPath string) (filesystem.Reader, error) {
	subPath = varutil.CleanPath(subPath)
	node, err := fs.root.GetByPath(subPath)
	if err != nil {
		return nil, err
	}
	if node.IsDir() {
		return nil, goaterr.Errorf("node is a dir %v", node)
	}
	file := node.(*File)
	file.ResetPointer()
	return filesystem.Reader(file), nil
}

// ReadFile return file data
func (fs *Filespace) ReadFile(subPath string) ([]byte, error) {
	subPath = varutil.CleanPath(subPath)
	return fs.root.ReadFile(subPath)
}

// WriteFile write file data
func (fs *Filespace) WriteFile(subPath string, data []byte, perm os.FileMode) error {
	subPath = varutil.CleanPath(subPath)
	return fs.root.WriteFile(subPath, data, perm)
}

// Filespace get directory node and return it as filespace
func (fs *Filespace) Filespace(subPath string) (filesystem.Filespace, error) {
	subPath = varutil.CleanPath(subPath)
	node, err := fs.root.GetByPath(subPath)
	if err != nil {
		return nil, err
	}
	if !node.IsDir() {
		return nil, goaterr.Errorf("Path is not a directory " + subPath)
	}
	return filesystem.Filespace(&Filespace{
		root: node.(*Dir),
	}), nil
}

// Remove delete node by path
func (fs *Filespace) Remove(subPath string) error {
	subPath = varutil.CleanPath(subPath)
	return fs.root.Remove(subPath, true)
}

// RemoveAll delete node by path recursively
func (fs *Filespace) RemoveAll(subPath string) error {
	subPath = varutil.CleanPath(subPath)
	return fs.root.Remove(subPath, false)
}

// Lstat returns a FileInfo describing the named file.
func (fs *Filespace) Lstat(subPath string) (os.FileInfo, error) {
	subPath = varutil.CleanPath(subPath)
	return fs.root.GetByPath(subPath)
}

// DebugPrint print filespace tree
func (fs *Filespace) DebugPrint() {
	debugPrint("", fs.root)
}

func debugPrint(basePath string, d *Dir) {
	for _, node := range d.GetNodes() {
		if node.IsDir() {
			debugPrint(basePath+"/"+node.Name(), node.(*Dir))
		} else {
			fmt.Println("### ", basePath+"/"+node.Name(), ":")
			var file = node.(*File)
			fmt.Println(string(file.GetData()))
		}
	}
}
