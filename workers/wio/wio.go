package wio

import (
	"io"

	"github.com/goatcms/goatcore/workers"
	"github.com/goatcms/goatcore/workers/jobsync"
)

// Copy copies from src to dst until either EOF is reached
// on src or an error occurs. It returns the number of bytes
// copied and the first error encountered while copying, if any.
//
// A successful Copy returns err == nil, not err == EOF.
// Because Copy is defined to read from src until EOF, it does
// not treat an EOF from Read as an error to be reported.
//
// If src implements the WriterTo interface,
// the copy is implemented by calling src.WriteTo(dst).
// Otherwise, if dst implements the ReaderFrom interface,
// the copy is implemented by calling dst.ReadFrom(src).
func Copy(w io.Writer, r io.Reader, lf *jobsync.Lifecycle) (written int64, err error) {
	return copyBuffer(w, r, nil, lf)
}

// CopyBuffer is identical to Copy except that it stages through the
// provided buffer (if one is required) rather than allocating a
// temporary one. If buf is nil, one is allocated; otherwise if it has
// zero length, CopyBuffer panics.
func CopyBuffer(dst io.Writer, src io.Reader, buf []byte, lf *jobsync.Lifecycle) (written int64, err error) {
	if buf != nil && len(buf) == 0 {
		panic("empty buffer in io.CopyBuffer")
	}
	if lf == nil {
		panic("empty lifecycle")
	}
	return copyBuffer(dst, src, buf, lf)
}

// copyBuffer is the actual implementation of Copy and CopyBuffer.
// if buf is nil, one is allocated.
func copyBuffer(dst io.Writer, src io.Reader, buf []byte, lf *jobsync.Lifecycle) (written int64, err error) {
	if buf == nil {
		buf = make([]byte, 32*1024)
	}
	for {
		if lf.IsKilled() {
			return written, workers.KilledError
		}
		nr, er := src.Read(buf)
		if nr > 0 {
			nw, ew := dst.Write(buf[0:nr])
			if nw > 0 {
				written += int64(nw)
			}
			if ew != nil {
				err = ew
				break
			}
			if nr != nw {
				err = io.ErrShortWrite
				break
			}
		}
		if er != nil {
			if er != io.EOF {
				err = er
			}
			break
		}
	}
	return written, err
}
