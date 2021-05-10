package pipc

import (
	"strings"
	"testing"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/goatapp"
	"github.com/goatcms/goatcore/app/terminal"
)

func TestTryOnBodySuccess(t *testing.T) {
	t.Parallel()
	var (
		err         error
		mapp        *goatapp.MockupApp
		bootstraper app.Bootstrap
	)
	if mapp, bootstraper, err = newApp(goatapp.Params{
		Arguments: []string{`appname`, `pip:try`, `--name=name`, `--body="bodyCommand"`, `--success=successCommand`, `--fail=failCommand`, `--finally=finallyCommand`, `--silent=false`},
	}); err != nil {
		t.Error(err)
		return
	}
	term := mapp.Terminal()
	term.SetCommand(terminal.NewCommand(terminal.CommandParams{
		Name: "bodyCommand",
		Callback: func(a app.App, ctx app.IOContext) (err error) {
			ctx.IO().Out().Printf("bodyOutput")
			return nil
		},
	}))
	term.SetCommand(terminal.NewCommand(terminal.CommandParams{
		Name: "successCommand",
		Callback: func(a app.App, ctx app.IOContext) (err error) {
			ctx.IO().Out().Printf("successOutput")
			return nil
		},
	}))
	term.SetCommand(terminal.NewCommand(terminal.CommandParams{
		Name: "failCommand",
		Callback: func(a app.App, ctx app.IOContext) (err error) {
			ctx.IO().Out().Printf("failOutput")
			return nil
		},
	}))
	term.SetCommand(terminal.NewCommand(terminal.CommandParams{
		Name: "finallyCommand",
		Callback: func(a app.App, ctx app.IOContext) (err error) {
			ctx.IO().Out().Printf("finallyOutput")
			return nil
		},
	}))
	// test
	if err = bootstraper.Run(); err != nil {
		t.Error(err)
		return
	}
	if err = mapp.Scopes().App().Wait(); err != nil {
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
