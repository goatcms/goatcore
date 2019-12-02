package dockersb

import (
	"strings"

	"github.com/goatcms/goatcore/app/modules/pipelinem/services"
	"github.com/goatcms/goatcore/dependency"
)

// DockerSandboxBuilder create new docker sandbox builder
type DockerSandboxBuilder struct {
}

// NewDockerSandboxBuilder create DockerSandboxBuilder
func NewDockerSandboxBuilder() *DockerSandboxBuilder {
	return &DockerSandboxBuilder{}
}

// DockerSandboxBuilderFactory create DockerSandboxBuilder
func DockerSandboxBuilderFactory(dp dependency.Provider) (ins interface{}, err error) {
	return services.SandboxBuilder(&DockerSandboxBuilder{}), nil
}

// Is return true if name is match to terminal factory
func (factory *DockerSandboxBuilder) Is(name string) bool {
	return strings.HasPrefix(name, "docker:")
}

// Build return terminal sandbox
func (factory *DockerSandboxBuilder) Build(name string) (sandbox services.Sandbox, err error) {
	return NewDockerSandbox(name[len("docker:"):])
}
