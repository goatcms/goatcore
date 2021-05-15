package termc

import (
	"github.com/goatcms/goatcore/app"
)

// RunExit close terminal
func RunExit(a app.App, ctx app.IOContext) (err error) {
	a.Scopes().App().Stop()
	return nil
}
