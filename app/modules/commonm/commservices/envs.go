package commservices

import "github.com/goatcms/goatcore/app"

// SSHCert represent ssh certificate
type SSHCert struct {
	Secret string
	Public string
}

// EnvironmentsUnit is a unit to manage environment variables
type EnvironmentsUnit interface {
	Envs(scp app.Scope) (envs Environments, err error)
}

// Environments contains and provide environment variables
type Environments interface {
	Get(name string) string
	All() map[string]string
	SetAll(values map[string]string) (err error)
	Set(key, value string) (err error)
	SSHCert() SSHCert
	SetSSHCert(cert SSHCert)
}
