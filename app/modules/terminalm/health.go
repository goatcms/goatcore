package terminalm

import (
	"sort"
	"strings"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/varutil"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

// HealthComamnd run health command. It show application helthy message.
func HealthComamnd(a app.App, ctx app.IOContext) (err error) {
	var (
		deps struct {
			CommandScope app.Scope `dependency:"CommandScope"`
		}
		keys     []string
		first    = true
		errs     []error
		msg      string
		cb       app.HealthCheckerCallback
		io       = ctx.IO()
		ctxScope = ctx.Scope()
		ins      interface{}
	)
	if err = a.DependencyProvider().InjectTo(&deps); err != nil {
		return err
	}
	keys = varutil.ToStringArr(deps.CommandScope.Keys())
	sort.Strings(keys)
	for _, key := range keys {
		if !strings.HasPrefix(key, healthPrefix) {
			continue
		}
		if first {
			io.Out().Printf("\nHealth:\n")
			first = false
		}
		ins = deps.CommandScope.Value(key)
		cb = ins.(app.HealthCheckerCallback)
		if msg, err = cb(a, ctxScope); err != nil {
			errs = append(errs, err)
			io.Out().Printf("[FAIL]  %s\n", msg)
		} else {
			io.Out().Printf("[OK]    %s\n", msg)
		}
	}
	io.Out().Printf("\n")
	return goaterr.ToError(errs)
}
