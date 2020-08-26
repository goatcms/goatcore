package memfs

import (
	"fmt"
	"os"
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
func NewFilespace() (filesystem.Filespace, error) {
	return &Filespace{
		root: NewDir("ROOT", filesystem.DefaultUnixDirMode, time.Now(), []os.FileInfo{}),
	}, nil
}

// Copy duplicate a file or directory
func (fs *Filespace) Copy(src, dest string) (err error) {
	var (
		destDirPath  []string
		destNodeName string
		destDir      *Dir
		srcNode      os.FileInfo
		copiedNode   os.FileInfo
	)
	src = varutil.CleanPath(src)
	dest = varutil.CleanPath(dest)
	if destDirPath, destNodeName, err = splitContainsPath(dest); err != nil {
		return err
	}
	if srcNode, err = getNodeByPath(fs.root, src); err != nil {
		return err
	}
	if destDir, err = mkdirAllNodes(fs.root, destDirPath, filesystem.DefaultUnixDirMode); err != nil {
		return err
	}
	if copiedNode, err = copyNode(srcNode, destNodeName); err != nil {
		return err
	}
	return destDir.addNode(copiedNode)
}

// CopyDirectory duplicate a directory
func (fs *Filespace) CopyDirectory(src, dest string) (err error) {
	var (
		destDirPath  []string
		destNodeName string
		destDir      *Dir
		srcDir       *Dir
		copiedDir    *Dir
	)
	src = varutil.CleanPath(src)
	dest = varutil.CleanPath(dest)
	if destDirPath, destNodeName, err = splitContainsPath(dest); err != nil {
		return err
	}
	if srcDir, err = getDirByPath(fs.root, src); err != nil {
		return err
	}
	if destDir, err = mkdirAllNodes(fs.root, destDirPath, filesystem.DefaultUnixDirMode); err != nil {
		return err
	}
	if copiedDir, err = copyDir(srcDir, destNodeName); err != nil {
		return err
	}
	return destDir.addNode(copiedDir)
}

// CopyFile duplicate a file
func (fs *Filespace) CopyFile(src, dest string) (err error) {
	var (
		destDirPath  []string
		destNodeName string
		destDir      *Dir
		srcFile      *File
		copiedFile   *File
	)
	src = varutil.CleanPath(src)
	dest = varutil.CleanPath(dest)
	if destDirPath, destNodeName, err = splitContainsPath(dest); err != nil {
		return err
	}
	if srcFile, err = getFileByPath(fs.root, src); err != nil {
		return err
	}
	if destDir, err = mkdirAllNodes(fs.root, destDirPath, filesystem.DefaultUnixDirMode); err != nil {
		return err
	}
	if copiedFile, err = copyFile(srcFile, destNodeName); err != nil {
		return err
	}
	return destDir.addNode(copiedFile)
}

// ReadDir return directory nodes
func (fs *Filespace) ReadDir(srcPath string) (nodes []os.FileInfo, err error) {
	var srcDir *Dir
	srcPath = varutil.CleanPath(srcPath)
	if srcDir, err = getDirByPath(fs.root, srcPath); err != nil {
		return nil, err
	}
	return srcDir.getNodes(), nil
}

// IsExist return true if node exist
func (fs *Filespace) IsExist(srcPath string) bool {
	if srcNode, err := getNodeByPath(fs.root, srcPath); err != nil || srcNode == nil {
		return false
	}
	return true
}

// IsFile return true if node exist and is a file
func (fs *Filespace) IsFile(srcPath string) bool {
	srcPath = varutil.CleanPath(srcPath)
	if srcNode, err := getFileByPath(fs.root, srcPath); err != nil || srcNode == nil {
		return false
	}
	return true
}

// IsDir return true if node exist and is a directory
func (fs *Filespace) IsDir(srcPath string) bool {
	srcPath = varutil.CleanPath(srcPath)
	if srcNode, err := getDirByPath(fs.root, srcPath); err != nil || srcNode == nil {
		return false
	}
	return true
}

// MkdirAll create directory recursively
func (fs *Filespace) MkdirAll(destPath string, filemode os.FileMode) (err error) {
	destPath = varutil.CleanPath(destPath)
	_, err = mkdirAll(fs.root, destPath, filemode)
	return err
}

// Writer return a file node writer
func (fs *Filespace) Writer(destPath string) (writer filesystem.Writer, err error) {
	var (
		destDirPath  []string
		destNodeName string
		dir          *Dir
		file         *File
		node         os.FileInfo
		ok           bool
	)
	if destPath, err = reducePath(destPath); err != nil {
		return nil, err
	}
	if destDirPath, destNodeName, err = splitContainsPath(destPath); err != nil {
		return nil, err
	}
	if dir, err = mkdirAllNodes(fs.root, destDirPath, filesystem.DefaultUnixDirMode); err != nil {
		return nil, err
	}
	dir.Lock()
	defer dir.Unlock()
	if node, err = dir.getNode(destNodeName); err != nil {
		file = NewFile(destNodeName, filesystem.DefaultUnixFileMode, time.Now(), []byte{})
		if err = dir.addNode(file); err != nil {
			return nil, err
		}
	} else {
		if file, ok = node.(*File); !ok {
			return nil, goaterr.Errorf("Node %s must be a file", destPath)
		}
		file.time = time.Now()
	}
	return NewFileHandler(file), nil
}

// Reader return a file node reader
func (fs *Filespace) Reader(srcPath string) (reader filesystem.Reader, err error) {
	var file *File
	srcPath = varutil.CleanPath(srcPath)
	if file, err = getFileByPath(fs.root, srcPath); err != nil {
		return nil, err
	}
	return NewFileHandler(file), nil
}

// ReadFile return file data
func (fs *Filespace) ReadFile(srcPath string) (data []byte, err error) {
	var file *File
	srcPath = varutil.CleanPath(srcPath)
	if file, err = getFileByPath(fs.root, srcPath); err != nil {
		return nil, err
	}
	return file.getData(), nil
}

// WriteFile write file data
func (fs *Filespace) WriteFile(destPath string, data []byte, filemode os.FileMode) (err error) {
	var (
		destDirPath  []string
		destNodeName string
		dir          *Dir
		file         *File
		node         os.FileInfo
		ok           bool
	)
	destPath = varutil.CleanPath(destPath)
	if destDirPath, destNodeName, err = splitContainsPath(destPath); err != nil {
		return err
	}
	if dir, err = mkdirAllNodes(fs.root, destDirPath, filemode); err != nil {
		return err
	}
	dir.Lock()
	defer dir.Unlock()
	if node, err = dir.getNode(destNodeName); err != nil {
		file = NewFile(destNodeName, filesystem.DefaultUnixFileMode, time.Now(), data)
		return dir.addNode(file)
	}
	if file, ok = node.(*File); !ok {
		return goaterr.Errorf("Node %s must be a file", destPath)
	}
	file.time = time.Now()
	file.setData(data)
	return nil
}

// Filespace get directory node and return it as filespace
func (fs *Filespace) Filespace(basePath string) (filesystem.Filespace, error) {
	return NewFilespaceWrapper(fs, basePath)
}

// Remove delete node by path
func (fs *Filespace) Remove(nodePath string) error {
	nodePath = varutil.CleanPath(nodePath)
	return removeNodeByPath(fs.root, nodePath, true)
}

// RemoveAll delete node by path recursively
func (fs *Filespace) RemoveAll(nodePath string) error {
	nodePath = varutil.CleanPath(nodePath)
	return removeNodeByPath(fs.root, nodePath, false)
}

// Lstat returns a FileInfo describing the named file.
func (fs *Filespace) Lstat(nodePath string) (os.FileInfo, error) {
	nodePath = varutil.CleanPath(nodePath)
	return getNodeByPath(fs.root, nodePath)
}

// DebugPrint print filespace tree
func (fs *Filespace) DebugPrint() {
	debugPrint("", fs.root)
}

func debugPrint(basePath string, d *Dir) {
	for _, node := range d.getNodes() {
		if node.IsDir() {
			debugPrint(basePath+"/"+node.Name(), node.(*Dir))
		} else {
			fmt.Println("### ", basePath+"/"+node.Name(), ":")
			var file = node.(*File)
			fmt.Println(string(file.getData()))
		}
	}
}
