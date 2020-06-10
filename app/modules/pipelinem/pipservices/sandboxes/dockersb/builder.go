package dockersb

import (
	"strings"

	"github.com/goatcms/goatcore/app/modules/commonm/commservices"
	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices"
	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

// DockerSandboxBuilder create new docker sandbox builder
type DockerSandboxBuilder struct {
	deps deps
}

// NewDockerSandboxBuilder create DockerSandboxBuilder
func NewDockerSandboxBuilder(envUnit commservices.EnvironmentsUnit) *DockerSandboxBuilder {
	return &DockerSandboxBuilder{
		deps: deps{
			EnvironmentsUnit: envUnit,
		},
	}
}

// DockerSandboxBuilderFactory create DockerSandboxBuilder
func DockerSandboxBuilderFactory(dp dependency.Provider) (ins interface{}, err error) {
	builder := &DockerSandboxBuilder{}
	if err = dp.InjectTo(&builder.deps); err != nil {
		return nil, err
	}
	return pipservices.SandboxBuilder(builder), nil
}

// Is return true if name is match to terminal factory
func (factory *DockerSandboxBuilder) Is(name string) bool {
	return strings.HasPrefix(name, "docker:") || strings.HasPrefix(name, "dockerb:")
}

// Build return terminal sandbox
func (factory *DockerSandboxBuilder) Build(name string) (sandbox pipservices.Sandbox, err error) {
	if strings.HasPrefix(name, "docker:") {
		return NewDockerSandbox(name[len("docker:"):], "sh", factory.deps)
	} else if strings.HasPrefix(name, "dockerb:") {
		return NewDockerSandbox(name[len("dockerb:"):], "bash", factory.deps)
	}
	return nil, goaterr.Errorf("incorrect sandbox %s", name)

}
