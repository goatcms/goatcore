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
			pip:run --name=first --body="kill" --silent=false
			pip:run --name=second --wait=first --body="killStatus" --silent=false
			`),
		Args: []string{`appname`, `terminal`},
	}); err != nil {
		t.Error(err)
		return
	}
	if err = goaterr.ToError(goaterr.AppendError(nil, app.RegisterCommand(mapp, "killStatus", func(a app.App, ctx app.IOContext) (err error) {
		time.Sleep(10 * time.Millisecond)
		// it will never executed because return error by command stop pipeline
		if ctx.Scope().IsKilled() {
			return ctx.IO().Out().Printf("is_killed")
		}
		return ctx.IO().Out().Printf("is_not_killed")
	}, ""), app.RegisterCommand(mapp, "kill", func(a app.App, ctx app.IOContext) (err error) {
		// we kill pipeline by return "some error"
		return fmt.Errorf("some error")
	}, ""))); err != nil {
		t.Error(err)
		return
	}
	// test
	if err = bootstraper.Run(); err != nil {
		t.Error(err)
		return
	}
	if err = mapp.AppScope().Wait(); err != nil {
		t.Error(err)
		return
	}
	result := mapp.OutputBuffer().String()
	if strings.Contains(result, "is_killed") || strings.Contains(result, "is_not_killed") {
		t.Errorf("expected stopped pipeline before killStatus command and take '%s'", result)
		return
	}
}
