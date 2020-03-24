package aesgcm256cfs

import (
	"io"
	"io/ioutil"

	"github.com/goatcms/goatcore/filesystem"
)

type reader struct {
	data   []byte
	stream io.ReadCloser
}

func newReader(key []byte, stream io.ReadCloser) (_out filesystem.Reader, err error) {
	var (
		buf []byte
	)
	if buf, err = ioutil.ReadAll(stream); err != nil {
		return nil, err
	}
	if err = stream.Close(); err != nil {
		return nil, err
	}
	if buf, err = NewCipher().Decrypt(key, buf); err != nil {
		return nil, err
	}
	return &reader{
		data: buf,
	}, nil
}

func (r *reader) Read(p []byte) (n int, err error) {
	n = copy(p, r.data)
	r.data = r.data[n:]
	if len(r.data) == 0 {
		return n, io.EOF
	}
	return n, nil
}

func (r *reader) Close() (err error) {
	r.data = nil
	return nil
}
