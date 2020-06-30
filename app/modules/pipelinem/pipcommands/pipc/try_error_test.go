package pipc

import (
	"fmt"
	"strings"
	"testing"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/mockupapp"
)

func TestTryErrorOnSuccess(t *testing.T) {
	t.Parallel()
	var (
		err         error
		mapp        *mockupapp.App
		bootstraper app.Bootstrap
	)
	if mapp, bootstraper, err = newApp(mockupapp.MockupOptions{
		Input: strings.NewReader(` `),
		Args:  []string{`appname`, `pip:try`, `--name=name`, `--body="emptyCommand"`, `--success=errorCommand`, `--silent=false`},
	}); err != nil {
		t.Error(err)
		return
	}
	if err = app.RegisterCommand(mapp, "emptyCommand", func(a app.App, ctx app.IOContext) (err error) {
		return nil
	}, "description"); err != nil {
		t.Error(err)
		return
	}
	if err = app.RegisterCommand(mapp, "errorCommand", func(a app.App, ctx app.IOContext) (err error) {
		return fmt.Errorf("Some error")
	}, "description"); err != nil {
		t.Error(err)
		return
	}
	// test
	if err = bootstraper.Run(); err == nil {
		t.Errorf("expected error")
		return
	}
}

func TestTryErrorOnFail(t *testing.T) {
	t.Parallel()
	var (
		err         error
		mapp        *mockupapp.App
		bootstraper app.Bootstrap
	)
	if mapp, bootstraper, err = newApp(mockupapp.MockupOptions{
		Input: strings.NewReader(` `),
		Args:  []string{`appname`, `pip:try`, `--name=name`, `--body="errorCommand"`, `--fail=errorCommand`, `--silent=false`},
	}); err != nil {
		t.Error(err)
		return
	}
	if err = app.RegisterCommand(mapp, "emptyCommand", func(a app.App, ctx app.IOContext) (err error) {
		return nil
	}, "description"); err != nil {
		t.Error(err)
		return
	}
	if err = app.RegisterCommand(mapp, "errorCommand", func(a app.App, ctx app.IOContext) (err error) {
		return fmt.Errorf("Some error")
	}, "description"); err != nil {
		t.Error(err)
		return
	}
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
		mapp        *mockupapp.App
		bootstraper app.Bootstrap
	)
	if mapp, bootstraper, err = newApp(mockupapp.MockupOptions{
		Input: strings.NewReader(` `),
		Args:  []string{`appname`, `pip:try`, `--name=name`, `--body="emptyCommand"`, `--finally=errorCommand`, `--silent=false`},
	}); err != nil {
		t.Error(err)
		return
	}
	if err = app.RegisterCommand(mapp, "emptyCommand", func(a app.App, ctx app.IOContext) (err error) {
		return nil
	}, "description"); err != nil {
		t.Error(err)
		return
	}
	if err = app.RegisterCommand(mapp, "errorCommand", func(a app.App, ctx app.IOContext) (err error) {
		return fmt.Errorf("Some error")
	}, "description"); err != nil {
		t.Error(err)
		return
	}
	// test
	if err = bootstraper.Run(); err == nil {
		t.Errorf("expected error")
		return
	}
}
