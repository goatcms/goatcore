package bufferio

import (
	"sync"
)

// Buffer cache input data
type Buffer struct {
	mu    sync.Mutex
	cache string
}

// NewBuffer create new buffer instance
func NewBuffer() *Buffer {
	return &Buffer{}
}

// Write data to buffer
func (buffer *Buffer) Write(s string) {
	buffer.mu.Lock()
	defer buffer.mu.Unlock()
	buffer.cache += s
}

// Read return buffer
func (buffer *Buffer) String() (s string) {
	buffer.mu.Lock()
	defer buffer.mu.Unlock()
	return buffer.cache
}

// ReadAndClean return buffor and clean buffor content
func (buffer *Buffer) ReadAndClean() (s string) {
	buffer.mu.Lock()
	defer buffer.mu.Unlock()
	s = buffer.cache
	buffer.cache = ""
	return s
}
