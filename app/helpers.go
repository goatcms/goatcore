package app

import (
	"github.com/goatcms/goatcore/varutil/goaterr"
)

// CloseApp close application
func CloseApp(a App) error {
	// close in inverted order (to init)
	return goaterr.ToErrors(goaterr.AppendError(nil,
		a.CommandScope().Trigger(KillEvent, nil),
		a.AppScope().Trigger(KillEvent, nil),
		a.DependencyScope().Trigger(KillEvent, nil),
		a.FilespaceScope().Trigger(KillEvent, nil),
		a.ArgsScope().Trigger(KillEvent, nil),
		a.EngineScope().Trigger(KillEvent, nil),
	))
}

// RegisterComand add new command to application
func RegisterComand(a App, name string, callback CommandCallback, help string) (err error) {
	commandScope := a.CommandScope()
	commandScope.Set("help.command."+name, help)
	commandScope.Set("command."+name, callback)
	return nil
}

// RegisterArgument add new argument definition to application
func RegisterArgument(a App, name string, help string) (err error) {
	a.CommandScope().Set("help.argument."+name, help)
	return nil
}
