package bufferio

import (
	"bytes"
	"fmt"
	"io"
	"sync"

	"github.com/goatcms/goatcore/varutil/goaterr"
)

// Buffer cache input data
type Buffer struct {
	mu   sync.Mutex
	data bytes.Buffer
}

// NewBuffer create new buffer instance
func NewBuffer() *Buffer {
	return &Buffer{}
}

func newReaderBuffer() io.Reader {
	return &Buffer{}
}

func newWriteBuffer() io.Writer {
	return &Buffer{}
}

// WriteString write string to buffer
func (buffer *Buffer) WriteString(s string) (err error) {
	var (
		n    int
		data = []byte(s)
	)
	buffer.mu.Lock()
	defer buffer.mu.Unlock()
	if n, err = buffer.data.Write(data); err != nil {
		return err
	}
	if n != len(data) {
		return goaterr.Errorf("expected write %d bytes and writed %d bytes", len(data), n)
	}
	return nil
}

// Read from buffer
func (buffer *Buffer) Read(p []byte) (n int, err error) {
	buffer.mu.Lock()
	defer buffer.mu.Unlock()
	return buffer.data.Read(p)
}

// WriteString write string to buffer
func (buffer *Buffer) Write(p []byte) (n int, err error) {
	buffer.mu.Lock()
	defer buffer.mu.Unlock()
	return buffer.data.Write(p)
}

// Printf to buffer
func (buffer *Buffer) Printf(format string, a ...interface{}) (err error) {
	_, err = buffer.Write([]byte(fmt.Sprintf(format, a...)))
	return
}

// Read return buffer
func (buffer *Buffer) String() string {
	buffer.mu.Lock()
	defer buffer.mu.Unlock()
	return buffer.data.String()
}

// Bytes return buffer
func (buffer *Buffer) Bytes() []byte {
	buffer.mu.Lock()
	defer buffer.mu.Unlock()
	return buffer.data.Bytes()
}

// ReadAndClean return buffor and clean buffor content
func (buffer *Buffer) ReadAndClean() (s string) {
	buffer.mu.Lock()
	defer buffer.mu.Unlock()
	s = buffer.data.String()
	buffer.data.Reset()
	return s
}
