package selfsb

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices"
	"github.com/goatcms/goatcore/app/modules/terminalm/termservices"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

// SelfSandbox is termal sandbox
type SelfSandbox struct {
	terminal termservices.Terminal
}

// NewSelfSandbox create a SelfSandbox instance
func NewSelfSandbox(terminal termservices.Terminal) (ins pipservices.Sandbox, err error) {
	if terminal == nil {
		return nil, goaterr.Errorf("terminal argument is required")
	}
	return &SelfSandbox{
		terminal: terminal,
	}, nil
}

// Run run code in sandbox
func (sandbox *SelfSandbox) Run(ctx app.IOContext) (err error) {
	return sandbox.terminal.RunLoop(ctx, "")
}
