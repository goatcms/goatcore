package pipelinem

import (
	"testing"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/bootstrap"
	"github.com/goatcms/goatcore/app/mockupapp"
	"github.com/goatcms/goatcore/app/modules/commonm"
	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices"
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
		bootstrap.Register(terminalm.NewModule()),
		bootstrap.Register(commonm.NewModule()),
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
		SandboxesManager pipservices.SandboxesManager `dependency:"PipSandboxesManager"`
		NamespacesUnit   pipservices.NamespacesUnit   `dependency:"PipNamespacesUnit"`
		Runner           pipservices.Runner           `dependency:"PipRunner"`
		TasksUnit        pipservices.TasksUnit        `dependency:"PipTasksUnit"`
	}
	if err = mapp.DependencyProvider().InjectTo(&deps); err != nil {
		t.Error(err)
		return
	}
}
