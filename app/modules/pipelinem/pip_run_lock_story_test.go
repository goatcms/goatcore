package pipelinem

import (
	"strings"
	"testing"
	"time"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/gio"
	"github.com/goatcms/goatcore/app/goatapp"
	"github.com/goatcms/goatcore/app/terminal"
)

func TestPipRunLockStory(t *testing.T) {
	t.Parallel()
	var (
		err         error
		mapp        *goatapp.MockupApp
		bootstraper app.Bootstrap
	)
	if mapp, bootstraper, err = newApp(goatapp.Params{
		IO: goatapp.IO{
			In: gio.NewAppInput(strings.NewReader(`
			pip:run --name=first --rlock=resource --body="long" --silent=false
			pip:run --name=second --wlock=resource --body="short" --silent=false
			`)),
		},
		Arguments: []string{`appname`, `terminal`},
	}); err != nil {
		t.Error(err)
		return
	}
	mapp.Terminal().SetCommand(
		terminal.NewCommand(terminal.CommandParams{
			Name: "long",
			Callback: func(a app.App, ctx app.IOContext) (err error) {
				ctx.IO().Out().Printf("lock")
				time.Sleep(30 * time.Millisecond)
				return ctx.IO().Out().Printf("unlock")
			},
		}),
		terminal.NewCommand(terminal.CommandParams{
			Name: "short",
			Callback: func(a app.App, ctx app.IOContext) (err error) {
				time.Sleep(10 * time.Millisecond)
				return ctx.IO().Out().Printf("write")
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
	buffer := mapp.OutputBuffer()
	result := buffer.String()
	posLock := strings.Index(result, "lock")
	posUnlock := strings.Index(result, "unlock")
	posWrite := strings.Index(result, "write")
	if posLock == -1 && posUnlock == -1 || posWrite == -1 {
		t.Errorf("expected output contains 'lock' 'unlock' and 'write' and take '%s'", result)
		return
	}
	if posWrite > posLock && posWrite < posUnlock {
		t.Errorf("other pipe con not write to lock resource")
		return
	}
}
