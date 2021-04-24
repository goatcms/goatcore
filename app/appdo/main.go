package appdo

import "github.com/goatcms/goatcore/app"

// Close shutdown application by send close event
func Close(a app.App) (err error) {
	return a.AppScope().Close()
}

// Emit a event to all application scopes
func Emit(a app.App, event int, data interface{}) {
	a.ArgsScope().Trigger(event, nil)
	a.ConfigScope().Trigger(event, nil)
	a.EngineScope().Trigger(event, nil)
	a.CommandScope().Trigger(event, nil)
	a.FilespaceScope().Trigger(event, nil)
	a.AppScope().Trigger(event, nil)
}
