package pipelinem

import (
	"testing"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/bootstrap"
	"github.com/goatcms/goatcore/app/goatapp"
	"github.com/goatcms/goatcore/app/modules/commonm"
	"github.com/goatcms/goatcore/app/modules/ocm"
	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices"
	"github.com/goatcms/goatcore/app/modules/terminalm"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

func TestSandboxes(t *testing.T) {
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
		bootstrap.Register(NewModule()),
		bootstrap.Register(terminalm.NewModule()),
		bootstrap.Register(commonm.NewModule()),
		bootstrap.Register(ocm.NewModule()),
	)); err != nil {
		t.Error(err)
		return
	}
	if err = bootstrap.Init(); err != nil {
		t.Error(err)
		return
	}
	// test
	var (
		deps struct {
			SandboxesManager pipservices.SandboxesManager `dependency:"PipSandboxesManager"`
		}
		sandbox pipservices.Sandbox
	)
	if err = mapp.DependencyProvider().InjectTo(&deps); err != nil {
		t.Error(err)
		return
	}
	if sandbox, err = deps.SandboxesManager.Get(""); err != nil {
		t.Error(err)
		return
	}
	if sandbox == nil {
		t.Errorf("Expected default sandbox and take nil")
		return
	}
	if sandbox, err = deps.SandboxesManager.Get("self"); err != nil {
		t.Error(err)
		return
	}
	if sandbox == nil {
		t.Errorf("Expected terminal sandbox and take nil")
		return
	}
	if sandbox, err = deps.SandboxesManager.Get("container:ubuntu:disco"); err != nil {
		t.Error(err)
		return
	}
	if sandbox == nil {
		t.Errorf("Expected docker sandbox and take nil")
		return
	}
}
