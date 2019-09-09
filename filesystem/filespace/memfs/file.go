package memfs

import (
	"os"
	"sync"
	"time"
)

// File is a single file
type File struct {
	sync.RWMutex
	name     string
	filemode os.FileMode
	time     time.Time
	data     []byte
	dataMU   sync.RWMutex
}

// NewFile create new File instance
func NewFile(name string, filemode os.FileMode, t time.Time, data []byte) *File {
	return &File{
		name:     name,
		filemode: filemode,
		time:     t,
		data:     data,
	}
}

// Name is a file name
func (f *File) Name() string {
	return f.name
}

// Mode is a unix file/directory mode
func (f *File) Mode() os.FileMode {
	return f.filemode
}

// ModTime is modification time
func (f *File) ModTime() time.Time {
	return f.time
}

// Sys return native system object
func (f *File) Sys() interface{} {
	return nil
}

// Size is length in bytes for regular files; system-dependent for others
func (f *File) Size() int64 {
	return int64(len(f.data))
}

// IsDir return true if node is a directory
func (f *File) IsDir() bool {
	return false
}

// getData return file data bytes
func (f *File) getData() []byte {
	f.dataMU.RLock()
	defer f.dataMU.RUnlock()
	return f.data
}

// setData set new file data bytes
func (f *File) setData(data []byte) {
	f.dataMU.Lock()
	defer f.dataMU.Unlock()
	f.time = time.Now()
	f.data = data
}
