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

func TestNestedPipWaitStory(t *testing.T) {
	t.Parallel()
	var (
		err         error
		mapp        *goatapp.MockupApp
		bootstraper app.Bootstrap
	)
	if mapp, bootstraper, err = newApp(goatapp.Params{
		IO: goatapp.IO{
			In: gio.NewAppInput(strings.NewReader(`
pip:run --name=first --silent=false --body=<<EOF
				echoOne
				pip:run --name=firstnested --body="echoTwo" --silent=false
				pip:run --name=secondnested --wait=firstnested --body="echoThree" --silent=false
EOF
			pip:run --name=second --wait=first --body="echoFour" --silent=false
			pip:run --name=last --wait=second --body="echoFive" --silent=false
			`)),
		},
		Arguments: []string{`appname`, `terminal`},
	}); err != nil {
		t.Error(err)
		return
	}
	mapp.Terminal().SetCommand(
		terminal.NewCommand(terminal.CommandParams{
			Name: "echoOne",
			Callback: func(a app.App, ctx app.IOContext) (err error) {
				time.Sleep(10 * time.Millisecond)
				return ctx.IO().Out().Printf("1")
			},
		}),
		terminal.NewCommand(terminal.CommandParams{
			Name: "echoTwo",
			Callback: func(a app.App, ctx app.IOContext) (err error) {
				time.Sleep(20 * time.Millisecond)
				return ctx.IO().Out().Printf("2")
			},
		}),
		terminal.NewCommand(terminal.CommandParams{
			Name: "echoThree",
			Callback: func(a app.App, ctx app.IOContext) (err error) {
				return ctx.IO().Out().Printf("3")
			},
		}),
		terminal.NewCommand(terminal.CommandParams{
			Name: "echoFour",
			Callback: func(a app.App, ctx app.IOContext) (err error) {
				return ctx.IO().Out().Printf("4")
			},
		}),
		terminal.NewCommand(terminal.CommandParams{
			Name: "echoFive",
			Callback: func(a app.App, ctx app.IOContext) (err error) {
				return ctx.IO().Out().Printf("5")
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
	onePos := strings.Index(result, "1")
	twoPos := strings.Index(result, "2")
	threePos := strings.Index(result, "3")
	fourPos := strings.Index(result, "4")
	fivePos := strings.Index(result, "5")
	if onePos == -1 && twoPos == -1 || threePos == -1 || fourPos == -1 || fivePos == -1 {
		t.Errorf("expected all numbers (1,2,3,4,5): %s", result)
		return
	}
	if onePos > twoPos || twoPos > threePos || threePos > fourPos || fourPos > fivePos {
		t.Errorf("numbers order is incorrect for result: '%s'", result)
		return
	}
}
