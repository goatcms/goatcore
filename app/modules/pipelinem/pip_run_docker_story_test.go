package pipelinem

import (
	"strings"
	"testing"
	"time"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/mockupapp"
	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/filesystem/filespace/diskfs"
	"github.com/goatcms/goatcore/testbase"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

func TestPipRunDockerStory(t *testing.T) {
	t.Parallel()
	var (
		err         error
		mapp        *mockupapp.App
		bootstraper app.Bootstrap
		cwd         filesystem.Filespace
	)
	if _, err = testbase.LoadDockerTestConfig(); err != nil {
		t.Skip(err.Error())
		return
	}
	if cwd, err = diskfs.NewFilespace("./"); err != nil {
		t.Error(err)
		return
	}
	if mapp, bootstraper, err = newApp(mockupapp.MockupOptions{
		RootFilespace: cwd,
		Input: strings.NewReader(`
			pip:run --name=first --sandbox="docker:alpine" --body=<<FIRSTEND
echo "outputAla"
FIRSTEND --silent=false
			pip:run --name=second --wait=first --body="echoMa" --silent=false
			pip:run --name=last --wait=second --body="echoKota" --silent=false
			pip:wait
			`),
		Args: []string{`appname`, `terminal`},
	}); err != nil {
		t.Error(err)
		return
	}
	if err = goaterr.ToError(goaterr.AppendError(nil, app.RegisterCommand(mapp, "echoMa", func(a app.App, ctx app.IOContext) (err error) {
		time.Sleep(20 * time.Millisecond)
		return ctx.IO().Out().Printf("outputMa")
	}, ""), app.RegisterCommand(mapp, "echoKota", func(a app.App, ctx app.IOContext) (err error) {
		return ctx.IO().Out().Printf("outputKota")
	}, ""))); err != nil {
		t.Error(err)
		return
	}
	// test
	if err = bootstraper.Run(); err != nil {
		t.Error(err)
		return
	}
	if err = mapp.AppScope().Wait(); err != nil {
		t.Error(err)
		return
	}
	buffer := mapp.OutputBuffer()
	result := buffer.String()
	posAla := strings.Index(result, "outputAla")
	posMa := strings.Index(result, "outputMa")
	posKota := strings.Index(result, "outputKota")
	if posAla == -1 && posMa == -1 || posKota == -1 {
		t.Errorf("expected outputcontains 'outputAla' 'outputMa' and  'outputKota' and take '%s'", result)
		return
	}
	if posAla > posMa || posMa > posKota {
		t.Errorf("order incorrect for result: '%s'", result)
		return
	}
}
