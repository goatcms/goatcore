package terminalm

import (
	"strings"
	"testing"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/bootstrap"
	"github.com/goatcms/goatcore/app/mockupapp"
	"github.com/goatcms/goatcore/app/modules"
)

func TestRunLoop(t *testing.T) {
	var (
		err  error
		mapp *mockupapp.App
	)
	t.Parallel()
	// prepare mockup application
	if mapp, err = mockupapp.NewApp(mockupapp.MockupOptions{
		Input: strings.NewReader("help"),
	}); err != nil {
		t.Error(err)
		return
	}
	bootstrap := bootstrap.NewBootstrap(mapp)
	if err = bootstrap.Register(NewModule()); err != nil {
		t.Error(err)
		return
	}
	if err = bootstrap.Init(); err != nil {
		t.Error(err)
		return
	}
	// test
	var deps struct {
		Terminal modules.Terminal `dependency:"TerminalService"`
		AppScope app.Scope        `dependency:"AppScope"`
	}
	if err = mapp.DependencyProvider().InjectTo(&deps); err != nil {
		t.Error(err)
		return
	}
	if err = deps.Terminal.RunLoop(mapp.IOContext()); err != nil {
		t.Error(err)
		return
	}
	out := mapp.OutputBuffer().String()
	if !strings.Contains(out, "Commands:") {
		t.Errorf("Expectend command list in result and take: %s", out)
		return
	}
}
