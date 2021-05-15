package containersb

import (
	"strings"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/modules/commonm/commservices"
	"github.com/goatcms/goatcore/app/modules/ocm/ocservices"
	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

// ContainerSandboxBuilder create new container sandbox builder
type ContainerSandboxBuilder struct {
	deps deps
}

// NewContainerSandboxBuilder create ContainerSandboxBuilder
func NewContainerSandboxBuilder(envUnit commservices.EnvironmentsUnit, ocManager ocservices.Manager) *ContainerSandboxBuilder {
	return &ContainerSandboxBuilder{
		deps: deps{
			EnvironmentsUnit: envUnit,
			OCManager:        ocManager,
		},
	}
}

// ContainerSandboxBuilderFactory create ContainerSandboxBuilder
func ContainerSandboxBuilderFactory(dp app.DependencyProvider) (ins interface{}, err error) {
	builder := &ContainerSandboxBuilder{}
	if err = dp.InjectTo(&builder.deps); err != nil {
		return nil, err
	}
	return pipservices.SandboxBuilder(builder), nil
}

// Is return true if name is match to terminal factory
func (factory *ContainerSandboxBuilder) Is(name string) bool {
	return strings.HasPrefix(name, "container:") || strings.HasPrefix(name, "containerb:")
}

// Build return terminal sandbox
func (factory *ContainerSandboxBuilder) Build(name string) (sandbox pipservices.Sandbox, err error) {
	if strings.HasPrefix(name, "container:") {
		return NewContainerSandbox(name[len("container:"):], "sh", factory.deps)
	} else if strings.HasPrefix(name, "containerb:") {
		return NewContainerSandbox(name[len("containerb:"):], "bash", factory.deps)
	}
	return nil, goaterr.Errorf("incorrect sandbox %s", name)
}
