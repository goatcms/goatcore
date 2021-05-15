package helpc

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/terminal"
)

func HealtCommand() app.TerminalCommand {
	return terminal.NewCommand(terminal.CommandParams{
		Name:     "health",
		Callback: RunHealth,
		Help:     "Show application health status",
	})
}

func HelpCommand() app.TerminalCommand {
	return terminal.NewCommand(terminal.CommandParams{
		Name:          "help",
		Callback:      RunHelp,
		Help:          "Show help (add command name to show command details",
		MainArguments: []string{"name"},
		Arguments: terminal.NewArguments(
			terminal.NewArgument(terminal.ArgumentParams{
				Name:     "name",
				Required: false,
				Type:     app.TerminalTextArgument,
			}),
		),
	})
}

func Commands() []app.TerminalCommand {
	return []app.TerminalCommand{
		HelpCommand(),
		HealtCommand(),
	}
}
