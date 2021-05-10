package infos

import (
	"github.com/goatcms/goatcore/app/modules/commonm/commservices"
	"github.com/goatcms/goatcore/dependency"
)

type Params struct {
	Authors     []commservices.InfoPerson
	Description string
	License     commservices.InfoLicense
	PoweredBy   string
}

// Info SandboxesInfo is a tool to menage sandboxes.
type Info struct {
	params Params
}

func NewInfo(params Params) commservices.Info {
	return &Info{
		params: params,
	}
}

// InfoFactory create an environment variables info instance
func InfoFactory(dp dependency.Provider) (ins interface{}, error error) {
	return commservices.Info(&Info{
		params: Params{
			License:   NewLicense(LicenseParams{}),
			PoweredBy: commservices.DefaultPoweredBy,
		},
	}), nil
}

// Authors return application authors list
func (info *Info) Authors() []commservices.InfoPerson {
	return info.params.Authors
}

// Description return application description
func (info *Info) Description() string {
	return info.params.Description
}

// License return application license
func (info *Info) License() commservices.InfoLicense {
	return info.params.License
}

// License return application license
func (info *Info) PoweredBy() string {
	return info.params.PoweredBy
}
