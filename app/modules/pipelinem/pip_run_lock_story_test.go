package pipelinem

import (
	"strings"
	"testing"
	"time"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/mockupapp"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

func TestPipRunLockStory(t *testing.T) {
	t.Parallel()
	var (
		err         error
		mapp        *mockupapp.App
		bootstraper app.Bootstrap
	)
	if mapp, bootstraper, err = newApp(mockupapp.MockupOptions{
		Input: strings.NewReader(`
			pip:run --name=first --rlock=resource --body="long"
			pip:run --name=second --wlock=resource --body="short"
			`),
		Args: []string{`appname`, `terminal`, ``, `--body="testCommand"`},
	}); err != nil {
		t.Error(err)
		return
	}
	if err = goaterr.ToErrors(goaterr.AppendError(nil, app.RegisterCommand(mapp, "long", func(a app.App, ctx app.IOContext) (err error) {
		ctx.IO().Out().Printf("lock")
		time.Sleep(30 * time.Millisecond)
		return ctx.IO().Out().Printf("unlock")
	}, ""), app.RegisterCommand(mapp, "short", func(a app.App, ctx app.IOContext) (err error) {
		time.Sleep(10 * time.Millisecond)
		return ctx.IO().Out().Printf("write")
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
	posLock := strings.Index(result, "lock")
	posUnlock := strings.Index(result, "unlock")
	posWrite := strings.Index(result, "write")
	if posLock == -1 && posUnlock == -1 || posWrite == -1 {
		t.Errorf("expected output contains 'lock' 'unlock' and 'write' and take '%s'", result)
		return
	}
	if posWrite > posLock && posWrite < posUnlock {
		t.Errorf("other pipe con not write to lock resource")
		return
	}
}
