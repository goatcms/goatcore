package termc

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/terminal"
)

func Commands() []app.TerminalCommand {
	return []app.TerminalCommand{
		terminal.NewCommand(terminal.CommandParams{
			Name:     "health",
			Callback: RunHealth,
			Help:     "Show application health status",
		}),
		terminal.NewCommand(terminal.CommandParams{
			Name:     "help",
			Callback: RunHelp,
			Help:     "Show help",
		}),
		terminal.NewCommand(terminal.CommandParams{
			Name:     "terminal",
			Callback: RunHelp,
			Help:     "Run internal terminal",
		}),
	}
}
