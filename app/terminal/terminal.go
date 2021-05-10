package terminal

import (
	"github.com/goatcms/goatcore/app"
)

// TerminalManager is implements app.TerminalManager
type TerminalManager struct {
	Arguments
	Commands
}

// NewTerminalManager create new instance of app.TerminalManager
func NewTerminalManager() (terminal app.TerminalManager) {
	return &TerminalManager{
		Arguments: *newArguments(),
		Commands:  *newCommands(),
	}
}
