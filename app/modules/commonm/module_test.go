package commonm

import (
	"testing"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/bootstrap"
	"github.com/goatcms/goatcore/app/mockupapp"
	"github.com/goatcms/goatcore/app/modules/commonm/commservices"
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
	if mapp, err = mockupapp.NewApp(mockupapp.MockupOptions{}); err != nil {
		t.Error(err)
		return
	}
	bootstrap := bootstrap.NewBootstrap(mapp)
	if err = goaterr.ToErrors(goaterr.AppendError(nil,
		bootstrap.Register(terminal.NewModule()),
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
		SharedMutex commservices.SharedMutex `dependency:"CommonSharedMutex"`
		WaitManager commservices.WaitManager `dependency:"CommonWaitManager"`
	}
	if err = mapp.DependencyProvider().InjectTo(&deps); err != nil {
		t.Error(err)
		return
	}
}
