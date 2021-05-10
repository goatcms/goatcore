package pipelinem

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/gio"
	"github.com/goatcms/goatcore/app/goatapp"
	"github.com/goatcms/goatcore/app/terminal"
)

func TestPipRunChildScopeKillStory(t *testing.T) {
	t.Parallel()
	var (
		err         error
		mapp        *goatapp.MockupApp
		bootstraper app.Bootstrap
	)
	// it check async execution
	if mapp, bootstraper, err = newApp(goatapp.Params{
		IO: goatapp.IO{
			In: gio.NewAppInput(strings.NewReader(`
			pip:run --name=first --body="kill" --silent=false
			pip:run --name=second --wait=first --body="killStatus" --silent=false
			`)),
		},
		Arguments: []string{`appname`, `terminal`},
	}); err != nil {
		t.Error(err)
		return
	}
	mapp.Terminal().SetCommand(
		terminal.NewCommand(terminal.CommandParams{
			Name: "killStatus",
			Callback: func(a app.App, ctx app.IOContext) (err error) {
				time.Sleep(10 * time.Millisecond)
				// it will never executed because return error by command stop pipeline
				if ctx.Scope().IsDone() {
					return ctx.IO().Out().Printf("is_killed")
				}
				return ctx.IO().Out().Printf("is_not_killed")
			},
		}),
		terminal.NewCommand(terminal.CommandParams{
			Name: "kill",
			Callback: func(a app.App, ctx app.IOContext) (err error) {
				return fmt.Errorf("some error")
			},
		}),
	)
	// test
	if err = bootstraper.Run(); err != nil {
		t.Error(err)
		return
	}
	if err = mapp.Scopes().App().Wait(); err != nil {
		t.Error(err)
		return
	}
	result := mapp.OutputBuffer().String()
	if strings.Contains(result, "is_killed") || strings.Contains(result, "is_not_killed") {
		t.Errorf("expected stopped pipeline before killStatus command and take '%s'", result)
		return
	}
}
