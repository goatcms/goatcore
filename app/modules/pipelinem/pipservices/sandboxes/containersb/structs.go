package containersb

import (
	"github.com/goatcms/goatcore/app/modules/commonm/commservices"
	"github.com/goatcms/goatcore/app/modules/ocm/ocservices"
)

// deps contains package dependencies
type deps struct {
	EnvironmentsUnit commservices.EnvironmentsUnit
	OCManager        ocservices.Manager
}
