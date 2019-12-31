package pipelinem

import (
	"strings"
	"testing"
	"time"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/mockupapp"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

func TestNestedPipWaitStory(t *testing.T) {
	t.Parallel()
	var (
		err         error
		mapp        *mockupapp.App
		bootstraper app.Bootstrap
	)
	if mapp, bootstraper, err = newApp(mockupapp.MockupOptions{
		Input: strings.NewReader(`
pip:run --name=first --body=<<EOF
				echoOne
				pip:run --name=firstnested --body="echoTwo"
				pip:run --name=secondnested --wait=firstnested --body="echoThree"
EOF
			pip:run --name=second --wait=first --body="echoFour"
			pip:run --name=last --wait=second --body="echoFive"
			`),
		Args: []string{`appname`, `terminal`},
	}); err != nil {
		t.Error(err)
		return
	}
	if err = goaterr.ToErrors(goaterr.AppendError(nil, app.RegisterCommand(mapp, "echoOne", func(a app.App, ctx app.IOContext) (err error) {
		time.Sleep(10 * time.Millisecond)
		return ctx.IO().Out().Printf("1")
	}, ""), app.RegisterCommand(mapp, "echoTwo", func(a app.App, ctx app.IOContext) (err error) {
		time.Sleep(20 * time.Millisecond)
		return ctx.IO().Out().Printf("2")
	}, ""), app.RegisterCommand(mapp, "echoThree", func(a app.App, ctx app.IOContext) (err error) {
		return ctx.IO().Out().Printf("3")
	}, ""), app.RegisterCommand(mapp, "echoFour", func(a app.App, ctx app.IOContext) (err error) {
		return ctx.IO().Out().Printf("4")
	}, ""), app.RegisterCommand(mapp, "echoFive", func(a app.App, ctx app.IOContext) (err error) {
		return ctx.IO().Out().Printf("5")
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
