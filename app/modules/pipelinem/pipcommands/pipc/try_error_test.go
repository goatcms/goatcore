package pipc

import (
	"fmt"
	"testing"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/goatapp"
	"github.com/goatcms/goatcore/app/terminal"
)

func TestTryErrorOnSuccess(t *testing.T) {
	t.Parallel()
	var (
		err         error
		mapp        *goatapp.MockupApp
		bootstraper app.Bootstrap
	)
	if mapp, bootstraper, err = newApp(goatapp.Params{
		Arguments: []string{`appname`, `pip:try`, `--name=name`, `--body="emptyCommand"`, `--success=errorCommand`, `--silent=false`},
	}); err != nil {
		t.Error(err)
		return
	}
	term := mapp.Terminal()
	term.SetCommand(terminal.NewCommand(terminal.CommandParams{
		Name: "emptyCommand",
		Callback: func(a app.App, ctx app.IOContext) (err error) {
			return nil
		},
		Help: "Return success",
	}))
	term.SetCommand(terminal.NewCommand(terminal.CommandParams{
		Name: "errorCommand",
		Callback: func(a app.App, ctx app.IOContext) (err error) {
			return fmt.Errorf("Some error")
		},
		Help: "Fail since execution",
	}))
	if err = bootstraper.Run(); err == nil {
		t.Errorf("expected error")
		return
	}
}

func TestTryErrorOnFail(t *testing.T) {
	t.Parallel()
	var (
		err         error
		mapp        *goatapp.MockupApp
		bootstraper app.Bootstrap
	)
	if mapp, bootstraper, err = newApp(goatapp.Params{
		Arguments: []string{`appname`, `pip:try`, `--name=name`, `--body="errorCommand"`, `--fail=errorCommand`, `--silent=false`},
	}); err != nil {
		t.Error(err)
		return
	}
	term := mapp.Terminal()
	term.SetCommand(terminal.NewCommand(terminal.CommandParams{
		Name: "emptyCommand",
		Callback: func(a app.App, ctx app.IOContext) (err error) {
			return nil
		},
		Help: "Return success",
	}))
	term.SetCommand(terminal.NewCommand(terminal.CommandParams{
		Name: "errorCommand",
		Callback: func(a app.App, ctx app.IOContext) (err error) {
			return fmt.Errorf("Some error")
		},
		Help: "Fail since execution",
	}))
	// test
	if err = bootstraper.Run(); err == nil {
		t.Errorf("expected error")
		return
	}
}

func TestTryErrorOnFinally(t *testing.T) {
	t.Parallel()
	var (
		err         error
		mapp        *goatapp.MockupApp
		bootstraper app.Bootstrap
	)
	if mapp, bootstraper, err = newApp(goatapp.Params{
		Arguments: []string{`appname`, `pip:try`, `--name=name`, `--body="emptyCommand"`, `--finally=errorCommand`, `--silent=false`},
	}); err != nil {
		t.Error(err)
		return
	}
	term := mapp.Terminal()
	term.SetCommand(terminal.NewCommand(terminal.CommandParams{
		Name: "emptyCommand",
		Callback: func(a app.App, ctx app.IOContext) (err error) {
			return nil
		},
		Help: "Return success",
	}))
	term.SetCommand(terminal.NewCommand(terminal.CommandParams{
		Name: "errorCommand",
		Callback: func(a app.App, ctx app.IOContext) (err error) {
			return fmt.Errorf("Some error")
		},
		Help: "Fail since execution",
	}))
	// test
	if err = bootstraper.Run(); err == nil {
		t.Errorf("expected error")
		return
	}
}
