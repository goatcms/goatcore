package envs

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/modules/commonm/commservices"
)

// Unit SandboxesUnit is a tool to menage sandboxes.
type Unit struct{}

// UnitFactory create an environment variables unit instance
func UnitFactory(dp app.DependencyProvider) (ins interface{}, error error) {
	return commservices.EnvironmentsUnit(&Unit{}), nil
}

// Envs return scope environments container
func (unit *Unit) Envs(scp app.Scope) (envs commservices.Environments, err error) {
	var (
		envsi  interface{}
		locker = scp.LockData()
	)
	envsi = locker.Value(envKey)
	if envsi == nil {
		envs = NewEnvironments()
		locker.SetValue(envKey, envs)
	} else {
		envs = envsi.(commservices.Environments)
	}
	return envs, locker.Commit()
}
