package smtpmail

import (
	"encoding/base64"
	"io"

	"github.com/goatcms/goatcore/varutil/goaterr"
)

const (
	base64InputGrid  = 3
	base64OutputGrid = 4
)

// Base64Encoder is a bease64 encoder reader
type Base64Encoder struct {
	base     io.Reader
	encoding *base64.Encoding
	buff     []byte
	buffSize int
	eof      error
}

// NewBase64Encoder create new Base64Encoder instance from reader
func NewBase64Encoder(base io.Reader) io.Reader {
	return &Base64Encoder{
		base:     base,
		encoding: base64.StdEncoding,
		buffSize: 0,
	}
}

func (reader *Base64Encoder) Read(dst []byte) (n int, err error) {
	var (
		tmp            []byte
		bytesToProcess int
	)
	if reader.buff == nil {
		reader.buff = make([]byte, len(dst))
	} else if len(dst) < base64OutputGrid {
		return 0, goaterr.Errorf("Output buffer is too small. Minimal buffer size is %v", base64OutputGrid)
	} else if len(dst) > len(reader.buff) {
		oldBuff := reader.buff
		reader.buff = make([]byte, len(dst))
		copy(reader.buff, oldBuff)
	}
	if reader.buffSize == 0 && reader.eof != nil {
		return 0, reader.eof
	}
	if reader.eof == nil && reader.buffSize < len(reader.buff) {
		tmp = reader.buff[reader.buffSize:]
		if n, reader.eof = reader.base.Read(tmp); reader.eof != nil && reader.eof != io.EOF {
			return n, err
		}
		reader.buffSize += n
	}
	if reader.eof != nil {
		bytesToProcess = reader.buffSize
	} else {
		bytesToProcess = reader.buffSize / base64InputGrid * base64InputGrid
	}
	maxSize := len(dst) / base64OutputGrid * base64InputGrid
	if maxSize < bytesToProcess {
		bytesToProcess = maxSize
	}
	reader.encoding.Encode(dst, reader.buff[:bytesToProcess])
	copy(reader.buff, reader.buff[bytesToProcess:])
	reader.buffSize -= bytesToProcess
	return (bytesToProcess + base64InputGrid - 1) / base64InputGrid * base64OutputGrid, nil
}
