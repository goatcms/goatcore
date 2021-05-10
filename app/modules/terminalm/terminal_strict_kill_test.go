package terminalm

import (
	"fmt"
	"strings"
	"testing"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/gio"
	"github.com/goatcms/goatcore/app/goatapp"
	"github.com/goatcms/goatcore/app/terminal"
)

func TestTerminalPipRunChildScopeKillStory(t *testing.T) {
	t.Parallel()
	var (
		err         error
		mapp        *goatapp.MockupApp
		bootstraper app.Bootstrap
	)
	// it check async execution
	if mapp, bootstraper, err = newApp(goatapp.Params{
		IO: goatapp.IO{
			In: gio.NewAppInput(strings.NewReader(`kill`)),
		},
		Arguments: []string{`appname`, `terminal`, `--strict=true`},
	}); err != nil {
		t.Error(err)
		return
	}
	mapp.Terminal().SetCommand(
		terminal.NewCommand(terminal.CommandParams{
			Name: "kill",
			Callback: func(a app.App, ctx app.IOContext) (err error) {
				ctx.Scope().Kill()
				return fmt.Errorf("some error")
			},
		}),
	)
	// test
	if err = bootstraper.Run(); err == nil {
		t.Errorf("Expected error")
		return
	}
}
