package memfs

import (
	"fmt"
	"io"
	"os"
	"path"

	"github.com/goatcms/goat-core/filesystem"
)

const (
	currentDir = "."
)

type Filespace struct {
	root *Dir
}

func NewFilespace() (*Filespace, error) {
	return &Filespace{
		root: &Dir{
			nodes: []os.FileInfo{},
		},
	}, nil
}

func (fs *Filespace) Copy(src, dest string) error {
	srcNode, err := fs.root.GetByPath(src)
	if err != nil {
		return err
	}
	if srcNode.IsDir() {
		return fs.CopyDirectory(src, dest)
	} else {
		return fs.CopyFile(src, dest)
	}
}

func (fs *Filespace) CopyDirectory(src, dest string) error {
	srcNode, err := fs.root.GetByPath(src)
	if err != nil {
		return err
	}
	if !srcNode.IsDir() {
		return fmt.Errorf("Source node must be a directory")
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
		return fmt.Errorf("Destination path exist and there is a file " + dest)
	}
	copiedDir.name = path.Base(dest)
	var destDir *Dir = destNode.(*Dir)
	destDir.AddNode(copiedDir)
	return nil
}

func (fs *Filespace) CopyFile(src, dest string) error {
	srcNode, err := fs.root.GetByPath(src)
	if err != nil {
		return err
	}
	if srcNode.IsDir() {
		return fmt.Errorf("Source node must be a file")
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
		return fmt.Errorf("Destination path exist and there is a file " + dest)
	}
	copiedFile.name = path.Base(dest)
	var destDir *Dir = destNode.(*Dir)
	destDir.AddNode(copiedFile)
	return nil
}

func (fs *Filespace) ReadDir(subPath string) ([]os.FileInfo, error) {
	srcNode, err := fs.root.GetByPath(subPath)
	if err != nil {
		return nil, err
	}
	if !srcNode.IsDir() {
		return nil, fmt.Errorf(subPath + " is not a directory")
	}
	var dir *Dir = srcNode.(*Dir)
	return dir.GetNodes(), nil
}

func (fs *Filespace) IsExist(subPath string) bool {
	srcNode, err := fs.root.GetByPath(subPath)
	if err != nil || srcNode == nil {
		return false
	}
	return true
}

func (fs *Filespace) IsFile(subPath string) bool {
	return !fs.IsDir(subPath)
}

func (fs *Filespace) IsDir(subPath string) bool {
	srcNode, err := fs.root.GetByPath(subPath)
	if err != nil || srcNode == nil {
		return false
	}
	return srcNode.IsDir()
}

func (fs *Filespace) MkdirAll(subPath string, filemode os.FileMode) error {
	return fs.root.MkdirAll(subPath, filemode)
}

func (fs *Filespace) Writer(subPath string) (io.Writer, error) {
	if err := fs.root.WriteFile(subPath, []byte{}, filesystem.DefaultUnixFileMode); err != nil {
		return nil, err
	}
	node, err := fs.root.GetByPath(subPath)
	if err != nil {
		return nil, err
	}
	if node.IsDir() {
		return nil, fmt.Errorf("node is a dir %v", node)
	}
	return node.(io.Writer), nil
}

func (fs *Filespace) Reader(subPath string) (io.Reader, error) {
	node, err := fs.root.GetByPath(subPath)
	if err != nil {
		return nil, err
	}
	if node.IsDir() {
		return nil, fmt.Errorf("node is a dir %v", node)
	}
	file := node.(*File)
	file.ResetPointer()
	return io.Reader(file), nil
}

func (fs *Filespace) ReadFile(subPath string) ([]byte, error) {
	return fs.root.ReadFile(subPath)
}

func (fs *Filespace) WriteFile(subPath string, data []byte, perm os.FileMode) error {
	return fs.root.WriteFile(subPath, data, perm)
}

func (fs *Filespace) Filespace(subPath string) (filesystem.Filespace, error) {
	node, err := fs.root.GetByPath(subPath)
	if err != nil {
		return nil, err
	}
	if !node.IsDir() {
		return nil, fmt.Errorf("Path is not a directory " + subPath)
	}
	return filesystem.Filespace(&Filespace{
		root: node.(*Dir),
	}), nil
}

func (fs *Filespace) DebugPrint() {
	debugPrint("", fs.root)
}

func debugPrint(basePath string, d *Dir) {
	for _, node := range d.GetNodes() {
		if node.IsDir() {
			debugPrint(basePath+"/"+node.Name(), node.(*Dir))
		} else {
			fmt.Println("### ", basePath+"/"+node.Name(), ":")
			var file *File = node.(*File)
			fmt.Println(string(file.GetData()))
		}
	}
}
