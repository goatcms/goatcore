package app

import (
	"io"

	"github.com/goatcms/goatcore/filesystem"
)

// Input represent a standard input
type Input interface {
	io.Reader
	ReadWord() (string, error)
	ReadLine() (string, error)
}

// Output represent a standard output
type Output interface {
	io.Writer
	Printf(format string, a ...interface{}) error
}

// IO represent a standard input/output
type IO interface {
	In() Input
	Out() Output
	Err() Output
	CWD() filesystem.Filespace
}

// IOContext represent application execution context
type IOContext interface {
	IO() IO
	Scope() Scope
	Close() (err error)
}

// Broadcast buffer and write data to each added writer data
type Broadcast interface {
	Output
	Add(writer io.Writer) (err error)
}

// BufferedBroadcast buffer and write data to each added writer data
type BufferedBroadcast interface {
	Broadcast
	String() string
	Bytes() []byte
}
