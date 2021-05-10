package ocm

import (
	"testing"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/bootstrap"
	"github.com/goatcms/goatcore/app/goatapp"
	"github.com/goatcms/goatcore/app/modules/ocm/ocservices"
	"github.com/goatcms/goatcore/app/modules/terminalm"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

func TestModule(t *testing.T) {
	var (
		err  error
		mapp app.App
	)
	t.Parallel()
	// prepare mockup application
	if mapp, err = goatapp.NewMockupApp(goatapp.Params{}); err != nil {
		t.Error(err)
		return
	}
	bootstrap := bootstrap.NewBootstrap(mapp)
	if err = goaterr.ToError(goaterr.AppendError(nil,
		bootstrap.Register(terminalm.NewModule()),
		bootstrap.Register(NewModule()),
	)); err != nil {
		t.Error(err)
		return
	}
	if err := bootstrap.Init(); err != nil {
		t.Error(err)
		return
	}
	// test
	var deps struct {
		Manager ocservices.Manager `dependency:"OCManager"`
	}
	if err = mapp.DependencyProvider().InjectTo(&deps); err != nil {
		t.Error(err)
		return
	}
}
