package commservices

import "github.com/goatcms/goatcore/app"

// EnvironmentsUnit is a unit to manage environment variables
type EnvironmentsUnit interface {
	Envs(scp app.Scope) (envs Environments, err error)
}

// Environments contains and provide environment variables
type Environments interface {
	SetAll(values map[string]string) (err error)
	Set(key, value string) (err error)
	GetAll() map[string]string
	Get(name string) string
}
