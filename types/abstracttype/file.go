package abstracttype

import (
	"github.com/goatcms/goat-core/filesystem"
	"github.com/goatcms/goat-core/types"
)

// File represent dingle file
type File struct {
	fs   filesystem.Filespace
	path string
}

// NewFile create new file record
func NewFile(fs filesystem.Filespace, path string) types.File {
	return &File{
		fs:   fs,
		path: path,
	}
}

// Filespace return filespace of file
func (f *File) Filespace() filesystem.Filespace {
	return f.fs
}

// Path return file path of filespace
func (f *File) Path() string {
	return f.path
}

// Reader crate file reader
func (f *File) Reader() (filesystem.Reader, error) {
	return f.fs.Reader(f.path)
}

// Writer crate file writer
func (f *File) Writer() (filesystem.Writer, error) {
	return f.fs.Writer(f.path)
}
