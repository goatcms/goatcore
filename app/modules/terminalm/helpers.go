package terminalm

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/gio"
	"github.com/goatcms/goatcore/app/injector"
	"github.com/goatcms/goatcore/app/scope"
	"github.com/goatcms/goatcore/app/scope/argscope"
)

func fixSpace(s string, l int) string {
	ll := l - len(s)
	prefix := make([]rune, ll)
	for i := range prefix {
		prefix[i] = ' '
	}
	return string(prefix) + s
}

func newCommandContext(ctx app.IOContext, args []string) (commCtx app.IOContext, err error) {
	argsData := &scope.DataScope{
		Data: make(map[string]interface{}),
	}
	if err = argscope.InjectArgsToScope(args, argsData); err != nil {
		return nil, err
	}
	ctxScope := ctx.Scope()
	commandScope := &scope.Scope{
		EventScope: ctxScope,
		DataScope:  ctxScope,
		SyncScope:  ctxScope,
		Injector: injector.NewMultiInjector([]app.Injector{
			ctxScope,
			argsData.Injector("command"),
		}),
	}
	return gio.NewIOContext(commandScope, ctx.IO()), nil
}
