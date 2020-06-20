package bootstrap

import (
	"github.com/goatcms/goatcore/app"
)

func testBootstrapImplementInterface(a app.App) app.Bootstrap {
	return NewBootstrap(a)
}
