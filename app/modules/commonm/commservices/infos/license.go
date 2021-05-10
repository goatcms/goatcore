package infos

import (
	"time"

	"github.com/goatcms/goatcore/app/modules/commonm/commservices"
)

type LicenseParams struct {
	Company string
	End     time.Time
	Start   time.Time
	Type    string
	URL     string
}

// License SandboxesLicense is a tool to menage sandboxes.
type License struct {
	params LicenseParams
}

func NewLicense(params LicenseParams) commservices.InfoLicense {
	return &License{
		params: params,
	}
}

// Company return licence company
func (license *License) Company() string {
	return license.params.Company
}

// End return expired license time
func (license *License) End() time.Time {
	return license.params.End
}

// Start return begin license time
func (license *License) Start() time.Time {
	return license.params.End
}

// Type return begin license time
func (license *License) Type() string {
	return license.params.Type
}

// URL return url to license agreement
func (license *License) URL() string {
	return license.params.URL
}
