package envs

import (
	"sync"

	"github.com/goatcms/goatcore/app/modules/commonm/commservices"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

// Environments SandboxesEnvironments is a tool to menage sandboxes.
type Environments struct {
	data map[string]string
	mu   sync.RWMutex
	cert commservices.SSHCert
}

// NewEnvironments return new empty Environments instance
func NewEnvironments() commservices.Environments {
	return &Environments{
		data: make(map[string]string),
	}
}

// SetAll valid and add all map values to sandboxes environments.
// Env key must cntains letters and underscores (first char must be letter).
func (envs *Environments) SetAll(values map[string]string) (err error) {
	if err = envs.valid(values); err != nil {
		return err
	}
	envs.mu.Lock()
	defer envs.mu.Unlock()
	for key, value := range values {
		envs.data[key] = value
	}
	return nil
}

// Set valid and add a environment to sandboxes environments.
// Env key must cntains letters and underscores (first char must be letter).
func (envs *Environments) Set(key, value string) (err error) {
	if err = envs.validKey(key); err != nil {
		return err
	}
	envs.mu.Lock()
	defer envs.mu.Unlock()
	envs.data[key] = value
	return nil
}

// valid return error if map is incorrect
func (envs *Environments) valid(values map[string]string) (err error) {
	for key := range values {
		if err = envs.validKey(key); err != nil {
			return err
		}
	}
	return nil
}

func (envs *Environments) validKey(key string) (err error) {
	if !envNamePattern.MatchString(key) {
		return goaterr.Errorf("Sandbox environment name is incorrect: %s", key)
	}
	return nil
}

// All return all sandboxes environments
func (envs *Environments) All() (result map[string]string) {
	result = make(map[string]string)
	envs.mu.RLock()
	defer envs.mu.RUnlock()
	for key, value := range envs.data {
		result[key] = value
	}
	return result
}

// Get return environment variable by name
func (envs *Environments) Get(name string) string {
	envs.mu.RLock()
	defer envs.mu.RUnlock()
	return envs.data[name]
}

// SSHCert return ssh certificate
func (envs *Environments) SSHCert() commservices.SSHCert {
	envs.mu.Lock()
	defer envs.mu.Unlock()
	return envs.cert
}

// SetSSHCert set new ssh certificate
func (envs *Environments) SetSSHCert(cert commservices.SSHCert) {
	envs.mu.RLock()
	defer envs.mu.RUnlock()
	envs.cert = cert
}
