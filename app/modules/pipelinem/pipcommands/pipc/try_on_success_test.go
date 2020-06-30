package pipc

import (
	"strings"
	"testing"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/mockupapp"
)

func TestTryOnBodySuccess(t *testing.T) {
	t.Parallel()
	var (
		err         error
		mapp        *mockupapp.App
		bootstraper app.Bootstrap
	)
	if mapp, bootstraper, err = newApp(mockupapp.MockupOptions{
		Input: strings.NewReader(` `),
		Args:  []string{`appname`, `pip:try`, `--name=name`, `--body="bodyCommand"`, `--success=successCommand`, `--fail=failCommand`, `--finally=finallyCommand`, `--silent=false`},
	}); err != nil {
		t.Error(err)
		return
	}
	if err = app.RegisterCommand(mapp, "bodyCommand", func(a app.App, ctx app.IOContext) (err error) {
		ctx.IO().Out().Printf("bodyOutput")
		return nil
	}, "description"); err != nil {
		t.Error(err)
		return
	}
	if err = app.RegisterCommand(mapp, "successCommand", func(a app.App, ctx app.IOContext) (err error) {
		ctx.IO().Out().Printf("successOutput")
		return nil
	}, "description"); err != nil {
		t.Error(err)
		return
	}
	if err = app.RegisterCommand(mapp, "failCommand", func(a app.App, ctx app.IOContext) (err error) {
		ctx.IO().Out().Printf("failOutput")
		return nil
	}, "description"); err != nil {
		t.Error(err)
		return
	}
	if err = app.RegisterCommand(mapp, "finallyCommand", func(a app.App, ctx app.IOContext) (err error) {
		ctx.IO().Out().Printf("finallyOutput")
		return nil
	}, "description"); err != nil {
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
	output := mapp.OutputBuffer().String()
	if !strings.Contains(output, "bodyOutput") {
		t.Errorf("expected 'bodyOutput' and take '%s'", output)
		return
	}
	if !strings.Contains(output, "finallyOutput") {
		t.Errorf("expected 'finallyOutput' and take '%s'", output)
		return
	}
	if !strings.Contains(output, "successOutput") {
		t.Errorf("expected 'successOutput' and take '%s'", output)
		return
	}
	if strings.Contains(output, "failOutput") {
		t.Errorf("unexpected 'failOutput' and take '%s'", output)
		return
	}
}
