package terminalm

import (
	"fmt"
	"strings"
	"testing"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/mockupapp"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

func TestTerminalPipRunChildScopeKillStory(t *testing.T) {
	t.Parallel()
	var (
		err         error
		mapp        *mockupapp.App
		bootstraper app.Bootstrap
	)
	// it check async execution
	if mapp, bootstraper, err = newApp(mockupapp.MockupOptions{
		Input: strings.NewReader(`kill`),
		Args:  []string{`appname`, `terminal`, `--strict=true`},
	}); err != nil {
		t.Error(err)
		return
	}
	if err = goaterr.ToError(goaterr.AppendError(nil, app.RegisterCommand(mapp, "kill", func(a app.App, ctx app.IOContext) (err error) {
		ctx.Scope().Kill()
		return fmt.Errorf("some error")
	}, ""))); err != nil {
		t.Error(err)
		return
	}
	// test
	if err = bootstraper.Run(); err == nil {
		t.Errorf("Expected error")
		return
	}
}
