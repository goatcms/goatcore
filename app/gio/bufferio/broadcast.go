package bufferio

import (
	"fmt"
	"io"

	"github.com/goatcms/goatcore/varutil/goaterr"
)

// Broadcast is helper to brodcast call to many writers
type Broadcast struct {
	buffer  *Buffer
	writers []io.Writer
}

// NewBroadcast return new Broadcast instance
func NewBroadcast(buffer *Buffer, writers []io.Writer) *Broadcast {
	if buffer == nil {
		buffer = NewBuffer()
	}
	return &Broadcast{
		buffer:  buffer,
		writers: append([]io.Writer{buffer}, writers...),
	}
}

// Writer is the interface that wraps the basic Write method.
func (broadcast *Broadcast) Write(p []byte) (n int, err error) {
	for _, out := range broadcast.writers {
		if n, err = out.Write(p); err != nil {
			return n, err
		}
		if n != len(p) {
			return n, goaterr.Errorf("Can not write %d bytes (%d bytes writen)", len(p), n)
		}
	}
	return n, err
}

// Printf print to multiple outputs.
func (broadcast *Broadcast) Printf(format string, a ...interface{}) (err error) {
	_, err = broadcast.Write([]byte(fmt.Sprintf(format, a...)))
	return
}

// Add writer (write buffored data to new outputs buffored )
func (broadcast *Broadcast) Add(writer io.Writer) (err error) {
	if _, err = writer.Write(broadcast.buffer.Bytes()); err != nil {
		return err
	}
	broadcast.writers = append(broadcast.writers, writer)
	return
}

// String return writed content
func (broadcast *Broadcast) String() string {
	return broadcast.buffer.String()
}

// Bytes return writed content
func (broadcast *Broadcast) Bytes() []byte {
	return broadcast.buffer.Bytes()
}
