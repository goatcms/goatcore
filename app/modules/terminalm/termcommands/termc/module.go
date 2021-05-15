package termc

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/terminal"
)

func ExitCommand() app.TerminalCommand {
	return terminal.NewCommand(terminal.CommandParams{
		Name:     "exit",
		Callback: RunExit,
		Help:     "close application",
	})
}

func TerminalCommand() app.TerminalCommand {
	return terminal.NewCommand(terminal.CommandParams{
		Name:     "terminal",
		Callback: RunTerminal,
		Help:     "Show application health status",
	})
}

func Commands() []app.TerminalCommand {
	return []app.TerminalCommand{
		ExitCommand(),
		TerminalCommand(),
	}
}
