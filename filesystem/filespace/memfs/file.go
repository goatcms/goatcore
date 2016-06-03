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
