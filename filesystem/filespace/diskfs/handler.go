package diskfs

import (
	"os"
)

// FileHandler is a handle to a open file
type FileHandler struct {
	*os.File
}

// NewFilespace create new Filespace instance
func NewFileHandler(os *os.File) *FileHandler {
	return &FileHandler{
		File: os,
	}
}

// Close handler
func (handler *FileHandler) Close() (err error) {
	if err = handler.File.Sync(); err != nil {
		return
	}
	return handler.File.Close()
}
