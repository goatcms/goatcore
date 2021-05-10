package pipelinem

import (
	"strings"
	"testing"
	"time"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/gio"
	"github.com/goatcms/goatcore/app/goatapp"
	"github.com/goatcms/goatcore/app/terminal"
)

func TestPipRunWaitStory(t *testing.T) {
	t.Parallel()
	var (
		err         error
		mapp        *goatapp.MockupApp
		bootstraper app.Bootstrap
	)
	if mapp, bootstraper, err = newApp(goatapp.Params{
		IO: goatapp.IO{
			In: gio.NewAppInput(strings.NewReader(`
			pip:run --name=first --body="echoAla" --silent=false
			pip:run --name=second --wait=first --body="echoMa" --silent=false
			pip:run --name=last --wait=second --body="echoKota" --silent=false
			`)),
		},
		Arguments: []string{`appname`, `terminal`, ``, `--body="testCommand"`},
	}); err != nil {
		t.Error(err)
		return
	}
	mapp.Terminal().SetCommand(
		terminal.NewCommand(terminal.CommandParams{
			Name: "echoAla",
			Callback: func(a app.App, ctx app.IOContext) (err error) {
				time.Sleep(10 * time.Millisecond)
				return ctx.IO().Out().Printf("outputAla")
			},
		}),
		terminal.NewCommand(terminal.CommandParams{
			Name: "echoMa",
			Callback: func(a app.App, ctx app.IOContext) (err error) {
				time.Sleep(20 * time.Millisecond)
				return ctx.IO().Out().Printf("outputMa")
			},
		}),
		terminal.NewCommand(terminal.CommandParams{
			Name: "echoKota",
			Callback: func(a app.App, ctx app.IOContext) (err error) {
				return ctx.IO().Out().Printf("outputKota")
			},
		}),
	)
	// test
	if err = bootstraper.Run(); err != nil {
		t.Error(err)
		return
	}
	if err = mapp.Scopes().App().Wait(); err != nil {
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
