package memfs

import (
	"io"
	"os"
	"time"
)

// File is a single file
type File struct {
	name     string
	filemode os.FileMode
	time     time.Time
	data     []byte
	pointer  int
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

// GetData return file data bytes
func (f *File) GetData() []byte {
	return f.data
}

// SetData set new file data bytes
func (f *File) SetData(data []byte) {
	f.data = data
}

func (f *File) Write(p []byte) (n int, err error) {
	f.data = append(f.data, p...)
	return len(p), nil
}

// Close closes the File, rendering it unusable for I/O. It returns an error, if any.
func (f *File) Close() error {
	return nil
}

// Read a cupe of bytes
func (f *File) Read(p []byte) (int, error) {
	n := copy(p, f.data[f.pointer:])
	f.pointer += n
	if f.pointer == len(f.data) {
		return n, io.EOF
	}
	return n, nil
}

// ResetPointer set file pointer to zero
func (f *File) ResetPointer() {
	f.pointer = 0
}

// Copy file and return new instance
func (f *File) Copy() (*File, error) {
	var datacopy = make([]byte, len(f.data))
	copy(datacopy[:], f.data)
	return &File{
		name:     f.name,
		filemode: f.filemode,
		time:     f.time,
		data:     datacopy,
	}, nil
}
