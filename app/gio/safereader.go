package gio

import (
	"io"
	"sync"
)

// SafeReader wrap io.Reader to single gorutine access
type SafeReader struct {
	mu     sync.Mutex
	reader io.Reader
}

// NewSafeReader create new SafeReader instance
func NewSafeReader(reader io.Reader) *SafeReader {
	return &SafeReader{
		reader: reader,
	}
}

// Read from buffer
func (sr *SafeReader) Read(p []byte) (n int, err error) {
	sr.mu.Lock()
	defer sr.mu.Unlock()
	return sr.reader.Read(p)
}
