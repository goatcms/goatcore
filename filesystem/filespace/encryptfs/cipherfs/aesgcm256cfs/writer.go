package aesgcm256cfs

import (
	"io"

	"github.com/goatcms/goatcore/filesystem"
)

type writer struct {
	data   []byte
	key    []byte
	stream io.WriteCloser
}

func newWriter(key []byte, stream io.WriteCloser) filesystem.Writer {
	return &writer{
		key:    key,
		stream: stream,
	}
}

func (w *writer) Write(p []byte) (n int, err error) {
	w.data = append(w.data, p...)
	return len(p), nil
}

func (w *writer) Close() (err error) {
	var data []byte
	if data, err = NewCipher().Encrypt(w.key, w.data); err != nil {
		return err
	}
	w.data = nil
	if _, err = w.stream.Write(data); err != nil {
		return err
	}
	return w.stream.Close()
}
