package termexec

import (
	"strings"
	"testing"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/gio"
	"github.com/goatcms/goatcore/app/goatapp"
	"github.com/goatcms/goatcore/app/terminal"
)

func TestRunLoop(t *testing.T) {
	var (
		err  error
		mapp *goatapp.MockupApp
	)
	t.Parallel()
	// prepare mockup application
	if mapp, err = goatapp.NewMockupApp(goatapp.Params{
		IO: goatapp.IO{
			In: gio.NewAppInput(strings.NewReader("commandName")),
		},
	}); err != nil {
		t.Error(err)
		return
	}
	commands := terminal.NewCommands(
		terminal.NewCommand(terminal.CommandParams{
			Name: "commandName",
			Help: "",
			Callback: func(a app.App, ctx app.IOContext) (err error) {
				ctx.IO().Out().Printf("expected resutl")
				return nil
			},
			Arguments: nil,
		}),
	)
	rctx := NewRunCtx(RunCtxParams{
		Application: mapp,
		Ctx:         mapp.IOContext(),
		Commands:    commands,
	})
	if err = RunLoop(rctx, "\n>"); err != nil {
		t.Error(err)
		return
	}
	out := mapp.OutputBuffer().String()
	if !strings.Contains(out, "expected resutl") {
		t.Errorf("expected 'expected resutl' and: %s", out)
		return
	}
}
