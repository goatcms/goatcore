package termservices

import (
	"io"

	"github.com/goatcms/goatcore/app"
)

type Terminal interface {
	RunLoop(io app.IOContext, prompt string) (err error)
	RunString(io app.IOContext, s string) (err error)
	RunCommand(io app.IOContext, args []string) (err error)
	RunCommandFromReader(io app.IOContext, reader io.Reader) (eof bool, err error)
	Extends(exts ...app.TerminalCommand) Terminal
}
