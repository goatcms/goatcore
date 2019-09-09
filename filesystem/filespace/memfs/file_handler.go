package memfs

import (
	"io"
	"time"
)

// FileHandler is a handler to single file
type FileHandler struct {
	file    *File
	pointer int
}

// NewFileHandler create new FileHandler instance
func NewFileHandler(file *File) (handler *FileHandler) {
	file.dataMU.Lock()
	return &FileHandler{
		file:    file,
		pointer: 0,
	}
}

// Write write data to stream
func (h *FileHandler) Write(p []byte) (n int, err error) {
	h.file.time = time.Now()
	h.file.data = append(h.file.data, p...)
	return len(p), nil
}

// Read a cupe of bytes from file
func (h *FileHandler) Read(p []byte) (int, error) {
	n := copy(p, h.file.data[h.pointer:])
	h.pointer += n
	if h.pointer == len(h.file.data) {
		return n, io.EOF
	}
	return n, nil
}

// ResetPointer set file pointer to zero
func (h *FileHandler) ResetPointer() {
	h.pointer = 0
}

// Close closes the File, rendering it unusable for I/O. It returns an error, if any.
func (h *FileHandler) Close() error {
	h.file.dataMU.Unlock()
	return nil
}
