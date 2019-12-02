package modules

import (
	"io"

	"github.com/goatcms/goatcore/app"
)

const (
	// TerminalService is a key for Terminal service
	TerminalService = "TerminalService"
)

// Terminal is global terminal interface
type Terminal interface {
	RunLoop(io app.IOContext) (err error)
	RunString(io app.IOContext, s string) (err error)
	RunCommand(io app.IOContext, args []string) (err error)
	RunCommandFromReader(io app.IOContext, reader io.Reader) (eof bool, err error)
}
