package terminal

import (
	"github.com/goatcms/goatcore/app"
)

// HelpComamnd show help message
func HelpComamnd(a app.App, ctxScope app.Scope) error {
	return PrintHelp(a)
}
