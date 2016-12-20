package memfs

import (
	"os"
	"time"
)

type File struct {
	name     string
	filemode os.FileMode
	time     time.Time
	data     []byte
	pointer  int
}

func (f *File) Name() string {
	return f.name
}

func (f *File) Mode() os.FileMode {
	return f.filemode
}

func (f *File) ModTime() time.Time {
	return f.time
}

func (f *File) Sys() interface{} {
	return nil
}

func (f *File) Size() int64 {
	return int64(len(f.data))
}

func (f *File) IsDir() bool {
	return false
}

func (f *File) GetData() []byte {
	return f.data
}

func (f *File) SetData(data []byte) {
	f.data = data
}

func (f *File) Write(p []byte) (n int, err error) {
	f.data = append(f.data, p...)
	return len(p), nil
}

func (f *File) Close() error {
	return nil
}

func (f *File) Read(p []byte) (int, error) {
	n := copy(p, f.data[f.pointer:])
	f.pointer += n
	return n, nil
}

func (f *File) ResetPointer() {
	f.pointer = 0
}

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
