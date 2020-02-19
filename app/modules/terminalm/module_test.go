package terminalm

import (
	"testing"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/bootstrap"
	"github.com/goatcms/goatcore/app/mockupapp"
	"github.com/goatcms/goatcore/app/modules"
)

func TestModule(t *testing.T) {
	var (
		err  error
		mapp app.App
	)
	t.Parallel()
	// prepare mockup application
	if mapp, err = mockupapp.NewApp(mockupapp.MockupOptions{}); err != nil {
		t.Error(err)
		return
	}
	bootstrap := bootstrap.NewBootstrap(mapp)
	if err = bootstrap.Register(NewModule()); err != nil {
		t.Error(err)
		return
	}
	if err := bootstrap.Init(); err != nil {
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
	if err = deps.Terminal.RunString(mapp.IOContext(), "help"); err != nil {
		t.Error(err)
		return
	}
}
