package sshsb

import (
	"github.com/goatcms/goatcore/app/modules/commonm/commservices"
)

// deps contains package dependencies
type deps struct {
	EnvironmentsUnit commservices.EnvironmentsUnit `dependency:"CommonEnvironmentsUnit"`
}
