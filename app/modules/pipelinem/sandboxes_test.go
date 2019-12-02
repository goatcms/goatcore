package pipelinem

import (
	"testing"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/bootstrap"
	"github.com/goatcms/goatcore/app/mockupapp"
	"github.com/goatcms/goatcore/app/modules/pipelinem/services"
	"github.com/goatcms/goatcore/app/modules/terminal"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

func TestSandboxes(t *testing.T) {
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
		bootstrap.Register(NewModule()),
		bootstrap.Register(terminal.NewModule()),
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
			SandboxsManager services.SandboxsManager `dependency:"SandboxsManager"`
		}
		sandbox services.Sandbox
	)
	if err = mapp.DependencyProvider().InjectTo(&deps); err != nil {
		t.Error(err)
		return
	}
	if sandbox, err = deps.SandboxsManager.Get(""); err != nil {
		t.Error(err)
		return
	}
	if sandbox == nil {
		t.Errorf("Expected default sandbox and take nil")
		return
	}
	if sandbox, err = deps.SandboxsManager.Get("terminal"); err != nil {
		t.Error(err)
		return
	}
	if sandbox == nil {
		t.Errorf("Expected terminal sandbox and take nil")
		return
	}
	if sandbox, err = deps.SandboxsManager.Get("docker:ubuntu:disco"); err != nil {
		t.Error(err)
		return
	}
	if sandbox == nil {
		t.Errorf("Expected terminal sandbox and take nil")
		return
	}
}
