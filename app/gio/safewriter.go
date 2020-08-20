package gio

import (
	"io"
	"sync"
)

// SafeWriter wrap io.Writer to multigorutines safe version
type SafeWriter struct {
	mu     sync.Mutex
	writer io.Writer
}

// NewSafeWriter create new buffer instance
func NewSafeWriter(writer io.Writer) *SafeWriter {
	return &SafeWriter{
		writer: writer,
	}
}

// Write write data
func (sw *SafeWriter) Write(p []byte) (n int, err error) {
	sw.mu.Lock()
	defer sw.mu.Unlock()
	return sw.writer.Write(p)
}
