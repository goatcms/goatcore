package sshsb

import (
	"strings"

	"github.com/goatcms/goatcore/app/modules/commonm/commservices"
	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices"
	"github.com/goatcms/goatcore/dependency"
)

// SSHSandboxBuilder create new docker sandbox builder
type SSHSandboxBuilder struct {
	deps deps
}

// NewSSHSandboxBuilder create SSHSandboxBuilder
func NewSSHSandboxBuilder(envUnit commservices.EnvironmentsUnit) *SSHSandboxBuilder {
	return &SSHSandboxBuilder{
		deps: deps{
			EnvironmentsUnit: envUnit,
		},
	}
}

// SSHSandboxBuilderFactory create SSHSandboxBuilder
func SSHSandboxBuilderFactory(dp dependency.Provider) (ins interface{}, err error) {
	builder := &SSHSandboxBuilder{}
	if err = dp.InjectTo(&builder.deps); err != nil {
		return nil, err
	}
	return pipservices.SandboxBuilder(builder), nil
}

// Is return true if name is match to terminal factory
func (factory *SSHSandboxBuilder) Is(name string) bool {
	return strings.HasPrefix(name, "ssh:")
}

// Build return terminal sandbox
func (factory *SSHSandboxBuilder) Build(name string) (sandbox pipservices.Sandbox, err error) {
	return NewSSHSandbox(name[len("ssh:"):], factory.deps)
}
