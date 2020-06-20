package gio

import (
	"fmt"
	"io"

	"github.com/goatcms/goatcore/app"
)

const (
	inMode   = 1
	outMode  = 2
	errMode  = 3
	nullMode = 4

	inPrompt  = "\n<<<<<<<<<<<<<<<<<<<<:\n"
	outPrompt = "\n>>>>>>>>>>>>>>>>>>>>:\n"
	errPrompt = "\n!!!!!!!!!!!!!!!!!!!!:\n"
)

// Repeater represent system output
type Repeater struct {
	output      app.Output
	err         app.Output
	input       app.Input
	mode        byte
	errRepeater errRepeater
}

type errRepeater struct {
	repeater *Repeater
}

// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
func (erro errRepeater) Printf(format string, a ...interface{}) (err error) {
	return erro.repeater.PrintErrf(format, a...)
}

// Write data to output
func (erro errRepeater) Write(p []byte) (n int, err error) {
	return erro.repeater.WriteErr(p)
}

// NewRepeater returns a new Repeater.
func NewRepeater(output, err app.Output, input app.Input) *Repeater {
	repeater := &Repeater{
		output: output,
		err:    err,
		input:  input,
		mode:   nullMode,
	}
	repeater.errRepeater = errRepeater{
		repeater: repeater,
	}
	return repeater
}

func (repeater *Repeater) Read(p []byte) (n int, err error) {
	var err2 error
	if n, err = repeater.input.Read(p); err != nil && err != io.EOF {
		return
	}
	if n == 0 {
		return
	}
	if repeater.mode != inMode {
		if _, err2 = repeater.output.Write([]byte(inPrompt)); err2 != nil {
			return 0, err2
		}
		repeater.mode = inMode
	}
	if _, err2 = repeater.output.Write(p[:n]); err2 != nil {
		return 0, err2
	}
	return n, err
}

func (repeater *Repeater) ReadWord() (word string, err error) {
	var err2 error
	if word, err = repeater.input.ReadWord(); err != nil && err != io.EOF {
		return
	}
	if word == "" {
		return
	}
	if repeater.mode != inMode {
		if err2 = repeater.output.Printf(inPrompt); err2 != nil {
			return "", err2
		}
		repeater.mode = inMode
	}
	if err2 = repeater.output.Printf(word); err2 != nil {
		return "", err2
	}
	return word, err
}

func (repeater *Repeater) ReadLine() (line string, err error) {
	var err2 error
	if line, err = repeater.input.ReadLine(); err != nil && err != io.EOF {
		return
	}
	if line == "" {
		return
	}
	if repeater.mode != inMode {
		if err2 = repeater.output.Printf(inPrompt); err2 != nil {
			return "", err2
		}
		repeater.mode = outMode
	}
	if err2 = repeater.output.Printf(line); err2 != nil {
		return "", err2
	}
	return line, nil
}

// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
func (repeater *Repeater) Printf(format string, a ...interface{}) (err error) {
	if repeater.mode != outMode {
		if err = repeater.output.Printf(outPrompt); err != nil {
			return
		}
		repeater.mode = outMode
	}
	return repeater.output.Printf(fmt.Sprintf(format, a...))
}

// Write data to output
func (repeater *Repeater) Write(p []byte) (n int, err error) {
	if repeater.mode != outMode {
		if n, err = repeater.output.Write([]byte(outPrompt)); err != nil {
			return
		}
		repeater.mode = outMode
	}
	return repeater.output.Write(p)
}

// PrintErrf formats according to a format specifier and writes to error output.
// It returns the number of bytes written and any write error encountered.
func (repeater *Repeater) PrintErrf(format string, a ...interface{}) (err error) {
	if repeater.mode != errMode {
		if err = repeater.err.Printf(errPrompt); err != nil {
			return
		}
		repeater.mode = errMode
	}
	return repeater.err.Printf(fmt.Sprintf(format, a...))
}

// WriteErr write data to error output
func (repeater *Repeater) WriteErr(p []byte) (n int, err error) {
	if repeater.mode != errMode {
		if n, err = repeater.err.Write([]byte(errPrompt)); err != nil {
			return
		}
		repeater.mode = errMode
	}
	return repeater.err.Write(p)
}

func (repeater *Repeater) Err() app.Output {
	return repeater.errRepeater
}
