package pipc

import (
	"strings"
	"testing"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/goatapp"
)

func TestRunner(t *testing.T) {
	t.Parallel()
	var (
		err         error
		mapp        *goatapp.MockupApp
		bootstraper app.Bootstrap
	)
	if mapp, bootstraper, err = newApp(goatapp.Params{
		Arguments: []string{`appname`, `pip:run`, `--name=name`, `--body="testCommand"`, `--silent=false`},
	}); err != nil {
		t.Error(err)
		return
	}
	if err = bootstraper.Run(); err != nil {
		t.Error(err)
		return
	}
	if err = mapp.Scopes().App().Wait(); err != nil {
		t.Error(err)
		return
	}
	buffer := mapp.OutputBuffer()
	if !strings.Contains(buffer.String(), "output") {
		t.Errorf("expected 'output' and take '%s'", buffer.String())
		return
	}
}
