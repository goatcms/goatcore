package terminalsb

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/modules"
	"github.com/goatcms/goatcore/app/modules/pipelinem/services"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

// TerminalSandbox is termal sandbox
type TerminalSandbox struct {
	terminal modules.Terminal
}

// NewTerminalSandbox create a TerminalSandbox instance
func NewTerminalSandbox(terminal modules.Terminal) (ins services.Sandbox, err error) {
	if terminal == nil {
		return nil, goaterr.Errorf("terminal argument is required")
	}
	return &TerminalSandbox{
		terminal: terminal,
	}, nil
}

// Run run code in sandbox
func (sandbox *TerminalSandbox) Run(ctx app.IOContext) (err error) {
	return sandbox.terminal.RunLoop(ctx)
}
