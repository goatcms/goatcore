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

func TestPipRunWaitForErrorStory(t *testing.T) {
	t.Parallel()
	var (
		err         error
		mapp        *mockupapp.App
		bootstraper app.Bootstrap
	)
	if mapp, bootstraper, err = newApp(mockupapp.MockupOptions{
		Input: strings.NewReader(`
			pip:run --name=first --body="error"
			pip:run --name=second --wait=first --body="afterKillCommand"
			`),
		Args: []string{`appname`, `terminal`, `--strict=true`},
	}); err != nil {
		t.Error(err)
		return
	}
	if err = goaterr.ToErrors(goaterr.AppendError(nil, app.RegisterCommand(mapp, "error", func(a app.App, ctx app.IOContext) (err error) {
		time.Sleep(10 * time.Millisecond)
		ctx.Scope().AppendError(fmt.Errorf("some error"))
		ctx.IO().Out().Printf("error")
		return nil
	}, ""), app.RegisterCommand(mapp, "afterKillCommand", func(a app.App, ctx app.IOContext) (err error) {
		return ctx.IO().Out().Printf("Unexpected output")
	}, ""))); err != nil {
		t.Error(err)
		return
	}
	// test
	if err = bootstraper.Run(); err == nil {
		t.Errorf("Expected error when bootstraper.Run(). Output: %s", mapp.OutputBuffer().String())
		return
	}
	if err = mapp.AppScope().Wait(); err == nil {
		t.Errorf("Expected error when Wait(). Output: %s", mapp.OutputBuffer().String())
		return
	}
	result := mapp.OutputBuffer().String()
	if strings.Contains(result, "Unexpected output") {
		t.Errorf("afterKillCommand should be skipped. Result: %s", result)
		return
	}
}
