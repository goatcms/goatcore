package pipelinem

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/mockupapp"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

func TestPipRunChildScopeKillStory(t *testing.T) {
	t.Parallel()
	var (
		err         error
		mapp        *mockupapp.App
		bootstraper app.Bootstrap
	)
	// it check async execution
	if mapp, bootstraper, err = newApp(mockupapp.MockupOptions{
		Input: strings.NewReader(`
			pip:run --name=first --body="killStatus" --silent=false
			pip:run --name=second --body="kill" --silent=false
			`),
		Args: []string{`appname`, `terminal`},
	}); err != nil {
		t.Error(err)
		return
	}
	if err = goaterr.ToErrors(goaterr.AppendError(nil, app.RegisterCommand(mapp, "killStatus", func(a app.App, ctx app.IOContext) (err error) {
		time.Sleep(10 * time.Millisecond)
		if ctx.Scope().IsKilled() {
			return ctx.IO().Out().Printf("is_killed")
		}
		return ctx.IO().Out().Printf("is_not_killed")
	}, ""), app.RegisterCommand(mapp, "kill", func(a app.App, ctx app.IOContext) (err error) {
		ctx.Scope().Kill()
		return fmt.Errorf("some error")
	}, ""))); err != nil {
		t.Error(err)
		return
	}
	// test
	bootstraper.Run()
	if err = mapp.AppScope().Wait(); err != nil {
		t.Error(err)
		return
	}
	result := mapp.OutputBuffer().String()
	if !strings.Contains(result, "is_killed") {
		t.Errorf("expected 'is_killed' result and take '%s'", result)
		return
	}
}
