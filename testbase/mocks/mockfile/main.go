package mockfile

import "io"

// File is system file interface
type File interface {
	io.Reader
	io.ReaderAt
	io.Seeker
	io.Closer
}

// MockFile is mock for file type
type MockFile struct {
	data    []byte
	pointer int64
}

// NewMockFile return new mocked file
func NewMockFile(data []byte) File {
	return &MockFile{
		data:    data,
		pointer: 0,
	}
}

// Read read data from file
func (f *MockFile) Read(p []byte) (int, error) {
	if f.pointer >= int64(len(f.data)) {
		return 0, io.EOF
	}
	n := copy(p, f.data[f.pointer:])
	f.pointer += int64(n)
	return n, nil
}

// ReadAt read data at offset
func (f *MockFile) ReadAt(p []byte, off int64) (n int, err error) {
	f.pointer = off
	return f.Read(p)
}

// Seek set file pointer
func (f *MockFile) Seek(offset int64, whence int) (int64, error) {
	switch {
	case whence == 0:
		f.pointer = offset
	case whence == 1:
		f.pointer += offset
	case whence == 2:
		f.pointer = int64(len(f.data)) - offset
	}
	return f.pointer, nil
}

// Close file handle
func (f *MockFile) Close() error {
	f.data = nil
	return nil
}
