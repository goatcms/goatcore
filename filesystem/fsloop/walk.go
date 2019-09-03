package fsloop

import (
	"os"

	"github.com/goatcms/goatcore/filesystem"
)

// WalkFunc is Walk callback
type WalkFunc func(path string, info os.FileInfo) (err error)

// WalkFS walks the file tree rooted at root, calling walkFn for each file or directory in the tree,
// including root.
func WalkFS(fs filesystem.Filespace, root string, fileFunc WalkFunc, dirFunc WalkFunc) (err error) {
	var infos []os.FileInfo
	if infos, err = fs.ReadDir(root); err != nil {
		return err
	}
	for _, info := range infos {
		currentPath := root + "/" + info.Name()
		if info.IsDir() {
			if info.Name() == "." || info.Name() == ".." {
				continue
			}
			if err = WalkFS(fs, currentPath, fileFunc, dirFunc); err != nil {
				return err
			}
			if dirFunc != nil {
				if err = dirFunc(currentPath, info); err != nil {
					return err
				}
			}
		} else {
			if err = fileFunc(currentPath, info); err != nil {
				return err
			}
		}
	}
	return nil
}
