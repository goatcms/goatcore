package terminal

import (
	"github.com/goatcms/goatcore/app"
)

func MergeCommands(base app.TerminalCommands, exts ...app.TerminalCommand) app.TerminalCommands {
	var merged []app.TerminalCommand
	for _, name := range base.CommandNames() {
		merged = append(merged, base.Command(name))
	}
	for _, command := range exts {
		merged = append(merged, command)
	}
	return NewCommands(merged...)
}
