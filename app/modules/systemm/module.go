package systemm

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/goatcms/goatcore/app"
)

// Module is command unit
type Module struct{}

// NewModule create new command module instance
func NewModule() app.Module {
	return &Module{}
}

// RegisterDependencies is init callback to register module dependencies
func (m *Module) RegisterDependencies(a app.App) error {
	return nil
}

// InitDependencies is init callback to inject dependencies inside module
func (m *Module) InitDependencies(a app.App) error {
	return nil
}

// Run start command line loop
func (m *Module) Run(a app.App) (err error) {
	var (
		appScope = a.AppScope()
		sigs     = make(chan os.Signal, 1)
	)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		var (
			sig  os.Signal
			more bool
		)
		if sig, more = <-sigs; !more {
			return
		}
		if sig == syscall.SIGINT || sig == syscall.SIGTERM {
			appScope.Kill()
		}
	}()
	appScope.On(app.KillEvent, func(i interface{}) error {
		signal.Reset(syscall.SIGINT, syscall.SIGTERM)
		return nil
	})
	return nil
}
