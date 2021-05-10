package selfsb

import (
	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices"
	"github.com/goatcms/goatcore/app/modules/terminalm/termservices"
	"github.com/goatcms/goatcore/dependency"
)

// SandboxBuilder create sandbox for current terminal
type SandboxBuilder struct {
	sandbox pipservices.Sandbox
}

// NewSandboxBuilder create new SandboxBuilder instance
func NewSandboxBuilder(terminal termservices.Terminal) (ins *SandboxBuilder, err error) {
	instance := &SandboxBuilder{}
	if instance.sandbox, err = NewSelfSandbox(terminal); err != nil {
		return nil, err
	}
	return instance, nil
}

// SandboxBuilderFactory create a SandboxBuilder instance
func SandboxBuilderFactory(dp dependency.Provider) (ins interface{}, err error) {
	var deps struct {
		Terminal termservices.Terminal `dependency:"TerminalService"`
	}
	instance := &SandboxBuilder{}
	if err = dp.InjectTo(&deps); err != nil {
		return nil, err
	}
	if instance.sandbox, err = NewSelfSandbox(deps.Terminal); err != nil {
		return nil, err
	}
	return pipservices.SandboxBuilder(instance), nil
}

// Is return true if name is match to terminal factory
func (factory *SandboxBuilder) Is(name string) bool {
	return name == "" || name == "self"
}

// Build return terminal sandbox
func (factory *SandboxBuilder) Build(name string) (sandbox pipservices.Sandbox, err error) {
	return factory.sandbox, nil
}
