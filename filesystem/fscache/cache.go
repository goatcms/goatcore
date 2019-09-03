package fscache

import (
	"os"

	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/filesystem/filespace/memfs"
	"github.com/goatcms/goatcore/filesystem/fshelper"
	"github.com/goatcms/goatcore/varutil"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

// Cache synchronise buffer filesystem to remote filesystem
type Cache struct {
	bufferFS filesystem.Filespace
	bufferRO filesystem.Filespace
	remoteFS filesystem.Filespace
	changes  cacheHistory
}

// cacheHistory storage fiflesystem changes
type cacheHistory struct {
	remove    map[string]bool
	removeAll map[string]bool
	mkdirAll  map[string]os.FileMode
	write     map[string]bool
}

// newCache create new chache for remoteFS (use exists buffer filespace)
func newCache(bufferFS, remoteFS filesystem.Filespace) *Cache {
	return &Cache{
		bufferFS: bufferFS,
		bufferRO: fshelper.NewReadonlyFS(bufferFS),
		remoteFS: remoteFS,
		changes: cacheHistory{
			remove:    map[string]bool{},
			removeAll: map[string]bool{},
			mkdirAll:  map[string]os.FileMode{},
			write:     map[string]bool{},
		},
	}
}

// NewMemCache create new chache for remote filespace (storage changes into memory)
func NewMemCache(remoteFS filesystem.Filespace) (c *Cache, err error) {
	var bufferFS filesystem.Filespace
	if bufferFS, err = memfs.NewFilespace(); err != nil {
		return c, err
	}
	return newCache(bufferFS, remoteFS), nil
}

// Buffer return RO (Reado-Only) buffer filespace
func (c Cache) Buffer() filesystem.Filespace {
	return c.bufferRO
}

// Commit send buffered changes to remote filesystem
func (c Cache) Commit() (err error) {
	var (
		src      string
		filemode os.FileMode
	)
	for src = range c.changes.remove {
		if c.remoteFS.IsFile(src) {
			if err = c.remoteFS.Remove(src); err != nil {
				return err
			}
		}
	}
	for src = range c.changes.removeAll {
		if c.remoteFS.IsExist(src) {
			if err = c.remoteFS.RemoveAll(src); err != nil {
				return err
			}
		}
	}
	for src, filemode = range c.changes.mkdirAll {
		if c.bufferFS.IsDir(src) {
			if err = c.remoteFS.MkdirAll(src, filemode); err != nil {
				return err
			}
		}
	}
	for src = range c.changes.write {
		if c.bufferFS.IsFile(src) {
			if err = fshelper.StreamCopy(c.bufferFS, c.remoteFS, src); err != nil {
				return err
			}
		}
	}
	return nil
}

// Copy duplicate a file or directory
func (c Cache) srcFS(p string) (srcFS filesystem.Filespace, src string) {
	src = varutil.CleanPath(p)
	if c.bufferFS.IsExist(src) {
		srcFS = c.bufferFS
	} else {
		srcFS = c.remoteFS
	}
	return srcFS, src
}

// Copy duplicate a file or directory
func (c Cache) Copy(src, dest string) error {
	var srcFS filesystem.Filespace
	srcFS, src = c.srcFS(src)
	dest = varutil.CleanPath(dest)
	c.changes.write[dest] = true
	return (fshelper.Copier{
		SrcFS:    srcFS,
		SrcPath:  src,
		DestFS:   c.bufferFS,
		DestPath: dest,
	}).Do()
}

// CopyDirectory duplicate a directory
func (c Cache) CopyDirectory(src, dest string) error {
	var srcFS filesystem.Filespace
	srcFS, src = c.srcFS(src)
	dest = varutil.CleanPath(dest)
	if !srcFS.IsDir(src) {
		return goaterr.Errorf("Source node must be a directory")
	}
	c.changes.write[dest] = true
	return c.Copy(src, dest)
}

// CopyFile duplicate a file
func (c Cache) CopyFile(src, dest string) error {
	var srcFS filesystem.Filespace
	srcFS, src = c.srcFS(src)
	dest = varutil.CleanPath(dest)
	if !srcFS.IsFile(src) {
		return goaterr.Errorf("Source node must be a file")
	}
	c.changes.write[dest] = true
	return c.Copy(src, dest)
}

// ReadDir return directory nodes
func (c Cache) ReadDir(src string) (result []os.FileInfo, err error) {
	var (
		remoteDirs, bufferDirs []os.FileInfo
		remoteErr, bufferErr   error
	)
	src = varutil.CleanPath(src)
	remoteDirs, remoteErr = c.remoteFS.ReadDir(src)
	bufferDirs, bufferErr = c.bufferFS.ReadDir(src)
	if remoteErr != nil && bufferErr != nil {
		return nil, goaterr.ToErrors(goaterr.AppendError(nil, remoteErr, bufferErr))
	}
	result = remoteDirs
ReadDirLoop:
	for _, bnode := range bufferDirs {
		for _, cnode := range remoteDirs {
			if bnode.Name() == cnode.Name() {
				continue ReadDirLoop
			}
		}
		result = append(result, bnode)
	}
	return result, nil
}

// IsExist return true if node exist
func (c Cache) IsExist(src string) bool {
	src = varutil.CleanPath(src)
	return c.bufferFS.IsExist(src) || c.remoteFS.IsExist(src)
}

// IsFile return true if node exist and is a file
func (c Cache) IsFile(src string) bool {
	src = varutil.CleanPath(src)
	return c.bufferFS.IsFile(src) || c.remoteFS.IsFile(src)
}

// IsDir return true if node exist and is a directory
func (c Cache) IsDir(src string) bool {
	src = varutil.CleanPath(src)
	return c.bufferFS.IsDir(src) || c.remoteFS.IsDir(src)
}

// MkdirAll create directory recursively
func (c Cache) MkdirAll(dest string, filemode os.FileMode) error {
	dest = varutil.CleanPath(dest)
	c.changes.mkdirAll[dest] = filemode
	return c.bufferFS.MkdirAll(dest, filemode)
}

// Writer return a file node writer
func (c Cache) Writer(dest string) (filesystem.Writer, error) {
	dest = varutil.CleanPath(dest)
	c.changes.write[dest] = true
	return c.bufferFS.Writer(dest)
}

// Reader return a file node reader
func (c Cache) Reader(src string) (filesystem.Reader, error) {
	var srcFS filesystem.Filespace
	srcFS, src = c.srcFS(src)
	return srcFS.Reader(src)
}

// ReadFile return file data
func (c Cache) ReadFile(src string) ([]byte, error) {
	var srcFS filesystem.Filespace
	srcFS, src = c.srcFS(src)
	return srcFS.ReadFile(src)
}

// WriteFile write file data
func (c Cache) WriteFile(dest string, data []byte, perm os.FileMode) error {
	c.changes.write[dest] = true
	return c.bufferFS.WriteFile(dest, data, perm)
}

// Filespace get directory node and return it as filespace
func (c Cache) Filespace(subPath string) (filesystem.Filespace, error) {
	return fshelper.NewSubFS(c, subPath), nil
}

// Remove delete node by path
func (c Cache) Remove(dest string) (err error) {
	dest = varutil.CleanPath(dest)
	if c.bufferFS.IsExist(dest) {
		err = c.bufferFS.Remove(dest)
	}
	c.changes.remove[dest] = true
	return err
}

// RemoveAll delete node by path recursively
func (c Cache) RemoveAll(dest string) (err error) {
	dest = varutil.CleanPath(dest)
	if c.bufferFS.IsExist(dest) {
		err = c.bufferFS.RemoveAll(dest)
	}
	c.changes.removeAll[dest] = true
	return err
}

// Lstat returns a FileInfo describing the named file.
func (c Cache) Lstat(src string) (os.FileInfo, error) {
	var srcFS filesystem.Filespace
	srcFS, src = c.srcFS(src)
	return srcFS.Lstat(src)
}
