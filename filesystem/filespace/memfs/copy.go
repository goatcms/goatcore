package memfs

import "os"

// copyFile copy a file and return copied file instance
func copyFile(f *File, newName string) (*File, error) {
	f.dataMU.RLock()
	defer f.dataMU.RUnlock()
	var datacopy = make([]byte, len(f.data))
	copy(datacopy[:], f.data)
	return NewFile(newName, f.filemode, f.time, datacopy), nil
}

// copyDir copy directory and return new directories and files tree
func copyDir(d *Dir, newName string) (*Dir, error) {
	var err error
	d.mu.RLock()
	defer d.mu.RUnlock()
	var nodescopy = make([]os.FileInfo, len(d.nodes))
	for i := 0; i < len(d.nodes); i++ {
		if d.nodes[i].IsDir() {
			var dir = d.nodes[i].(*Dir)
			if nodescopy[i], err = copyDir(dir, dir.Name()); err != nil {
				return nil, err
			}
		} else {
			var file = d.nodes[i].(*File)
			if nodescopy[i], err = copyFile(file, file.Name()); err != nil {
				return nil, err
			}
		}
	}
	return NewDir(newName, d.filemode, d.time, nodescopy), nil
}

// copyNode copy directory or file and return it
func copyNode(node os.FileInfo, newName string) (result os.FileInfo, err error) {
	if node.IsDir() {
		return copyDir(node.(*Dir), newName)
	}
	return copyFile(node.(*File), newName)
}
