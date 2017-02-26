package wio

import (
	"bytes"
	"errors"
	"testing"
	"time"

	"github.com/goatcms/goatcore/workers"
	"github.com/goatcms/goatcore/workers/jobsync"
)

// An version of bytes.Buffer without ReadFrom and WriteTo
type Buffer struct {
	bytes.Buffer
}

// Simple tests, primarily to verify the ReadFrom and WriteTo callouts inside Copy, CopyBuffer and CopyN.

func TestCopy(t *testing.T) {
	rb := new(Buffer)
	wb := new(Buffer)
	lf := jobsync.NewLifecycle(2*time.Second, true)
	rb.WriteString("hello, world.")
	Copy(wb, rb, lf)
	if wb.String() != "hello, world." {
		t.Errorf("Copy did not work properly")
	}
}

func TestKillCopy(t *testing.T) {
	rb := new(Buffer)
	wb := new(Buffer)
	lf := jobsync.NewLifecycle(2*time.Second, true)
	lf.Kill()
	rb.WriteString("hello, world.")
	_, err := Copy(wb, rb, lf)
	if err != workers.KilledError {
		t.Errorf("Copy did not work properly")
	}
}

func TestCopyBuffer(t *testing.T) {
	rb := new(Buffer)
	wb := new(Buffer)
	lf := jobsync.NewLifecycle(2*time.Second, true)
	rb.WriteString("hello, world.")
	CopyBuffer(wb, rb, make([]byte, 1), lf) // Tiny buffer to keep it honest.
	if wb.String() != "hello, world." {
		t.Errorf("CopyBuffer did not work properly")
	}
}

func TestCopyBufferNil(t *testing.T) {
	rb := new(Buffer)
	wb := new(Buffer)
	lf := jobsync.NewLifecycle(2*time.Second, true)
	rb.WriteString("hello, world.")
	CopyBuffer(wb, rb, nil, lf) // Should allocate a buffer.
	if wb.String() != "hello, world." {
		t.Errorf("CopyBuffer did not work properly")
	}
}

type zeroErrReader struct {
	err error
}

func (r zeroErrReader) Read(p []byte) (int, error) {
	return copy(p, []byte{0}), r.err
}

type errWriter struct {
	err error
}

func (w errWriter) Write([]byte) (int, error) {
	return 0, w.err
}

// In case a Read results in an error with non-zero bytes read, and
// the subsequent Write also results in an error, the error from Write
// is returned, as it is the one that prevented progressing further.
func TestCopyReadErrWriteErr(t *testing.T) {
	lf := jobsync.NewLifecycle(2*time.Second, true)
	er, ew := errors.New("readError"), errors.New("writeError")
	r, w := zeroErrReader{err: er}, errWriter{err: ew}
	n, err := Copy(w, r, lf)
	if n != 0 || err != ew {
		t.Errorf("Copy(zeroErrReader, errWriter) = %d, %v; want 0, writeError", n, err)
	}
}
