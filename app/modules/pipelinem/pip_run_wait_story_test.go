package pipelinem

import (
	"strings"
	"testing"
	"time"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/mockupapp"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

func TestPipRunWaitStory(t *testing.T) {
	t.Parallel()
	var (
		err         error
		mapp        *mockupapp.App
		bootstraper app.Bootstrap
	)
	if mapp, bootstraper, err = newApp(mockupapp.MockupOptions{
		Input: strings.NewReader(`
			pip:run --name=first --body="echoAla"
			pip:run --name=second --wait=first --body="echoMa"
			pip:run --name=last --wait=second --body="echoKota"
			`),
		Args: []string{`appname`, `terminal`, ``, `--body="testCommand"`},
	}); err != nil {
		t.Error(err)
		return
	}
	if err = goaterr.ToErrors(goaterr.AppendError(nil, app.RegisterCommand(mapp, "echoAla", func(a app.App, ctx app.IOContext) (err error) {
		time.Sleep(10 * time.Millisecond)
		return ctx.IO().Out().Printf("outputAla")
	}, ""), app.RegisterCommand(mapp, "echoMa", func(a app.App, ctx app.IOContext) (err error) {
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
