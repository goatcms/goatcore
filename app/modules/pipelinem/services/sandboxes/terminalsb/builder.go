package terminalsb

import (
	"github.com/goatcms/goatcore/app/modules"
	"github.com/goatcms/goatcore/app/modules/pipelinem/services"
	"github.com/goatcms/goatcore/dependency"
)

// SandboxBuilder create sandbox for current terminal
type SandboxBuilder struct {
	sandbox services.Sandbox
}

// NewSandboxBuilder create new SandboxBuilder instance
func NewSandboxBuilder(terminal modules.Terminal) (ins *SandboxBuilder, err error) {
	instance := &SandboxBuilder{}
	if instance.sandbox, err = NewTerminalSandbox(terminal); err != nil {
		return nil, err
	}
	return instance, nil
}

// SandboxBuilderFactory create a SandboxBuilder instance
func SandboxBuilderFactory(dp dependency.Provider) (ins interface{}, err error) {
	var deps struct {
		Terminal modules.Terminal `dependency:"TerminalService"`
	}
	instance := &SandboxBuilder{}
	if err = dp.InjectTo(&deps); err != nil {
		return nil, err
	}
	if instance.sandbox, err = NewTerminalSandbox(deps.Terminal); err != nil {
		return nil, err
	}
	return services.SandboxBuilder(instance), nil
}

// Is return true if name is match to terminal factory
func (factory *SandboxBuilder) Is(name string) bool {
	return name == "" || name == "terminal"
}

// Build return terminal sandbox
func (factory *SandboxBuilder) Build(name string) (sandbox services.Sandbox, err error) {
	return factory.sandbox, nil
}
