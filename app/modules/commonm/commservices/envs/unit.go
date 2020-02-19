package envs

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/modules/commonm/commservices"
	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

// Unit SandboxesUnit is a tool to menage sandboxes.
type Unit struct{}

// UnitFactory create an environment variables unit instance
func UnitFactory(dp dependency.Provider) (ins interface{}, error error) {
	return commservices.EnvironmentsUnit(&Unit{}), nil
}

// Envs return scope environments container
func (unit *Unit) Envs(scp app.Scope) (envs commservices.Environments, err error) {
	var (
		envsi  interface{}
		locker = scp.LockData()
	)
	if envsi, err = locker.Get(envKey); err != nil {
		return nil, goaterr.ToError(goaterr.AppendError(nil, err, locker.Commit()))
	}
	if envsi == nil {
		envs = NewEnvironments()
		if err = locker.Set(envKey, envs); err != nil {
			return nil, goaterr.ToError(goaterr.AppendError(nil, err, locker.Commit()))
		}
	} else {
		envs = envsi.(commservices.Environments)
	}
	return envs, locker.Commit()
}
