package sshsb

import (
	"strings"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/varutil/goaterr"

	"github.com/goatcms/goatcore/app/modules/commonm/commservices"
	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices"
)

// SSHSandboxBuilder create new ssh sandbox builder
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
func SSHSandboxBuilderFactory(dp app.DependencyProvider) (ins interface{}, err error) {
	builder := &SSHSandboxBuilder{}
	if err = dp.InjectTo(&builder.deps); err != nil {
		return nil, err
	}
	return pipservices.SandboxBuilder(builder), nil
}

// Is return true if name is match to terminal factory
func (factory *SSHSandboxBuilder) Is(name string) bool {
	return strings.HasPrefix(name, "ssh:") || strings.HasPrefix(name, "sshb:")
}

// Build return terminal sandbox
func (factory *SSHSandboxBuilder) Build(name string) (sandbox pipservices.Sandbox, err error) {
	// sh for ssh:*
	if strings.HasPrefix(name, "ssh:") {
		return NewSSHSandbox(name[len("ssh:"):], "sh", factory.deps)
	}
	// bash for sshd:*
	if strings.HasPrefix(name, "sshb:") {
		return NewSSHSandbox(name[len("sshb:"):], "bash", factory.deps)
	}
	return nil, goaterr.Errorf("Incorrect sandbox %s", name)
}
