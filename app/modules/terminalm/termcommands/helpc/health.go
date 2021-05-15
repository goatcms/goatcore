package helpc

import (
	"sort"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/modules/terminalm/termcommands"
	"github.com/goatcms/goatcore/app/terminal/termformatter"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

// RunHealth run health command. It show application helthy status.
func RunHealth(a app.App, ctx app.IOContext) (err error) {
	var (
		errs      []error
		io        = ctx.IO()
		msg       string
		names     []string
		lineWidth = termcommands.LineWidth
	)
	names = a.HealthCheckerNames()
	sort.Strings(names)
	if len(names) > 0 {
		io.Out().Printf("Health:\n")
		formatter := termformatter.NewBlockFormatter(io.Out(), lineWidth,
			termformatter.NewBlockDef(8, termformatter.ToRight, termformatter.ToRight),
			termformatter.NewBlockDef(2, termformatter.ToRight, termformatter.ToRight),
			termformatter.NewBlockDef(lineWidth-10, termformatter.ToLeft, termformatter.ToLeft),
		)
		for _, name := range names {
			checker := a.HealthChecker(name)
			if msg, err = checker(a, ctx.Scope()); err != nil {
				errs = append(errs, err)
				formatter.PrintBlocks("[FAIL]", "", msg)
			} else {
				formatter.PrintBlocks("[OK]", "", msg)
			}
		}
	}
	io.Out().Printf("\n")
	return goaterr.ToError(errs)
}
