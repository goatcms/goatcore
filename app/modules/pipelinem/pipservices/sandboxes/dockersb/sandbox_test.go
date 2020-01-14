package dockersb

import (
	"strings"
	"testing"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/gio"
	"github.com/goatcms/goatcore/app/gio/bufferio"
	"github.com/goatcms/goatcore/app/modules/commonm/commservices"
	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices"
	"github.com/goatcms/goatcore/app/scope"
	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/filesystem/filespace/diskfs"
	"github.com/goatcms/goatcore/testbase"
)

func TestDockerSandbox(t *testing.T) {
	t.Parallel()
	var (
		err  error
		mapp app.App
		scp  = scope.NewScope(scope.Params{})
		cwd  filesystem.Filespace
	)
	if _, err = testbase.LoadDockerTestConfig(); err != nil {
		t.Skip(err.Error())
		return
	}
	if mapp, err = newApp(); err != nil {
		t.Error(err)
		return
	}
	// test
	var deps struct {
		Runner pipservices.Runner `dependency:"PipRunner"`
	}
	if err = mapp.DependencyProvider().InjectTo(&deps); err != nil {
		t.Error(err)
		return
	}
	buffer := bufferio.NewBuffer()
	if cwd, err = diskfs.NewFilespace("./"); err != nil {
		t.Error(err)
		return
	}
	if err = deps.Runner.Run(pipservices.Pip{
		Context: pipservices.PipContext{
			In:    gio.NewInput(strings.NewReader(`echo "output"`)),
			Out:   bufferio.NewBufferOutput(buffer),
			Err:   bufferio.NewBufferOutput(buffer),
			Scope: scp,
			CWD:   cwd,
		},
		Name:    "name",
		Sandbox: "docker:alpine",
		Lock:    commservices.LockMap{},
		Wait:    []string{},
	}); err != nil {
		t.Error(err)
		return
	}
	if err = scp.Wait(); err != nil {
		t.Error(err)
		return
	}
	if !strings.Contains(buffer.String(), "output") {
		t.Errorf("expected 'output'")
		return
	}
}
