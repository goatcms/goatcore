package terminalm

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/bootstrap"
	"github.com/goatcms/goatcore/app/mockupapp"
	"github.com/goatcms/goatcore/app/modules/commonm"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

func newApp(options mockupapp.MockupOptions) (mapp *mockupapp.App, bootstraper app.Bootstrap, err error) {
	if mapp, err = mockupapp.NewApp(options); err != nil {
		return nil, nil, err
	}
	bootstraper = bootstrap.NewBootstrap(mapp)
	if err = goaterr.ToError(goaterr.AppendError(nil,
		bootstraper.Register(commonm.NewModule()),
		bootstraper.Register(NewModule()),
	)); err != nil {
		return nil, nil, err
	}
	if err = bootstraper.Init(); err != nil {
		return nil, nil, err
	}
	return mapp, bootstraper, nil
}
