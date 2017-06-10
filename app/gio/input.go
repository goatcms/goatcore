package gio

import (
	"errors"
	"io"
	"sync"

	"github.com/goatcms/goatcore/app"
)

var (
	errNegativeRead = errors.New("gio: reader returned negative count from Read")
	errBufforFull   = errors.New("gio: reader buffer is full")
)

const (
	defaultBufSize           = 4096
	minReadBufferSize        = 16
	maxConsecutiveEmptyReads = 100
)

// Input represent system input
type Input struct {
	buf  []byte
	rd   io.Reader // reader provided by the client
	r, w int       // buf read and write positions
	mu   sync.Mutex
	eof  bool
}

// NewInputSize returns a new Input whose buffer has at least the specified
func NewInputSize(rd io.Reader, size int) *Input {
	if size < minReadBufferSize {
		size = minReadBufferSize
	}
	in := &Input{
		buf: make([]byte, size),
		rd:  rd,
		r:   0,
		w:   0,
	}
	in.fill()
	return in
}

// NewInput returns a new Input whose buffer has the default size.
func NewInput(rd io.Reader) *Input {
	return NewInputSize(rd, defaultBufSize)
}

// NewAppInput returns a new app.Input whose buffer has the default size.
func NewAppInput(rd io.Reader) app.Input {
	return app.Input(NewInputSize(rd, defaultBufSize))
}

func (in *Input) fill() error {
	// Slide existing data to beginning.
	if in.r > 0 {
		copy(in.buf, in.buf[in.r:in.w])
		in.w -= in.r
		in.r = 0
	}
	if in.w >= len(in.buf) {
		panic("gio.Input.fill: tried to fill full buffer")
	}
	// Read new data: try a limited number of times.
	for i := maxConsecutiveEmptyReads; i > 0; i-- {
		n, err := in.rd.Read(in.buf[in.w:])
		if n < 0 {
			return errNegativeRead
		}
		in.w += n
		if err != nil {
			return err
		}
		if n > 0 {
			return nil
		}
	}
	return io.ErrNoProgress
}

// ReadWord return next word from input stream
func (in *Input) ReadWord() (s string, err error) {
	in.mu.Lock()
	defer in.mu.Unlock()
	if in.eof {
		return "", io.EOF
	}
LoopSkipWhiteAtBeginOfString:
	for isWhiteChar(in.buf[in.r]) {
		for ; in.r < in.w; in.r++ {
			if !isWhiteChar(in.buf[in.r]) {
				break LoopSkipWhiteAtBeginOfString
			}
		}
		if err = in.fill(); err != nil {
			if err == io.EOF {
				in.buf = nil
				in.rd = nil
				in.eof = true
			}
			return "", err
		}
	}
	offset := 0
	filled := false
	pos := in.r + offset
	for {
		for pos < in.w {
			if isWhiteChar(in.buf[pos]) {
				s := string(in.buf[in.r:pos])
				in.r = pos
				return s, nil
			}
			offset++
			pos = in.r + offset
		}
		if filled {
			if in.eof {
				s := string(in.buf[in.r:pos])
				in.r = pos
				in.buf = nil
				in.rd = nil
				return s, io.EOF
			}
			return "", errBufforFull
		}
		filled = true
		if err = in.fill(); err != nil {
			if err == io.EOF {
				in.eof = true
			} else {
				return "", err
			}
		}
		pos = in.r + offset
	}
}

// ReadLine return next line from input stream
func (in *Input) ReadLine() (s string, err error) {
	in.mu.Lock()
	defer in.mu.Unlock()
	if in.eof {
		return "", io.EOF
	}
LoopSkipWhiteAtBeginOfString:
	for isWhiteChar(in.buf[in.r]) {
		for ; in.r < in.w; in.r++ {
			if !isWhiteChar(in.buf[in.r]) {
				break LoopSkipWhiteAtBeginOfString
			}
		}
		if err = in.fill(); err != nil {
			return "", err
		}
	}
	offset := 0
	filled := false
	pos := in.r + offset
	for {
		for pos < in.w {
			if isNewLine(in.buf[pos]) {
				s := string(in.buf[in.r:pos])
				in.r = pos
				return s, nil
			}
			offset++
			pos = in.r + offset
		}
		if filled {
			if in.eof {
				s := string(in.buf[in.r:pos])
				in.r = pos
				in.buf = nil
				in.rd = nil
				return s, io.EOF
			}
			return "", errBufforFull
		}
		filled = true
		if err = in.fill(); err != nil {
			if err == io.EOF {
				in.eof = true
			} else {
				return "", err
			}
		}
		pos = in.r + offset
	}
}
