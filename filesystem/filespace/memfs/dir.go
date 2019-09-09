package memfs

import (
	"os"
	"sync"
	"time"

	"github.com/goatcms/goatcore/varutil/goaterr"
)

// Dir is single directory
type Dir struct {
	sync.RWMutex
	name     string
	filemode os.FileMode
	time     time.Time
	nodes    []os.FileInfo
	index    map[string]os.FileInfo
	mu       sync.RWMutex
}

// NewDir create new directory with nodes
func NewDir(name string, filemode os.FileMode, t time.Time, nodes []os.FileInfo) *Dir {
	dir := &Dir{
		name:     name,
		filemode: filemode,
		time:     t,
		nodes:    nodes,
		index:    map[string]os.FileInfo{},
	}
	for _, node := range dir.nodes {
		dir.index[node.Name()] = node
	}
	return dir
}

// Name is a directory name
func (d *Dir) Name() string {
	return d.name
}

// Mode is a unix file/directory mode
func (d *Dir) Mode() os.FileMode {
	return d.filemode
}

// ModTime is modification time
func (d *Dir) ModTime() time.Time {
	return d.time
}

// Sys return native system object
func (d *Dir) Sys() interface{} {
	return nil
}

// Size is length in bytes for regular files; system-dependent for others
func (d *Dir) Size() int64 {
	return int64(len(d.nodes))
}

// IsDir return true if node is a directory
func (d *Dir) IsDir() bool {
	return true
}

// getNodes return nodes for directory
func (d *Dir) getNodes() []os.FileInfo {
	return d.nodes
}

// getNodes return nodes for directory
func (d *Dir) contains(name string) bool {
	d.mu.RLock()
	defer d.mu.RUnlock()
	_, ok := d.index[name]
	return ok
}

// getNode return single node by name
func (d *Dir) getNode(nodeName string) (node os.FileInfo, err error) {
	var ok bool
	d.mu.RLock()
	defer d.mu.RUnlock()
	if node, ok = d.index[nodeName]; !ok {
		return nil, goaterr.Errorf("No find node with name " + nodeName)
	}
	return node, nil
}

// getDir return single directory node by name
func (d *Dir) getDir(nodeName string) (dir *Dir, err error) {
	var (
		ok   bool
		node os.FileInfo
	)
	d.mu.RLock()
	defer d.mu.RUnlock()
	if node, ok = d.index[nodeName]; !ok {
		return nil, goaterr.Errorf("No find directory node with name %s", nodeName)
	}
	if dir, ok = node.(*Dir); !ok {
		return nil, goaterr.Errorf("%s is not a directory", nodeName)
	}
	return dir, nil
}

// addNode add new node to directory (name must be unique in directory)
func (d *Dir) addNode(newNode os.FileInfo) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	nodeName := newNode.Name()
	if _, ok := d.index[nodeName]; ok {
		return goaterr.Errorf("node named %s exists", nodeName)
	}
	d.nodes = append(d.nodes, newNode)
	d.index[nodeName] = newNode
	return nil
}

// addNode add new node to directory (name must be unique in directory)
func (d *Dir) mkdir(name string, mode os.FileMode) (dir *Dir, err error) {
	var (
		node os.FileInfo
		ok   bool
	)
	if dir, err = d.getDir(name); err == nil {
		return dir, nil
	}
	d.mu.Lock()
	defer d.mu.Unlock()
	if node, ok = d.index[name]; ok {
		if dir, ok = node.(*Dir); !ok {
			return nil, goaterr.Errorf("Node named %s is not directory", name)
		}
		return dir, nil
	}
	dir = NewDir(name, mode, time.Now(), []os.FileInfo{})
	d.nodes = append(d.nodes, dir)
	d.index[name] = dir
	return dir, nil
}

// removeNodeByName remove a node by name
func (d *Dir) removeNodeByName(name string) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	for i := 0; i < len(d.nodes); i++ {
		if d.nodes[i].Name() == name {
			d.nodes = append(d.nodes[:i], d.nodes[i+1:]...)
			delete(d.index, name)
			return nil
		}
	}
	return goaterr.Errorf("Con not find node to remove (by name " + name + ")")
}
