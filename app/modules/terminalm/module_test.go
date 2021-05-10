package terminalm

import (
	"strings"
	"testing"

	"github.com/goatcms/goatcore/app/bootstrap"
	"github.com/goatcms/goatcore/app/goatapp"
	"github.com/goatcms/goatcore/app/modules/commonm"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

func TestModule(t *testing.T) {
	var (
		err  error
		mapp *goatapp.MockupApp
	)
	t.Parallel()
	// prepare mockup application
	if mapp, err = goatapp.NewMockupApp(goatapp.Params{
		Name:      "TestApp",
		Arguments: []string{"help"},
	}); err != nil {
		t.Error(err)
		return
	}
	bootstrap := bootstrap.NewBootstrap(mapp)
	if err = goaterr.ToError(goaterr.AppendError(nil,
		bootstrap.Register(NewModule()),
		bootstrap.Register(commonm.NewModule()),
	)); err != nil {
		t.Error(err)
		return
	}
	if err = bootstrap.Init(); err != nil {
		t.Error(err)
		return
	}
	if err = bootstrap.Run(); err != nil {
		t.Error(err)
		return
	}
	output := mapp.OutputBuffer().String()
	if !strings.Contains(output, "TestApp") {
		t.Errorf("Expected correct application name (displayed by help function) and take: %s", output)
	}
}
