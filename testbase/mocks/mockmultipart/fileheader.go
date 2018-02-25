package mockmultipart

import (
	"mime/multipart"

	"github.com/goatcms/goatcore/testbase/mocks/mockfile"
)

// FileHeader is mock represent FileHeader
type FileHeader struct {
	data []byte
}

// NewFileHeader create FileHeader instance
func NewFileHeader(data []byte) *FileHeader {
	return &FileHeader{
		data: data,
	}
}

// Open create new file handler
func (fh *FileHeader) Open() (multipart.File, error) {
	file := mockfile.NewMockFile(fh.data)
	return file, nil
}
