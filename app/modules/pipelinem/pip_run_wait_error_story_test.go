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

func TestPipRunWaitForErrorStory(t *testing.T) {
	t.Parallel()
	var (
		err         error
		mapp        *goatapp.MockupApp
		bootstraper app.Bootstrap
	)
	if mapp, bootstraper, err = newApp(goatapp.Params{
		IO: goatapp.IO{
			In: gio.NewAppInput(strings.NewReader(`
			pip:run --name=first --body="error"
			pip:run --name=second --wait=first --body="afterKillCommand"
			`)),
		},
		Arguments: []string{`appname`, `terminal`, `--strict=true`, `--silent=false`},
	}); err != nil {
		t.Error(err)
		return
	}
	mapp.Terminal().SetCommand(
		terminal.NewCommand(terminal.CommandParams{
			Name: "error",
			Callback: func(a app.App, ctx app.IOContext) (err error) {
				time.Sleep(10 * time.Millisecond)
				ctx.Scope().AppendError(fmt.Errorf("some error"))
				ctx.IO().Out().Printf("error")
				return nil
			},
		}),
		terminal.NewCommand(terminal.CommandParams{
			Name: "afterKillCommand",
			Callback: func(a app.App, ctx app.IOContext) (err error) {
				return ctx.IO().Out().Printf("Unexpected output")
			},
		}),
	)
	// test
	if err = bootstraper.Run(); err == nil {
		t.Errorf("Expected error when bootstraper.Run(). Output: %s", mapp.OutputBuffer().String())
		return
	}
	if err = mapp.Scopes().App().Wait(); err == nil {
		t.Errorf("Expected error when Wait(). Output: %s", mapp.OutputBuffer().String())
		return
	}
	result := mapp.OutputBuffer().String()
	if strings.Contains(result, "Unexpected output") {
		t.Errorf("afterKillCommand should be skipped. Result: %s", result)
		return
	}
}
