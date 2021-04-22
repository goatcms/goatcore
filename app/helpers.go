package app

import (
	"strings"

	"github.com/goatcms/goatcore/varutil/goaterr"
)

// CloseApp close application
func CloseApp(a App) error {
	// close in inverted order (to init)
	return goaterr.ToError(goaterr.AppendError(nil,
		a.CommandScope().Trigger(KillEvent, nil),
		a.AppScope().Trigger(KillEvent, nil),
		a.FilespaceScope().Trigger(KillEvent, nil),
		a.ArgsScope().Trigger(KillEvent, nil),
		a.EngineScope().Trigger(KillEvent, nil),
	))
}

// RegisterCommand add new command to application
func RegisterCommand(a App, name string, callback CommandCallback, help string) (err error) {
	return RegisterScopeCommand(a.CommandScope(), name, callback, help)
}

// RegisterScopeCommand add new command to scope
func RegisterScopeCommand(scp Scope, name string, callback CommandCallback, help string) (err error) {
	name = strings.ToLower(name)
	scp.Set("help.command."+name, help)
	scp.Set("command."+name, callback)
	return nil
}

// RegisterHealthChecker add new health checker to application
func RegisterHealthChecker(a App, name string, callback HealthCheckerCallback) (err error) {
	return RegisterScopeHealthChecker(a.CommandScope(), name, callback)
}

// RegisterHealthChecker add new health checker to scope
func RegisterScopeHealthChecker(scp Scope, name string, callback HealthCheckerCallback) (err error) {
	scp.Set("health."+name, callback)
	return nil
}

// RegisterArgument add new argument definition to application
func RegisterArgument(a App, name string, help string) (err error) {
	return RegisterScopeArgument(a.CommandScope(), name, help)
}

// RegisterScopeArgument add new argument definition to scope
func RegisterScopeArgument(scp Scope, name string, help string) (err error) {
	scp.Set("help.argument."+name, help)
	return nil
}
