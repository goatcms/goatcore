package memfs

import (
	"fmt"
	"os"
	"path"
	"strings"
	"time"
)

type Dir struct {
	name     string
	filemode os.FileMode
	time     time.Time
	nodes    []os.FileInfo
}

func (d *Dir) Name() string {
	return d.name
}

func (d *Dir) Mode() os.FileMode {
	return d.filemode
}

func (d *Dir) ModTime() time.Time {
	return d.time
}

func (d *Dir) Sys() interface{} {
	return nil
}

func (d *Dir) Size() int64 {
	return int64(len(d.nodes))
}

func (d *Dir) IsDir() bool {
	return true
}

func (d *Dir) GetNodes() []os.FileInfo {
	return d.nodes
}

func (d *Dir) GetNode(nodeName string) (os.FileInfo, error) {
	for _, node := range d.nodes {
		if nodeName == node.Name() {
			return node, nil
		}
	}
	return nil, fmt.Errorf("No find node with name " + nodeName)
}

func (d *Dir) AddNode(newNode os.FileInfo) error {
	for _, node := range d.nodes {
		if newNode.Name() == node.Name() {
			return fmt.Errorf("node named  " + newNode.Name() + " exists")
		}
	}
	d.nodes = append(d.nodes, newNode)
	return nil
}

func (d *Dir) RemoveNodeByName(name string) error {
	for i := 0; i < len(d.nodes); i++ {
		if d.nodes[i].Name() == name {
			d.nodes = append(d.nodes[:i], d.nodes[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Con not find node to remove (by name " + name + ")")
}

func (d *Dir) GetByPath(nodePath string) (os.FileInfo, error) {
	var currentNode os.FileInfo = d
	if nodePath == "." || nodePath == "./" {
		return currentNode, nil
	}
	pathNodes := strings.Split(path.Clean(nodePath), "/")
	for _, nodeName := range pathNodes {
		if currentNode.IsDir() != true {
			return nil, fmt.Errorf("Node by name " + currentNode.Name() + " must be dir to get sub node (path " + nodePath + " )")
		}
		var dir *Dir = currentNode.(*Dir)
		newNode, err := dir.GetNode(nodeName)
		if err != nil {
			return nil, err
		}
		currentNode = newNode
	}
	return currentNode, nil
}

func (d *Dir) Copy() (*Dir, error) {
	var err error
	var nodescopy = make([]os.FileInfo, len(d.nodes))
	for i := 0; i < len(d.nodes); i++ {
		if d.nodes[i].IsDir() {
			var dir *Dir = d.nodes[i].(*Dir)
			nodescopy[i], err = dir.Copy()
			if err != nil {
				return nil, err
			}
		} else {
			var file *File = d.nodes[i].(*File)
			nodescopy[i], err = file.Copy()
			if err != nil {
				return nil, err
			}
		}
	}
	return &Dir{
		name:     d.name,
		filemode: d.filemode,
		time:     d.time,
		nodes:    nodescopy,
	}, nil
}

func (d *Dir) MkdirAll(subPath string, filemode os.FileMode) error {
	pathNodes := strings.Split(path.Clean(subPath), "/")
	currentNode := d
	for i, nodeName := range pathNodes {
		newCurrentNode, err := currentNode.GetNode(nodeName)
		if err == nil {
			// get exist path node
			if !newCurrentNode.IsDir() {
				return fmt.Errorf(nodeName + " exists and is not dir in path " + subPath)
			}
			currentNode = newCurrentNode.(*Dir)
			continue
		}
		//create directories
		var subDir *Dir
		for i2 := len(pathNodes) - 1; i2 >= i; i2-- {
			newSubDir := &Dir{
				name:     pathNodes[i2],
				filemode: filemode,
				time:     time.Now(),
				nodes:    []os.FileInfo{},
			}
			if subDir != nil {
				newSubDir.AddNode(subDir)
			}
			subDir = newSubDir
		}
		currentNode.AddNode(subDir)
		return nil
	}
	return nil
}

func (d *Dir) ReadFile(subPath string) ([]byte, error) {
	node, err := d.GetByPath(subPath)
	if err != nil {
		return nil, err
	}
	if node.IsDir() {
		return nil, fmt.Errorf("Use ReadFile on directory ")
	}
	var fileNode *File = node.(*File)
	return fileNode.GetData(), nil
}

func (d *Dir) WriteFile(subPath string, data []byte, perm os.FileMode) error {
	dirPath := path.Dir(subPath)
	d.MkdirAll(dirPath, perm)
	node, err := d.GetByPath(subPath)
	if err != nil {
		//create new file if not exist
		node, err := d.GetByPath(dirPath)
		if err != nil {
			return err
		}
		if !node.IsDir() {
			return fmt.Errorf("There is a file on path " + dirPath)
		}
		var baseDir *Dir = node.(*Dir)
		baseDir.AddNode(&File{
			name:     path.Base(subPath),
			filemode: perm,
			time:     time.Now(),
			data:     data,
		})
		return nil
	}
	//overwrite file
	if node.IsDir() {
		return fmt.Errorf("Use WriteFile on directory")
	}
	var file *File = node.(*File)
	file.SetData(data)
	return nil
}
