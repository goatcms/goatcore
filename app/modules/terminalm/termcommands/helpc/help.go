package helpc

import (
	"github.com/goatcms/goatcore/app"
)

// RunHelp show help
func RunHelp(a app.App, ctx app.IOContext) (err error) {
	var deps struct {
		CommandName string `command:"?$1"`
	}
	if err = ctx.Scope().InjectTo(&deps); err != nil {
		return err
	}
	if deps.CommandName != "" {
		return runCommandHelp(a, ctx, deps.CommandName)
	}
	return runOverview(a, ctx)
}
