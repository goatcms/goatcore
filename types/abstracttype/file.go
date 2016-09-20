package abstracttype

import "github.com/goatcms/goat-core/filesystem"

// File represent dingle file
type File struct {
	fs   filesystem.Filespace
	path string
}

// NewFile create new file record
func NewFile(fs filesystem.Filespace, path string) *File {
	return &File{
		fs:   fs,
		path: path,
	}
}

// Filespace return filespace of file
func (f *File) Filespace() filesystem.Filespace {
	return f.Filespace()
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
